package autoload

type AppConfig struct {
	AppName string `env:"name" default:"my-web"`
	version string `env:"version" default:"1.0"`
}
