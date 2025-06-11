package cmd

import (
	"context"
	"fmt"
	"golang/config"
	"golang/internal/controller"
	"golang/internal/repository"
	"golang/internal/usecase"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"golang/proto/kafka_pb"

	"google.golang.org/grpc"
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
            repository.NewKafkaProducer,
            usecase.NewKafkaUC,
            controller.NewKafkaController,
            controller.NewGRPCController,
        ),
        fx.Invoke(registerHTTPServer,RegisterGRPCServer),
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

func RegisterGRPCServer(
	lc fx.Lifecycle,
	conf *config.Config,
	handler *controller.GRPCController,
) {
	grpcServer := grpc.NewServer()
	kafka_pb.RegisterKafkaServiceServer(grpcServer, handler)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				address := fmt.Sprintf(":%s", conf.GRPC.Port)
				listener, err := net.Listen("tcp", address)
				if err != nil {
					fmt.Printf("failed to listen: %v\n", err)
					return
				}
				fmt.Printf("gRPC server running at %s\n", address)
				if err := grpcServer.Serve(listener); err != nil {
					fmt.Printf("failed to start gRPC server: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("gRPC server shutting down...")
			grpcServer.GracefulStop()
			return nil
		},
	})
}
