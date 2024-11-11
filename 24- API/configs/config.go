package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var conf *config

// Config struct with private fields
type config struct {
	dbDriver     string `mapstructure:"DB_DRIVER"`
	dbHost       string `mapstructure:"DB_HOST"`
	dbPort       string `mapstructure:"DB_PORT"`
	dbUser       string `mapstructure:"DB_USER"`
	dbPass       string `mapstructure:"DB_PASS"`
	dbName       string `mapstructure:"DB_NAME"`
	httpPort     string `mapstructure:"HTTP_PORT"`
	jwtSecret    string `mapstructure:"JWT_SECRET"`
	jwtExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	tokenAuth    *jwtauth.JWTAuth
}

// NewConfig is a constructor to create a new Config instance
func NewConfig() *config {
	if conf == nil {
		loadConfig(".")
	}
	return conf
}

// Getters for the private fields
func (c *config) DBDriver() string {
	return c.dbDriver
}

func (c *config) DBHost() string {
	return c.dbHost
}

func (c *config) DBPort() string {
	return c.dbPort
}

func (c *config) DBUser() string {
	return c.dbUser
}

func (c *config) DBPass() string {
	return c.dbPass
}

func (c *config) DBName() string {
	return c.dbName
}

func (c *config) HTTPPort() string {
	return c.httpPort
}

// func (c *config) JWTSecret() string {
// 	return c.jwtSecret
// }

func (c *config) JWTExpiresIn() int {
	return c.jwtExpiresIn
}

func (c *config) TokenAuth() *jwtauth.JWTAuth {
	return c.tokenAuth
}

// loadConfig reads configuration using Viper
func loadConfig(path string) (*config, error) {
	viper.SetConfigName("app-config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	conf = &config{
		dbDriver:     viper.GetString("DB_DRIVER"),
		dbHost:       viper.GetString("DB_HOST"),
		dbPort:       viper.GetString("DB_PORT"),
		dbUser:       viper.GetString("DB_USER"),
		dbPass:       viper.GetString("DB_PASS"),
		dbName:       viper.GetString("DB_NAME"),
		httpPort:     viper.GetString("HTTP_PORT"),
		jwtSecret:    viper.GetString("JWT_SECRET"),
		jwtExpiresIn: viper.GetInt("JWT_EXPIRES_IN"),
	}

	conf.tokenAuth = jwtauth.New("HS256", []byte(conf.jwtSecret), nil)
	return conf, nil
}
