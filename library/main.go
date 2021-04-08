package library

//NewAppServer crea la instancia necesaria para utilizar todos los métodos de la librería
func NewAppServer(databaseHost string) *MobileConfiguration {
	configuration := MobileConfiguration{}
	configuration.InitServerApp(databaseHost)

	return &configuration
}
