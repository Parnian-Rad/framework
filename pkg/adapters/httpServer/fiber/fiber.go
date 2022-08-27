package fiber

import (
	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"

	sentryfiber "github.com/aldy505/sentry-fiber"
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber/v2"
)

type FiberHttpServer struct {
	app     *fiber.App
	address string
}

func New(address string) ports.HttpServer {
	if address == "" {
		address = "0.0.0.0:3000"
	}

	return &FiberHttpServer{
		app:     fiber.New(),
		address: address,
	}
}

func (fh *FiberHttpServer) ActiveApm() {
	fh.app.Use(apmfiber.Middleware())
}

func (fh *FiberHttpServer) ActiveSentry() {
	fh.app.Use(sentryfiber.New(sentryfiber.Options{}))
}

func (fh *FiberHttpServer) Listen() error {
	if err := fh.app.Listen(fh.address); err != nil {
		return err
	}
	return nil
}
