package library

type User struct {
	UserID  int64     `json:"user_id" gorm:"primary_key"`
	Devices []*Device `json:"devices" gorm:"foreignkey:UserID"`
}

type Device struct {
	DeviceID   int    `json:"device_id,omitempty" gorm:"primary_key"`
	DeviceType string `json:"device_type,omitempty"`
}
