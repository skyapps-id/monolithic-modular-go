package bootstrap

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/monolithic-modular-go/internal/config"
	"github.com/skyapps-id/monolithic-modular-go/internal/router"
	"github.com/skyapps-id/monolithic-modular-go/pkg/logger"
)

func Run(configPath string, modules []router.Module) {
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger.Init(cfg.Server.LogLevel)

	selected, err := pickModules(modules, cfg.Modules)
	if err != nil {
		log.Fatalf("failed to pick modules: %v", err)
	}

	if len(selected) == 0 {
		fmt.Println("no modules to run")
		os.Exit(0)
	}

	for _, m := range selected {
		log.Printf("loading module: %s", m.Name())
	}

	app := router.NewApp(echo.New(), cfg.Server.APIPrefix)
	app.Register(selected...)

	if err := app.Serve(cfg.Server.Addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func pickModules(all []router.Module, names []string) ([]router.Module, error) {
	if len(names) == 0 {
		return all, nil
	}

	pick := make(map[string]bool, len(names))
	for _, n := range names {
		pick[n] = true
	}

	var selected []router.Module
	for _, m := range all {
		if pick[m.Name()] {
			selected = append(selected, m)
			delete(pick, m.Name())
		}
	}

	if len(pick) > 0 {
		missing := make([]string, 0, len(pick))
		for n := range pick {
			missing = append(missing, n)
		}
		return nil, fmt.Errorf("module(s) not found: %s", strings.Join(missing, ", "))
	}

	return selected, nil
}
