package usecase

import (
	"context"
	"equipment-management/service/device/builder"
	"equipment-management/service/device/model"
	"equipment-management/service/device/repository"

	"github.com/labstack/gommon/log"
)

type DeviceUseCase interface {
	GetDevices(ctx context.Context, req model.GetDevicesRequest) (*model.GetDevicesResponse, error)
}

type deviceImpl struct {
	deviceRepository repository.DeviceRepo
}

func NewDeviceUseCase(deviceRepository repository.DeviceRepo) DeviceUseCase {
	return &deviceImpl{
		deviceRepository: deviceRepository,
	}
}

func (d *deviceImpl) GetDevices(
	ctx context.Context,
	req model.GetDevicesRequest,
) (*model.GetDevicesResponse, error) {
	request := builder.BuildFindRequestByGetDeviceRequest(req)
	data, total, err := d.deviceRepository.Find(ctx, request)
	if err != nil {
		log.Errorf("GetDevices: Find error: %s", err.Error())
		return nil, err
	}

	return &model.GetDevicesResponse{Data: data, Total: total}, nil
}
