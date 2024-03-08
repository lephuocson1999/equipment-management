package cmd

import (
	"context"
	"equipment-management/config"
	devicefx "equipment-management/di/devicefx"
	"equipment-management/di/postgresfx"
	"fmt"

	device_http "equipment-management/service/device/router"
	"errors"
	"log"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		NewServer().Run()
	},
	Version: "1.0.0",
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	app := fx.New(
		fx.Invoke(config.InitConfig),
		devicefx.Module,
		postgresfx.Module,
		fx.Provide(provideGinEngine),
		fx.Invoke(
			registerService,
		),
		fx.Invoke(startServer),
	)
	app.Run()
}

func provideGinEngine() *gin.Engine {
	return gin.Default()
}

func registerService(
	g *gin.Engine,
	deviceRouter device_http.DeviceRouter,
) {
	api := g.Group("")

	deviceRouter.Register(api)
}

func startServer(lifecycle fx.Lifecycle, g *gin.Engine) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					addr := fmt.Sprintf(":%v", config.ServerConfig().Port)
					err := endless.ListenAndServe(addr, g)
					if err != nil && !errors.Is(err, http.ErrServerClosed) {
						log.Fatal("failed to listen and serve from server: %v", err.Error())
					}
				}()
				return nil
			},
		},
	)
}
