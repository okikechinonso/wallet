package ports

import log "github.com/sirupsen/logrus"

type ILogger interface {
	MakeLogger(filename string, display bool) *log.Logger
	SetFormater()
}
