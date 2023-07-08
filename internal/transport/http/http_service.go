package http

import (
	"github.com/krobus00/asynqmon-multi-service/internal/config"
	"github.com/labstack/echo/v4"
)

type Delivery struct {
	e              *echo.Echo
	monitoringCtrl *MonitoringController
}

func NewDelivery() *Delivery {
	return new(Delivery)
}

func (t *Delivery) InitRoutes() {
	t.e.GET("", t.monitoringCtrl.Index)
	monitoringGroup := t.e.Group(config.MonitoringPath())
	monitoringGroup.Any("/:serviceName/*", t.monitoringCtrl.GetAsynqmonByServiceName)
}
