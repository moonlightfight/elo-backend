package config

type Configurations struct {
	Server   ServerConfigurations
	Database DatebaseConfigurations
	ApiKeys  ApiKeysConfigurations
}

type ServerConfigurations struct {
	Port   int
	Secret string
}

type DatebaseConfigurations struct {
	DBName string
	DBUser string
	DBPass string
}

type ApiKeysConfigurations struct {
	Challonge string
	Smash     string
}
