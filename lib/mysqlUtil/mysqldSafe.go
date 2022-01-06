package mysqlUtil

import (
	"hcc/horn/lib/cmd"
	"hcc/horn/lib/logger"
	"hcc/horn/lib/pid"
	"strconv"
	"time"
)

var mysqldWaitTimeSec time.Duration = 30

func runMySQLDSafe() {
	logger.Logger.Println("Running mysqld_safe...")

	go func() {
		err := cmd.RunCMD("mysqld_safe --skip-grant-tables &")
		if err != nil {
			panic(err)
		}
	}()

	logger.Logger.Println("Wait " + strconv.Itoa(int(mysqldWaitTimeSec)) + " seconds until mysqld is up")
	var tick = 0
	for true {
		time.Sleep(time.Second)

		running, mysqldPID, err := pid.IsMySQLDRunning()
		if running {
			logger.Logger.Println("running mysqld in safe mode with pid=" + strconv.Itoa(mysqldPID))
			break
		}
		if err != nil {
			logger.Logger.Println("Error occurred while checking status of mysqld. (" + err.Error() + ")")
		}

		tick++
		if tick == int(mysqldWaitTimeSec) {
			panic("Failed to run mysqld_safe")
		}
	}
}

func killMYSQLDSafe() error {
	err := cmd.RunCMD("killall mysqld_safe")
	if err != nil {
		return err

	}

	return nil
}

func killMYSQLD() error {
	err := cmd.RunCMD("killall mysqld")
	if err != nil {
		return err

	}

	return nil
}

func stopMySQLDSafe() error {
	err := killMYSQLDSafe()
	if err != nil {
		return err
	}

	err = killMYSQLD()
	if err != nil {
		return err
	}

	logger.Logger.Println("Wait " + strconv.Itoa(int(mysqldWaitTimeSec)) + " seconds until mysqld is down")
	var tick = 0
	for true {
		time.Sleep(time.Second)

		running, _, err := pid.IsMySQLDRunning()
		if !running {
			break
		}
		if err != nil {
			logger.Logger.Println("Error occurred while checking status of mysqld. (" + err.Error() + ")")
		}

		tick++
		if tick == int(mysqldWaitTimeSec) {
			panic("Failed to stop mysqld safe mode")
		}
	}

	return nil
}
