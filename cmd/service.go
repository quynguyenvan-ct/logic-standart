package cmd

import (
	"context"
	"fmt"
	"golang/config"
	"golang/internal/controller"
	"golang/internal/repository"
	"golang/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Command 'service'
var service = &cobra.Command{
    Use:   "service",
    Short: "Run HTTP API server",
    Run: func(cmd *cobra.Command, args []string) {
        app := buildApp()
        if err := app.Start(context.Background()); err != nil {
            fmt.Printf("failed to start service: %v\n", err)
            return
        }

        // Wait until the app is done (Ctrl+C or SIGTERM)
        <-app.Done()

        if err := app.Stop(context.Background()); err != nil {
            fmt.Printf("failed to stop service: %v\n", err)
        }
    },
}

// buildApp constructs the fx.App with dependencies
func buildApp() *fx.App {
    conf := config.LoadConfig()
    return fx.New(
        fx.WithLogger(func() fxevent.Logger {
            return &fxevent.ConsoleLogger{W: os.Stdout}
        }),
        fx.Provide(
            func() *config.Config { return conf },
            repository.NewUserRepository,
            usecase.NewUserUsecase,
            controller.NewController,
        ),
        fx.Invoke(registerHTTPServer),
    )
}

// registerHTTPServer starts the Gin HTTP server
func registerHTTPServer(
    lc fx.Lifecycle,
    conf *config.Config,
    ctrl *controller.Controller,
) {
    router := gin.Default()
    ctrl.RegisterRoutes(router)

    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            go func() {
                fmt.Printf("HTTP server running at port %s\n", conf.HTTP.Port)
                if err := router.Run(":" + conf.HTTP.Port); err != nil {
                    fmt.Printf("failed to start HTTP server: %v\n", err)
                }
            }()
            return nil
        },
        OnStop: func(ctx context.Context) error {
            fmt.Println("HTTP server shutting down...")
            return nil
        },
    })
}
