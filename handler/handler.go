package handler

import (
	"digimon/logger"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry
)

func init() {
	log = logger.GetLogger().WithField("pkg", "handler")
}
