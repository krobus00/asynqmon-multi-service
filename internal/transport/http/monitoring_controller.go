package http

import (
	"fmt"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/krobus00/asynqmon-multi-service/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type monitoringEndpoint struct {
	Name string
	URL  string
}

type MonitoringController struct {
	handlers     map[string]*asynqmon.HTTPHandler
	endpointList []monitoringEndpoint
}

func NewMonitoringController() (*MonitoringController, error) {
	controller := new(MonitoringController)
	controller.handlers = map[string]*asynqmon.HTTPHandler{}
	controller.endpointList = make([]monitoringEndpoint, 0)

	for _, service := range config.Services() {
		logrus.Info(fmt.Sprintf("register %s handler", service.Name))
		opts, err := redis.ParseURL(service.RedisHost)
		if err != nil {
			return nil, err
		}
		endpoint := fmt.Sprintf("%s/%s", config.MonitoringPath(), service.Name)
		handler := asynqmon.New(asynqmon.Options{
			RootPath: endpoint,
			RedisConnOpt: asynq.RedisClientOpt{
				Addr:     opts.Addr,
				Password: opts.Password,
				DB:       opts.DB,
			},
		})
		controller.handlers[service.Name] = handler
		controller.endpointList = append(controller.endpointList, monitoringEndpoint{
			Name: service.Name,
			URL:  endpoint,
		})
	}
	return controller, nil
}

func (t *MonitoringController) GetAsynqmonByServiceName(eCtx echo.Context) error {
	serviceName := eCtx.Param("serviceName")
	handler, ok := t.handlers[serviceName]
	if !ok {
		return eCtx.JSON(http.StatusNotFound, map[string]any{
			"message": "service not found",
		})
	}

	handler.ServeHTTP(eCtx.Response(), eCtx.Request())
	return nil
}

func (t *MonitoringController) Index(eCtx echo.Context) error {
	return eCtx.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title":    "List Service",
		"services": t.endpointList,
	})
}
