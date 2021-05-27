package gormtest

import "github.com/jinzhu/gorm"

//EsbProduct 商品
type EsbProduct struct {
	gorm.Model
	Name        string           `json:"name"`
	Notes       string           `json:"notes"`
	Content     string           `json:"content"`
	Price       int              `json:"price"`
	MarketPrice int              `json:"market_price"`
	Brokerage   int              `json:"brokerage"`
	State       int              `json:"state"`
	StockID     uint             `json:"stockId"`
	Rank        int              `json:"rank"`
	Stock       EsbProductStock  `json:"stock" gorm:"foreignkey:StockID"`
	Files       []EsbProductFile `json:"files" gorm:"foreignkey:PID"`
	UserID      uint             `json:"minUserId"`
	TypeID      uint             `json:"typeId"`
}

//EsbProductStock 商品库存销售信息
type EsbProductStock struct {
	ID        int `json:"stock_id,omitempty"`
	AddVolume int `json:"addVolume,omitempty" db:"add_volume"`
	SubVolume int `json:"subVolume,omitempty" db:"sub_volume"`
	Stock     int `json:"stock,omitempty"`
	// PID       int `json:"pid,omitempty"`
}

// 每个商品有一些文件。 文件有类型。

//EsbProductFile 商品图片信息
type EsbProductFile struct {
	PID  int     `json:"pid" gorm:"column:pid"`
	FID  int     `json:"fid" gorm:"column:fid"`
	Type int     `json:"type"`
	File EsbFile `json:"file" gorm:"foreignkey:FID"`
}

type EsbFile struct {
	ID       int64  `json:"id" gorm:"column:id"`
	PID      int64  `json:"pid" gorm:"column:pid"`
	Path     string `json:"path"`
	Abstarct string `json:"abstarct"`
	Suffix   string `json:"suffix"`
	Prefix   string `json:"prefix"`
	Type     string `json:"type"`
	Appid    string `json:"appid"`
	Count    int64  `json:"count"`
}
