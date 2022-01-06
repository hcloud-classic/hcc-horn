package mysqlUtil

import (
	"crypto/sha512"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"hcc/horn/lib/cmd"
	"hcc/horn/lib/config"
	"hcc/horn/lib/rsautil"
	"io/ioutil"
)

var generatedPassword string

func generatePassword() (string, error) {
	passGen, err := password.Generate(128, 32, 32, false, true)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha512.Sum512([]byte(passGen))), nil
}

func ChangePassword() error {
	var err error

	generatedPassword, err = generatePassword()
	if err != nil {
		return err
	}

	err = stopMYSQLD()
	if err != nil {
		return err
	}

	_ = stopMySQLDSafe()
	runMySQLDSafe()

	err = cmd.RunCMD("mysql -u" + config.Mysqld.User + " -e \"flush privileges;alter user '" +
		config.Mysqld.User + "'@'" + config.Mysqld.Host +
		"' identified with mysql_native_password by '" + generatedPassword + "';flush privileges;\"")
	if err != nil {
		return err
	}

	err = killMYSQLDSafe()
	if err != nil {
		return err
	}

	err = stopMySQLDSafe()
	if err != nil {
		return err
	}

	err = startMYSQLD()
	if err != nil {
		return err
	}

	return nil
}

func GetEncryptPassword() ([]byte, error) {
	pubKeyData, err := ioutil.ReadFile(config.Rsakey.PublicKeyFile)

	pubKey, err := rsautil.BytesToPublicKey(pubKeyData)
	if err != nil {
		return nil, err
	}

	encryptedData, err := rsautil.EncryptWithPublicKey([]byte(generatedPassword), pubKey)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}
