package testutils

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func NullLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)

	return logger
}
