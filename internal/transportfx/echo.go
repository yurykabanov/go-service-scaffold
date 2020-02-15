package transportfx

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/echo-swagger"
	"github.com/yurykabanov/echo-logrus"
	"go.uber.org/fx"

	"github.com/yurykabanov/service-scaffold/docs"
	"github.com/yurykabanov/service-scaffold/pkg/transport/http/controller"
	"github.com/yurykabanov/service-scaffold/pkg/version"
)

type MainEchoServerParams struct {
	fx.In

	Middleware struct {
		fx.In

		JaegerTracingMiddleware echo.MiddlewareFunc `name:"jaegertracing"`

		// TODO: Add your middleware here:
		// MyAwesomeMiddleware echo.MiddlewareFunc `name:"my_awesome_middleware"`
	}

	Controllers struct {
		fx.In

		// TODO: Add your controllers here:
		// MyAwesomeController controllers.MyAwesomeController
	}
}

func EchoServerProvider(
	params MainEchoServerParams,
	logrusLogger *logrus.Logger,
	stdLogger *log.Logger,
) *echo.Echo {
	server := echo.New()

	// Disable annoying logs
	server.HideBanner = true
	server.HidePort = true

	// NOTE: Echo.Logger is used in
	// - `DefaultHTTPErrorHandler` when it's unable to send error to client
	// - `echo.Response` when response is already committed
	// and is obtainable via:
	// - `echo.Context`'s `Logger` method as a fallback logger
	//
	// NOTE: `echo.Echo` messes with log level, given adapter ignores all `SetLevel` calls
	server.Logger = echologrus.LoggerAdapter{Logger: logrusLogger}

	// NOTE: Echo.StdLogger is a logger that set into `http.Server`'s `ErrorLog`
	server.StdLogger = stdLogger

	server.HTTPErrorHandler = controller.EchoErrorHandler

	// Middleware
	{
		server.Use(middleware.Recover())
		server.Use(echologrus.Middleware(logrusLogger))
		server.Use(middleware.RequestID())
		server.Use(params.Middleware.JaegerTracingMiddleware)
	}

	// Routes
	{
		// TODO: add your router here
		api := server.Group("/api/:version")
		_ = api

		server.GET("/version", controller.NewVersionController().Version)

		// TODO: consider adding "enable swagger docs" switch
		docs.SwaggerInfo.Version = version.Version
		server.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	return server
}

type EchoHttpServerHolder struct {
	fx.In

	EchoServer *echo.Echo
	HttpServer *http.Server
}

func RunServer(lc fx.Lifecycle, s fx.Shutdowner, holder EchoHttpServerHolder) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := holder.EchoServer.StartServer(holder.HttpServer)
				if err != nil {
					s.Shutdown()
				}
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			if err := holder.EchoServer.Shutdown(ctx); err != http.ErrServerClosed {
				return err
			}

			return nil
		},
	})
}
