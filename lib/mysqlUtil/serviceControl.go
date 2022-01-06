package mysqlUtil

import "hcc/horn/lib/cmd"

func stopMYSQLD() error {
	err := cmd.RunCMD("service mysql stop")
	if err != nil {
		return err

	}

	return nil
}

func startMYSQLD() error {
	err := cmd.RunCMD("service mysql start")
	if err != nil {
		return err

	}

	return nil
}
