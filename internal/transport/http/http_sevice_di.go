package http

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func (t *Delivery) InjectEcho(e *echo.Echo) error {
	if e == nil {
		return errors.New("invalid echo")
	}
	t.e = e
	return nil
}

func (t *Delivery) InjectMonitoringController(ctrl *MonitoringController) error {
	if ctrl == nil {
		return errors.New("invalid monitoring controller")
	}
	t.monitoringCtrl = ctrl
	return nil
}
