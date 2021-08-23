package configs

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}
