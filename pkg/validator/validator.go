package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/skyapps-id/monolithic-modular-go/pkg/apperror"
)

// Validator adapts go-playground/validator to echo.Validator, turning
// validation failures into *apperror.AppError with per-field details.
type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	v := validator.New()
	// Report errors using the json field name instead of the Go field name.
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "" || name == "-" {
			return field.Name
		}
		return name
	})
	return &Validator{validate: v}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}

	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return apperror.ValidationError(err.Error())
	}

	details := make(map[string]string, len(verrs))
	for _, fe := range verrs {
		details[fe.Field()] = messageFor(fe)
	}
	return apperror.ValidationError("validation failed").WithDetails(details)
}

func messageFor(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email"
	case "gt":
		return fe.Field() + " must be greater than " + fe.Param()
	case "min":
		return fe.Field() + " must be at least " + fe.Param()
	case "max":
		return fe.Field() + " must be at most " + fe.Param()
	default:
		return fe.Field() + " is invalid"
	}
}
