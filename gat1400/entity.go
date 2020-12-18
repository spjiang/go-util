package gat1400

// Gat1400Object 数据实体
type Gat1400Object struct {
	MotorVehicleListObject MotorVehicleListObject `json:"MotorVehicleListObject"`
}

// MotorVehicleListObject 机动车抓拍对象实体列表
type MotorVehicleListObject struct {
	MotorVehicleObject []MotorVehicleObject `json:"MotorVehicleObject"`
}

// MotorVehicleObject
type MotorVehicleObject struct {
	MotorVehicleID    string       `json:"MotorVehicleID"`
	InfoKind          int64        `json:"InfoKind"`
	SourceID          string       `json:"SourceID"`
	StorageURL1       string       `json:"StorageUrl1"`
	HasPlate          string       `json:"HasPlate"`
	VehicleColor      string       `json:"VehicleColor"`
	PlateClass        string       `json:"PlateClass"`
	PlateColor        string       `json:"PlateColor"`
	PlateNo           string       `json:"PlateNo"`
	TollgateID        string       `json:"TollgateID"`
	DeviceID          string       `json:"DeviceID"`
	LeftTopX          int64        `json:"LeftTopX"`
	LeftTopY          int64        `json:"LeftTopY"`
	RightBtmX         int64        `json:"RightBtmX"`
	RightBtmY         int64        `json:"RightBtmY"`
	MarkTime          string       `json:"MarkTime"`
	AppearTime        string       `json:"AppearTime"`
	DisAppearTime     string       `json:"DisAppearTime"`
	LaneNo            int64        `json:"LaneNo"`
	Speed             float64      `json:"Speed"`
	Direction         string       `json:"Direction"`
	DrivingStatusCode string       `json:"DrivingStatusCode"`
	VehicleClass      string       `json:"VehicleClass"`
	VehicleBrand      string       `json:"VehicleBrand"`
	VehicleModel      string       `json:"VehicleModel"`
	VehicleStyles     string       `json:"VehicleStyles"`
	VehicleLength     int64        `json:"VehicleLength"`
	VehicleColorDepth string       `json:"VehicleColorDepth"`
	PassTime          string       `json:"PassTime"`
	NameOfPassedRoad  string       `json:"NameOfPassedRoad"`
	Sunvisor          int64        `json:"Sunvisor"`
	PlateReliability  string       `json:"PlateReliability"`
	SubImageList      SubImageList `json:"SubImageList"`
}

// SubImageList ...
type SubImageList struct {
	SubImageInfoObject []SubImageInfoObject `json:"SubImageInfoObject"`
}

// SubImageInfoObject ...
type SubImageInfoObject struct {
	ImageID     string `json:"ImageID"`
	EventSort   int64  `json:"EventSort"`
	DeviceID    string `json:"DeviceID"`
	StoragePath string `json:"StoragePath"`
	Type        string `json:"Type"`
	FileFormat  string `json:"FileFormat"`
	ShotTime    string `json:"ShotTime"`
	Width       int64  `json:"Width"`
	Height      int64  `json:"Height"`
	Data        string `json:"Data"`
}
