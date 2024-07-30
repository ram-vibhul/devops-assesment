package config

type Config struct {
	App      App      `yaml:"app"`
	Postgres Postgres `yaml:"postgres"`
}

type App struct {
	Versions    map[string]uint `yaml:"versions"`
	Environment string          `yaml:"environment"`
	Host        string          `yaml:"host"`
	Port        string          `yaml:"port"`
}

type Postgres struct {
	Host        string `yaml:"host"`
	Database    string `yaml:"database"`
	UserName    string `yaml:"userName"`
	Password    string `yaml:"password"`
	SSLmode     string `yaml:"sslMode"`
	Automigrate bool   `yaml:"automigrate"`
}
