package main

import (
	"errors"
	"fmt"
	"hcc/horn/action/grpc/server"
	"hcc/horn/lib/config"
	"hcc/horn/lib/logger"
	"hcc/horn/lib/mysqlUtil"
	"hcc/horn/lib/pid"
	"hcc/horn/lib/rsautil"
	"hcc/horn/lib/syscheck"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func parseArgs(args []string) error {
	length := len(args)

	if length == 1 {
		return nil
	}

	if length == 2 {
		if args[1] == "genkey" {
			fmt.Println("Generating private and public keys...")

			err := rsautil.WritePrivateAndPublicKeys()
			if err != nil {
				return err
			}

			os.Exit(0)
		}
	}

	return errors.New("unknown args")
}

func init() {
	err := parseArgs(os.Args)
	if err != nil {
		fmt.Println("Usage: To running service, just type 'horn'. To generate keys, type 'horn genkey'.")
		panic(err)
	}

	err = syscheck.CheckRoot()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	hornRunning, hornPID, err := pid.IsHornRunning()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if hornRunning {
		fmt.Println("horn is already running. (PID: " + strconv.Itoa(hornPID) + ")")
		os.Exit(1)
	}
	err = pid.WriteHornPID()
	if err != nil {
		_ = pid.DeleteHornPID()
		fmt.Println(err)
		panic(err)
	}

	err = logger.Init()
	if err != nil {
		_ = pid.DeleteHornPID()
		panic(err)
	}

	config.Init()

	logger.Logger.Println("Changing mysqld password...")
	err = mysqlUtil.ChangePassword()
	if err != nil {
		_ = pid.DeleteHornPID()
		panic(err)
	}

	running, mysqldPID, err := pid.IsMySQLDRunning()
	if running {
		logger.Logger.Println("mysqld is running with pid=" + strconv.Itoa(mysqldPID))
	}
	if err != nil {
		logger.Logger.Panic("mysqld is not running! (" + err.Error() + ")")
	}
}

func end() {
	logger.End()
	_ = pid.DeleteHornPID()
}

func main() {
	// Catch the exit signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		end()
		fmt.Println("Exiting horn module...")
		os.Exit(0)
	}()

	server.Init()
}
