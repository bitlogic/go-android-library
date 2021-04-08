package library

import "log"

// MobileConfiguration .
type MobileConfiguration struct {
	Database DatabaseService
}

//InitServerApp sólo contiene la llamada a la inicialización a la base de datos,
// pero se deja así en caso de que en algún momento sea necesario agregarle más propiedades
func (configuration *MobileConfiguration) InitServerApp(databasepath string) {
	configuration.InitDatabase(databasepath)
}

//InitDatabase inicializa la base de datos
func (configuration *MobileConfiguration) InitDatabase(databasepath string) {
	dao, err := NewDao(databasepath)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	configuration.Database = NewDBService(dao)
}
