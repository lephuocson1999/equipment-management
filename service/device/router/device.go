package router

import (
	"equipment-management/pkg/middleware"
	"equipment-management/service/device/model"
	"equipment-management/service/device/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceRouter interface {
	Register(routerGroup gin.IRouter)
}

type deviceRouterIml struct {
	deviceUsecase usecase.DeviceUseCase
}

func NewDeviceRouter(
	deviceUsecase usecase.DeviceUseCase,
) DeviceRouter {
	return &deviceRouterIml{
		deviceUsecase: deviceUsecase,
	}
}

func (m *deviceRouterIml) Register(r gin.IRouter) {
	product := r.Group("/api/v1", middleware.Authentication)
	{
		product.GET("/devices", m.GetDevices)
	}
}

func (m *deviceRouterIml) GetDevices(c *gin.Context) {
	ctx := c.Request.Context()

	// Parse request
	var req model.GetDevicesRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := m.deviceUsecase.GetDevices(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data.Data,
		"total":   data.Total,
	})
}
