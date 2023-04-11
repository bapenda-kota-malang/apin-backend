package apicore

/*****************************************************************************\
The idea is to place all the needs for a web application to be running in one
place. All the user to do is set the configuration and the handlers since it
is main properties of a website.
\*****************************************************************************/

import (
	"net/http"

	"go.uber.org/zap"
	"gorm.io/gorm"
	// "github.com/karincake/nogosari/mailer"
)

type app struct {
	Code            string
	Name            string
	Env             string
	Version         string
	HttpConf        *httpConf
	DbConf          *dbConf
	MailerConf      *mailerConf
	RateLimiterConf *rateLimiterConf
	LoggerConf      *loggerConf
	BankJatimConf   *bankJatimConf
}

// Exported package vars for easier resolving
var (
	Self      *app // the core itself
	DB        *gorm.DB
	Logger    *zap.Logger
	BankJatim *BankJatimClient
)

// app starter
func Run(routerIn http.Handler, code string) {
	Self = &app{}

	// Intantiate all the conf
	Self.Code = code
	Self.HttpConf = &httpConf{}
	Self.DbConf = &dbConf{}
	Self.MailerConf = &mailerConf{}
	Self.LoggerConf = &loggerConf{}
	Self.BankJatimConf = &bankJatimConf{}

	// init all manually to make it clear of what's happening
	Self.initConfig()
	Self.initLogger()
	Self.initDb()
	Self.initBankJatimClient()
	// a.initMailer()
	Self.initExtCall()
	Self.initHttp(routerIn)
}
