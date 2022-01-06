package config

type rsakey struct {
	PublicKeyFile string `goconf:"rsakey:public_key_file"` // PublicKeyFile : RSA public key file for encrypt mysqld password
}

// Rsakey : rsakey config structure
var Rsakey rsakey
