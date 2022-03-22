package config

type (
	// DB represent database configuration.
	DB struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Name string `yaml:"name"`
		Code string `yaml:"code"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	}
)
