package gormtest

import "github.com/jinzhu/gorm"

type EsbUser struct {
	gorm.Model
	Phone     string    `json:"phone"`
	NickName  string    `json:"nick_name"`
	Name      string    `json:"name"`
	Pwd       string    `json:"pwd"`
	PlainText string    `json:"plain_text"`
	State     int       `json:"state"`
	Roles     []EsbRole `json:"roles" gorm:"many2many:esb_user_role"`
	//MinUser   *WxMinUser `json:"minUser" gorm:"foreignkey:UserID"`
	Ip string `json:"ip"`
}

const (
	USER_STATE_BAN    = iota // 封号
	USER_STATE_NORMAL        //正常
)

type EsbRole struct {
	gorm.Model
	Name      string        `json:"name"`
	State     int           `json:"state"`
	Resources []EsbResource `json:"resources" gorm:"many2many:esb_role_resource"`
}

type EsbResource struct {
	gorm.Model
	Uri        string `json:"uri"`
	Type       int    `json:"type"`
	State      int    `json:"state"`
	VerifyType int    `json:"verifyType"`
}

type EsbUsersResponse struct {
	//PaginationRep
	Users []*EsbUser `json:"users"`
}
