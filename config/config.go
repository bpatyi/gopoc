package config

import validation "github.com/go-ozzo/ozzo-validation/v4"

const (
	EnvDebug      = "debug"
	EnvStaging    = "staging"
	EnvProduction = "production"

	FrameworkGin   = "gin"
	FrameworkEcho  = "echo"
	FrameworkFiber = "fiber"

	databaseTypeSqlite   = "sqlite"
	databaseTypeMongoDb  = "mongodb"
	databaseTypeMysql    = "mysql"
	databaseTypePostgres = "postgres"
)

type Config struct {
	Env          string
	Framework    string
	DatabaseType string
	Logger       *Logger
}

func (c Config) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Env, validation.Required, validation.In(EnvDebug, EnvStaging, EnvProduction)),
		validation.Field(&c.Framework, validation.Required, validation.In(FrameworkEcho, FrameworkFiber, FrameworkGin)),
		validation.Field(&c.DatabaseType, validation.Required, validation.In(
			databaseTypeMongoDb,
			databaseTypeMysql,
			databaseTypePostgres,
			databaseTypeSqlite,
		)),
		validation.Field(&c.Logger),
	)
}
