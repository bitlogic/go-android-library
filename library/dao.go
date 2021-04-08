package library

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

type DaoRepository interface {
	AddUser(userInfo *User) error
	getUserDevices(userID int64) ([]*Device, error)
}

func NewDao(databasePath string) (*Dao, error) {
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening or creating the database %v ", err)
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("error migrating user model")
		return nil, err
	}

	err = db.AutoMigrate(&Device{})
	if err != nil {
		log.Fatal("error migrating devices model")
		return nil, err
	}

	dao := Dao{db}
	log.Println("database correctly migrated")
	return &dao, nil
}

//AddUser: adds new user to the database
func (dao *Dao) AddUser(user *User) error {
	currentConnection := dao.db
	result := currentConnection.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetUser
func (dao *Dao) getUserDevices(userID int64) ([]*Device, error) {
	var devices []*Device

	currentConnection := dao.db
	query := "user_id = ? "
	var args []interface{}
	args = append(args, userID)

	result := currentConnection.Where(query, args...).Find(&devices)
	if result.Error != nil {
		err := result.Error
		return []*Device{}, err
	}

	return devices, nil
}
