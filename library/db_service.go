package library

import "errors"

//serviceProperties .
type serviceProperties struct {
	daoRepository DaoRepository
}

type DatabaseService interface {
	CreateUser(userInfo *User) (*User, error)
	getUserDevices(userID int64) ([]*Device, error)
}

//NewDBService .
func NewDBService(daoRepository DaoRepository) DatabaseService {
	return &serviceProperties{daoRepository}
}

func (service *serviceProperties) CreateUser(user *User) (*User, error) {
	err := service.daoRepository.AddUser(user)
	if err != nil && err.Error() == "UNIQUE constraint failed: users.user_id" {
		return &User{}, errors.New("user already exists")
	} else if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (service *serviceProperties) getUserDevices(userID int64) ([]*Device, error) {
	devices, err := service.daoRepository.getUserDevices(userID)

	if err != nil && err.Error() == "record not found" {
		return devices, errors.New("user not found") // vendrá vacío
	} else if err != nil {
		return devices, err
	}

	return devices, err
}
