package postgres

type Config struct {
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	Schema     string `yaml:"schema"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Driver     string `yaml:"driver"`
	Migrations string `yaml:"migrations"`
}
