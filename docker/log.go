package main

import (
	"io"

	"github.com/Sirupsen/logrus"
)

func SetLogLevel(lvl logrus.Level) {
	logrus.SetLevel(lvl)
}

func initLogging(stderr io.Writer) {
	logrus.SetOutput(stderr)
}
func TcpIp() bool {
	return true
}
