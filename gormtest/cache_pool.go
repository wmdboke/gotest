package gormtest

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RlgwTrainDevice struct {
	TrainCode  string
	DeviceCode string
}

type RlgwDevice struct {
	gorm.Model
	DeviceCode      string
	DeviceState     byte
	NetState        byte
	LastOnline      time.Time
	FirmwareVersion string
	DiskTotalSize   float64
	DiskUsedSize    float64
	CreateUser      int
	UpdateUser      int
	TrainDevice     RlgwTrainDevice `gorm:"foreignkey:DeviceCode;association_foreignkey:DeviceCode"`
}

type RlgwTrain struct {
	gorm.Model
	TrainCode    string
	TrainStation string
	CreateUser   string
	UpdateUser   string
	Flag         string
}

type RlgwDeviceNetwork struct {
	DeviceCode string
	Mac        string
	Ip         string
}

type RlgwDeviceParams struct {
	ParamId    int16
	ParamCode  string
	ParamValue float32
	ParamInfo  string
}
