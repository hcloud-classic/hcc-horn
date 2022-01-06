package rsautil

import (
	"hcc/horn/lib/config"
	"io/ioutil"
)

func CheckKeyFiles() error {
	_, err := ioutil.ReadFile(config.Rsakey.PublicKeyFile)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadFile(config.Rsakey.PrivateKeyFile)
	if err != nil {
		return err
	}

	return nil
}
