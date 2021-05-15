package config

type Configurations struct {
	Server   ServerConfigurations
	Database DatebaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatebaseConfigurations struct {
	DBName string
	DBUser string
	DBPass string
}
