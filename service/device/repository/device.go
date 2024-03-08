package repository

import (
	"context"

	"equipment-management/pkg/dto"
	"equipment-management/service/device/model"

	"gorm.io/gorm"
)

type DeviceRepo interface {
	Find(ctx context.Context, req dto.FindRequest) ([]model.Device, int64, error)
}

type deviceImp struct {
	db *gorm.DB
}

func NewDeviceRepo(db *gorm.DB) DeviceRepo {
	return &deviceImp{
		db: db,
	}
}

func (i *deviceImp) WithDeviceContext(ctx context.Context) *gorm.DB {
	return i.db.WithContext(ctx).Debug().Table("equipments").Model(&model.Device{})
}

func (i *deviceImp) Find(
	ctx context.Context,
	req dto.FindRequest,
) (devices []model.Device, total int64, err error) {
	tx := i.WithDeviceContext(ctx)

	for _, v := range req.Preloads {
		tx = tx.Preload(v)
	}

	if req.Join != "" {
		tx = tx.Joins(req.Join)
	}
	if req.GroupBy != "" {
		tx = tx.Group(req.GroupBy)
	}
	if req.Having != "" {
		tx = tx.Having(req.Having)
	}

	for _, v := range req.Queries {
		if v.Args != nil {
			tx = tx.Where(v.Query, v.Args)
			continue
		}

		tx = tx.Where(v.Query)
	}

	tx = tx.Count(&total)

	if req.Limit != 0 {
		tx = tx.Offset(req.Offset).Limit(req.Limit)
	}

	err = tx.Order(req.Sort).Find(&devices).Error
	if err != nil {
		return nil, total, err
	}

	return devices, total, nil
}
