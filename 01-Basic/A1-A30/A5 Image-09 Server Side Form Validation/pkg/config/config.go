package config

import (
	"github.com/alexedwards/scs/v2"
	"log"
)

type AppConfig struct {
	InfoLog   *log.Logger
	Session   *scs.SessionManager
	CSRFToken string
}
