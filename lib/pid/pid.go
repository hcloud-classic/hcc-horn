package pid

import (
	"hcc/horn/lib/fileutil"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
)

var mysqldPIDFile = "/var/run/mysqld/mysqld.pid"

// IsMySQLDRunning : Check if mysqld is running
func IsMySQLDRunning() (running bool, pid int, err error) {
	if _, err := os.Stat(mysqldPIDFile); os.IsNotExist(err) {
		return false, 0, nil
	}

	pidStr, err := ioutil.ReadFile(mysqldPIDFile)
	if err != nil {
		return false, 0, err
	}

	mysqldPID, _ := strconv.Atoi(strings.TrimSpace(string(pidStr)))

	proc, err := os.FindProcess(mysqldPID)
	if err != nil {
		return false, 0, err
	}
	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true, mysqldPID, nil
	}

	return false, 0, nil
}

var hornPIDFileLocation = "/var/run"
var hornPIDFile = "/var/run/horn.pid"

// IsHornRunning : Check if horn is running
func IsHornRunning() (running bool, pid int, err error) {
	if _, err := os.Stat(hornPIDFile); os.IsNotExist(err) {
		return false, 0, nil
	}

	pidStr, err := ioutil.ReadFile(hornPIDFile)
	if err != nil {
		return false, 0, err
	}

	hornPID, _ := strconv.Atoi(string(pidStr))

	proc, err := os.FindProcess(hornPID)
	if err != nil {
		return false, 0, err
	}
	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true, hornPID, nil
	}

	return false, 0, nil
}

// WriteHornPID : Write horn PID to "/var/run/horn.pid"
func WriteHornPID() error {
	pid := os.Getpid()

	err := fileutil.CreateDirIfNotExist(hornPIDFileLocation)
	if err != nil {
		return err
	}

	err = fileutil.WriteFile(hornPIDFile, strconv.Itoa(pid))
	if err != nil {
		return err
	}

	return nil
}

// DeleteHornPID : Delete the horn PID file
func DeleteHornPID() error {
	err := fileutil.DeleteFile(hornPIDFile)
	if err != nil {
		return err
	}

	return nil
}
