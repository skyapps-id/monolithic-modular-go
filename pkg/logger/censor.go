package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type censorHandler struct {
	handler slog.Handler
}

func newCensorHandler(h slog.Handler) slog.Handler {
	return &censorHandler{handler: h}
}

func (h *censorHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *censorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return newCensorHandler(h.handler.WithAttrs(censorAttrs(attrs)))
}

func (h *censorHandler) WithGroup(name string) slog.Handler {
	return newCensorHandler(h.handler.WithGroup(name))
}

func (h *censorHandler) Handle(ctx context.Context, r slog.Record) error {
	var attrs []slog.Attr
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, censorAttr(a))
		return true
	})

	newRecord := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	newRecord.AddAttrs(attrs...)

	_, file, line, _ := runtime.Caller(0)
	const maxDepth = 10
	for i := 0; i < maxDepth; i++ {
		_, f, l, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !strings.Contains(f, "logger/") && !strings.Contains(f, "slog") && !strings.Contains(f, "runtime/") {
			file, line = f, l
			break
		}
	}
	relFile := file
	if wd, err := os.Getwd(); err == nil {
		if cut, ok := strings.CutPrefix(file, wd+"/"); ok {
			relFile = cut
		}
	}
	newRecord.AddAttrs(slog.String("source", filepath.Join(relFile)+":"+strconv.Itoa(line)))

	return h.handler.Handle(ctx, newRecord)
}

func censorAttrs(attrs []slog.Attr) []slog.Attr {
	for i := range attrs {
		attrs[i] = censorAttr(attrs[i])
	}
	return attrs
}

func censorAttr(a slog.Attr) slog.Attr {
	if a.Value.Kind() == slog.KindGroup {
		return slog.Attr{Key: a.Key, Value: slog.GroupValue(censorAttrs(a.Value.Group())...)}
	}

	v := a.Value.Resolve()
	if v.Kind() != slog.KindAny {
		return a
	}

	raw := v.Any()
	if raw == nil {
		return a
	}

	censored := censorValue(reflect.ValueOf(raw))
	if censored != nil {
		return slog.Any(a.Key, censored)
	}
	return a
}

func censorValue(v reflect.Value) interface{} {
	if !v.IsValid() {
		return nil
	}

	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return v.Interface()
	}

	t := v.Type()
	result := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		fieldVal := v.Field(i)
		if !fieldVal.CanInterface() {
			continue
		}

		if field.Tag.Get("log") == "sensitive" {
			result[field.Name] = "******"
			continue
		}

		nested := censorValue(fieldVal)
		if nested != nil {
			result[field.Name] = nested
		}
	}

	if len(result) == 0 {
		return v.Interface()
	}
	return result
}
