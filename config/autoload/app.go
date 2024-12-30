package autoload

type AppConfig struct {
	AppName    string `env:"APP_NAME" default:"competition-tool-api"`
	ProjectEnv string `env:"PROJECT_ENV" default:"dev"`
}
