package config

type mysqld struct {
	User string `goconf:"mysqld:user"` // User : mysqld user
	Host string `goconf:"mysqld:host"` // Address : mysqld host
}

// Mysqld : mysqld config structure
var Mysqld mysqld
