package utils

import "github.com/wonderivan/logger"

func ProcessError(err error) error {
	//_ = os.Stderr.Close()
	//os.Exit(0)
	logger.Error(err.Error())
	return err
}