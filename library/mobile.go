package library

import (
	"bytes"
	"encoding/json"
	"log"
)

//UserFactory: Los modelos exportados sólo contienen setters y getters para los tipos nativos (int, float, etc), pero no para los structs personalizados, como es el caso.
// Es por eso que se genera éste método, el cual tiene por objetivo inicializar un objeto usuario con sus respectivos "devices".
func (configuration *MobileConfiguration) UserFactory(device *Device) *User {
	return &User{
		Devices: []*Device{device},
	}
}

//AddUser Los objetos de tipo STRUCT deben pasarse por referencia al utilizarlos como parámetros (los retornos también)
func (configuration *MobileConfiguration) AddUser(userInfo *User) (*User, error) {
	responseUser, err := configuration.Database.CreateUser(userInfo)
	if err != nil {
		log.Printf("error adding user %v", err)
		return &User{}, err
	}
	log.Print("user correctly added")
	return responseUser, nil
}

//GetUserDevices devuelce un slice de BYTES, ya que actualmente el compilador de libs mobiles sólo admite este tipo de slices
func (configuration *MobileConfiguration) GetUserDevices(userID int64) ([]byte, error) {
	devices, err := configuration.Database.getUserDevices(userID)
	if err != nil {
		log.Printf("error searching user devices %v", err)
		return nil, err
	}

	sliceOfBytes := new(bytes.Buffer)
	err = json.NewEncoder(sliceOfBytes).Encode(devices)
	if err != nil {
		log.Printf("error encoding user devices information %v", err)
		return nil, err
	}

	log.Print("user devices were correctly retrieved")
	return sliceOfBytes.Bytes(), nil
}
