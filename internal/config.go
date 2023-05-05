package internal

type Config struct {
	Httpd    Httpd    `yaml:"httpd"`
	RabbitMQ RabbitMQ `yaml:"rabbitMQ"`
	Services Services `yaml:"services"`
}
type Httpd struct {
	Port string `yaml:"port"`
}
type RabbitMQ struct {
	ConnectionURL string `yaml:"connectionUrl"`
}
type AuthenticationService struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
type Services struct {
	AuthenticationService AuthenticationService `yaml:"authenticationService"`
}
