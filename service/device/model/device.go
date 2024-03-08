package model

type Device struct {
	ID                            int     `json:"id" gorm:"column:id"`
	Name                          string  `json:"name" gorm:"column:name"`
	Description                   string  `json:"description" gorm:"column:description"`
	Model                         string  `json:"model" gorm:"column:model"`
	Version                       string  `json:"version" gorm:"column:version"`
	Expense                       float64 `json:"expense" gorm:"column:expense"`
	DateOfManufacture             int64   `json:"date_of_manufacture" gorm:"column:date_of_manufacture"`
	Color                         string  `json:"color" gorm:"column:color"`
	BatteryLife                   string  `json:"battery_life" gorm:"column:battery_life"`
	Camera                        string  `json:"camera" gorm:"column:camera"`
	Length                        string  `json:"length" gorm:"column:length"`
	Width                         string  `json:"width" gorm:"column:width"`
	Height                        string  `json:"height" gorm:"column:height"`
	Weight                        string  `json:"weight" gorm:"column:weight"`
	Material                      string  `json:"material" gorm:"column:material"`
	Core                          string  `json:"core" gorm:"column:core"`
	Sensor                        string  `json:"sensor" gorm:"column:sensor"`
	Algorithm                     string  `json:"algorithm" gorm:"column:algorithm"`
	Category                      string  `json:"category" gorm:"column:category"`
	Language                      string  `json:"language" gorm:"column:language"`
	WorkingRadius                 string  `json:"working_radius" gorm:"column:working_radius"`
	ChargingVoltage               string  `json:"charging_voltage" gorm:"column:charging_voltage"`
	MotorType                     string  `json:"motor_type" gorm:"column:motor_type"`
	MovementMaximumspeed          string  `json:"movement_maximumspeed" gorm:"column:movement_maximumspeed"`
	Control                       string  `json:"control" gorm:"column:control"`
	CommunicationInterface        string  `json:"communication_interface" gorm:"column:communication_interface"`
	Bluetooth                     string  `json:"bluetooth" gorm:"column:bluetooth"`
	Wireless                      string  `json:"wireless" gorm:"column:wireless"`
	Manufacturer                  string  `json:"manufacturer" gorm:"column:manufacturer"`
	ProgrammingLanguage           string  `json:"programming_language" gorm:"column:programming_language"`
	ServiceLife                   string  `json:"service_life" gorm:"column:service_life"`
	ApplicableEquipment           string  `json:"applicable_equipment" gorm:"column:applicable_equipment"`
	Protocol                      string  `json:"protocol" gorm:"column:protocol"`
	Type                          string  `json:"type" gorm:"column:type"`
	LearningSpeed                 string  `json:"learning_speed" gorm:"column:learning_speed"`
	SuccessRate                   string  `json:"success_rate" gorm:"column:success_rate"`
	HoursOfOperation              string  `json:"hours_of_operation" gorm:"column:hours_of_operation"`
	RepetitivePositioningAccuracy string  `json:"repetitive_positioning_accuracy" gorm:"column:repetitive_positioning_accuracy"`
	Screen                        string  `json:"screen" gorm:"column:screen"`
	Other                         string  `json:"other" gorm:"column:other"`
}

type GetDevicesRequest struct {
	Q                             string   `json:"q" form:"q"`
	Model                         []string `json:"model" form:"model"`
	Version                       []string `json:"version" form:"version"`
	FromExpense                   float64  `json:"from_expense" form:"from_expense"`
	ToExpense                     float64  `json:"to_expense" form:"to_expense"`
	FromDateOfManufacture         int64    `json:"from_date_of_manufacture" form:"from_date_of_manufacture"`
	ToDateOfManufacture           int64    `json:"to_date_of_manufacture" form:"to_date_of_manufacture"`
	Color                         []string `json:"color" form:"color"`
	BatteryLife                   []string `json:"battery_life" form:"battery_life"`
	Camera                        []string `json:"camera" form:"camera"`
	Length                        []string `json:"length" form:"length"`
	Width                         []string `json:"width" form:"width"`
	Height                        []string `json:"height" form:"height"`
	Weight                        []string `json:"weight" form:"weight"`
	Material                      []string `json:"material" form:"material"`
	Core                          []string `json:"core" form:"core"`
	Sensor                        []string `json:"sensor" form:"sensor"`
	Algorithm                     []string `json:"algorithm" form:"algorithm"`
	Category                      []string `json:"category" form:"category"`
	Language                      []string `json:"language" form:"language"`
	WorkingRadius                 []string `json:"working_radius" form:"working_radius"`
	ChargingVoltage               []string `json:"charging_voltage" form:"charging_voltage"`
	MotorType                     []string `json:"motor_type" form:"motor_type"`
	MovementMaximumspeed          []string `json:"movement_maximumspeed" form:"movement_maximumspeed"`
	Control                       []string `json:"control" form:"control"`
	CommunicationInterface        []string `json:"communication_interface" form:"communication_interface"`
	Bluetooth                     []string `json:"bluetooth" form:"bluetooth"`
	Wireless                      []string `json:"wireless" form:"wireless"`
	Manufacturer                  []string `json:"manufacturer" form:"manufacturer"`
	ProgrammingLanguage           []string `json:"programming_language" form:"programming_language"`
	ServiceLife                   []string `json:"service_life" form:"service_life"`
	ApplicableEquipment           []string `json:"applicable_equipment" form:"applicable_equipment"`
	Protocol                      []string `json:"protocol" form:"protocol"`
	Type                          []string `json:"type" form:"type"`
	LearningSpeed                 []string `json:"learning_speed" form:"learning_speed"`
	SuccessRate                   []string `json:"success_rate" form:"success_rate"`
	HoursOfOperation              []string `json:"hours_of_operation" form:"hours_of_operation"`
	RepetitivePositioningAccuracy []string `json:"repetitive_positioning_accuracy" form:"repetitive_positioning_accuracy"`
	Screen                        []string `json:"screen" form:"screen"`
	Other                         []string `json:"other" form:"other"`
	Limit                         int      `json:"limit" form:"limit"`
	Page                          int      `json:"page" form:"page"`
	Sort                          string   `json:"sort" form:"sort"`
}

// func (g *GetDevicesRequest) InModels() string {
// 	var in string
// 	for _, model := range g.Model {

// 	}
// }

type GetDevicesResponse struct {
	Data  []Device `json:"data"`
	Total int64    `json:"total"`
}
