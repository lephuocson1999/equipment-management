package devicefx

import (
	"equipment-management/service/device/repository"
	device_http "equipment-management/service/device/router"
	"equipment-management/service/device/usecase"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(provideDeviceRepository, provideDeviceUsecase, provideDeviceRouter)

func provideDeviceUsecase(deviceRepo repository.DeviceRepo) usecase.DeviceUseCase {
	return usecase.NewDeviceUseCase(deviceRepo)
}

func provideDeviceRouter(
	deviceUsecase usecase.DeviceUseCase,
) device_http.DeviceRouter {
	return device_http.NewDeviceRouter(
		deviceUsecase,
	)
}

func provideDeviceRepository(
	db *gorm.DB,
) repository.DeviceRepo {
	return repository.NewDeviceRepo(db)
}
