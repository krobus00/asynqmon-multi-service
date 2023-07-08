package bootstrap

import (
	"context"
	"fmt"

	"github.com/krobus00/asynqmon-multi-service/internal/config"
	"github.com/krobus00/asynqmon-multi-service/internal/infrastructure"
	httpTransport "github.com/krobus00/asynqmon-multi-service/internal/transport/http"
	"github.com/krobus00/asynqmon-multi-service/internal/util"
	"github.com/sirupsen/logrus"
)

func StartServerBootstrap() {
	e := infrastructure.NewEcho()

	// init transport layer
	// controller
	monitoringCtrl, err := httpTransport.NewMonitoringController()
	util.ContinueOrFatal(err)

	// http service
	httpDelivery := httpTransport.NewDelivery()
	err = httpDelivery.InjectEcho(e)
	util.ContinueOrFatal(err)
	err = httpDelivery.InjectMonitoringController(monitoringCtrl)
	util.ContinueOrFatal(err)

	httpDelivery.InitRoutes()

	go func() {
		logrus.Info(fmt.Sprintf("http server started on :%s", config.PortHTTP()))
		err := e.Start(":" + config.PortHTTP())
		util.ContinueOrFatal(err)
	}()

	wait := gracefulShutdown(context.Background(), config.GracefulShutdownTimeOut(), map[string]operation{
		"http": func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	<-wait
}
