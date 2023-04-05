package config

import (
	"time"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// PublicRootPath is application controller path for public operations
const PublicRootPath = "/v1/public/knative"

// DefaultRequestTimeout is default timeout for internal http calls
const DefaultRequestTimeout = 30 * time.Second

type args struct {
	Port     string `arg:"env:PORT"`
	LogLevel string `arg:"env:LOG_LEVEL"`
	Data     string `arg:"env:ACS_CLIENT_URL"`
}

// Props is global variable for environment variables usage
var Props args

var opts struct {
	Profile string `arg:"-p" default:"default"`
}

// LoadConfig loads configuration when project run
func LoadConfig() {
	arg.MustParse(&opts)
	initLogger()
	log.Info("Application is starting with profile:", opts.Profile)

	initEnvVars()
	_ = arg.Parse(&Props)
	applyLoggerLevel()
}

func initEnvVars() {
	if godotenv.Load("config/profiles/default.env") != nil {
		log.Fatal("Error in loading environment variables from: profiles/default.env")
	} else {
		log.Info("Environment variables loaded from: profiles/default.env")
	}

	if opts.Profile != "default" {
		profileFileName := "config/profiles/" + opts.Profile + ".env"
		if godotenv.Overload(profileFileName) != nil {
			log.Fatal("Error in loading environment variables from:", profileFileName)
		} else {
			log.Info("Environment variables overloaded from:", profileFileName)
		}
	}
}

func initLogger() {
	log.SetLevel(log.InfoLevel)
	if opts.Profile == "default" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func applyLoggerLevel() {
	loglevel, err := log.ParseLevel(Props.LogLevel)
	if err != nil {
		loglevel = log.InfoLevel
	}
	log.SetLevel(loglevel)
}
