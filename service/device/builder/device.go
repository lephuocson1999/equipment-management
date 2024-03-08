package builder

import (
	"equipment-management/pkg/dto"
	"equipment-management/service/device/model"
	"fmt"
)

func BuildFindRequestByGetDeviceRequest(
	req model.GetDevicesRequest,
) dto.FindRequest {
	return dto.FindRequest{
		Queries: BuildQueryByGetDevicesRequest(req),
		Offset:  BuildOffset(req),
		Limit:   req.Limit,
		Sort:    req.Sort,
	}
}

func BuildOffset(
	req model.GetDevicesRequest,
) int {
	if req.Limit == 0 {
		req.Limit = 100
	}
	if req.Page == 0 {
		req.Page = 1
	}

	return (req.Page - 1) * req.Limit
}

func BuildQueryByGetDevicesRequest(
	req model.GetDevicesRequest,
) []dto.Query {
	var queries []dto.Query
	if req.Q != "" {
		query := fmt.Sprintf("name ILIKE '%%%s%%' OR description ILIKE '%%%s%%'", req.Q, req.Q)
		queries = append(queries, dto.Query{Query: query})
	}

	if len(req.Model) > 0 {
		queries = append(queries, dto.Query{Query: "model IN (?)", Args: req.Model})
	}

	if len(req.Version) > 0 {
		queries = append(queries, dto.Query{Query: "version IN (?)", Args: req.Version})
	}

	if req.FromExpense > 0 {
		queries = append(queries, dto.Query{Query: "expense >= ?", Args: req.FromExpense})
	}

	if req.ToExpense > 0 {
		queries = append(queries, dto.Query{Query: "expense <= ?", Args: req.ToExpense})
	}

	if req.FromDateOfManufacture > 0 {
		queries = append(queries, dto.Query{Query: "date_of_manufacture >=  ?", Args: req.FromDateOfManufacture})
	}

	if req.ToDateOfManufacture > 0 {
		queries = append(queries, dto.Query{Query: "date_of_manufacture <=  ?", Args: req.ToDateOfManufacture})
	}

	if len(req.Color) > 0 {
		queries = append(queries, dto.Query{Query: "color IN (?)", Args: req.Color})
	}

	if len(req.BatteryLife) > 0 {
		queries = append(queries, dto.Query{Query: "battery_life IN (?)", Args: req.BatteryLife})
	}

	if len(req.Camera) > 0 {
		queries = append(queries, dto.Query{Query: "camera IN (?)", Args: req.Camera})
	}

	if len(req.Length) > 0 {
		queries = append(queries, dto.Query{Query: "length IN (?)", Args: req.Length})
	}

	if len(req.Width) > 0 {
		queries = append(queries, dto.Query{Query: "width IN (?)", Args: req.Width})
	}

	if len(req.Weight) > 0 {
		queries = append(queries, dto.Query{Query: "weight IN (?)", Args: req.Weight})
	}

	if len(req.Material) > 0 {
		queries = append(queries, dto.Query{Query: "material IN (?)", Args: req.Material})
	}

	if len(req.Core) > 0 {
		queries = append(queries, dto.Query{Query: "core IN (?)", Args: req.Core})
	}

	if len(req.Sensor) > 0 {
		queries = append(queries, dto.Query{Query: "sensor IN (?)", Args: req.Sensor})
	}

	if len(req.Algorithm) > 0 {
		queries = append(queries, dto.Query{Query: "algorithm IN (?)", Args: req.Algorithm})
	}

	if len(req.Category) > 0 {
		queries = append(queries, dto.Query{Query: "category IN (?)", Args: req.Category})
	}

	if len(req.Language) > 0 {
		queries = append(queries, dto.Query{Query: "language IN (?)", Args: req.Language})
	}

	if len(req.WorkingRadius) > 0 {
		queries = append(queries, dto.Query{Query: "working_radius IN (?)", Args: req.WorkingRadius})
	}

	if len(req.ChargingVoltage) > 0 {
		queries = append(queries, dto.Query{Query: "charging_voltage IN (?)", Args: req.ChargingVoltage})
	}

	if len(req.MotorType) > 0 {
		queries = append(queries, dto.Query{Query: "motor_type IN (?)", Args: req.MotorType})
	}

	if len(req.MovementMaximumspeed) > 0 {
		queries = append(queries, dto.Query{Query: "movement_maximumspeed IN (?)", Args: req.MovementMaximumspeed})
	}

	if len(req.Control) > 0 {
		queries = append(queries, dto.Query{Query: "control IN (?)", Args: req.Control})
	}

	if len(req.CommunicationInterface) > 0 {
		queries = append(queries, dto.Query{Query: "communication_interface IN (?)", Args: req.CommunicationInterface})
	}

	if len(req.Bluetooth) > 0 {
		queries = append(queries, dto.Query{Query: "bluetooth IN (?)", Args: req.Bluetooth})
	}

	if len(req.Wireless) > 0 {
		queries = append(queries, dto.Query{Query: "wireless IN (?)", Args: req.Wireless})
	}

	if len(req.Manufacturer) > 0 {
		queries = append(queries, dto.Query{Query: "manufacturer IN (?)", Args: req.Manufacturer})
	}

	if len(req.ProgrammingLanguage) > 0 {
		queries = append(queries, dto.Query{Query: "programming_language IN (?)", Args: req.ProgrammingLanguage})
	}

	if len(req.ServiceLife) > 0 {
		queries = append(queries, dto.Query{Query: "service_life IN (?)", Args: req.ServiceLife})
	}

	if len(req.ApplicableEquipment) > 0 {
		queries = append(queries, dto.Query{Query: "applicable_equipment IN (?)", Args: req.ApplicableEquipment})
	}

	if len(req.Protocol) > 0 {
		queries = append(queries, dto.Query{Query: "protocol IN (?)", Args: req.Protocol})
	}

	if len(req.Type) > 0 {
		queries = append(queries, dto.Query{Query: "type IN (?)", Args: req.Type})
	}

	if len(req.LearningSpeed) > 0 {
		queries = append(queries, dto.Query{Query: "learning_speed IN (?)", Args: req.LearningSpeed})
	}

	if len(req.SuccessRate) > 0 {
		queries = append(queries, dto.Query{Query: "success_rate IN (?)", Args: req.SuccessRate})
	}

	if len(req.HoursOfOperation) > 0 {
		queries = append(queries, dto.Query{Query: "hours_of_operation IN (?)", Args: req.HoursOfOperation})
	}

	if len(req.RepetitivePositioningAccuracy) > 0 {
		queries = append(queries, dto.Query{Query: "repetitive_positioning_accuracy IN (?)", Args: req.RepetitivePositioningAccuracy})
	}

	if len(req.Screen) > 0 {
		queries = append(queries, dto.Query{Query: "screen IN (?)", Args: req.Screen})
	}

	if len(req.Other) > 0 {
		queries = append(queries, dto.Query{Query: "other IN (?)", Args: req.Other})
	}

	return queries
}
