package config

type rsakey struct {
	PrivateKeyFile string `goconf:"rsakey:private_key_file"` // PrivateKeyFile : RSA private key file for decrypt mysqld password
	PublicKeyFile  string `goconf:"rsakey:public_key_file"`  // PublicKeyFile : RSA public key file for encrypt mysqld password
}

// Rsakey : rsakey config structure
var Rsakey rsakey
