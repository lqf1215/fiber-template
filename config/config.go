package config

var Config = struct {
	DB   DB     `yaml:"db"`
	Port string `yaml:"port"`
	Aws  Aws    `json:"aws"`
}{}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Aws struct {
	EndpointUrl     string `json:"endpoint-url"`
	Region          string `json:"region"`
	AccessKeyID     string `json:"access-key-id"`
	SecretAccessKey string `json:"secret-access-key"`
	Bucket          string `json:"bucket"`
}
