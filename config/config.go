package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HTTPServer httpServer
	DB         DB
	Redis      redis
	Secret     secret
}

type httpServer struct {
	Host string `envconfig:"HTTP_SERVER_HOST" default:"localhost" firestore:"http_server_host"`
	Port int    `envconfig:"HTTP_SERVER_PORT" default:"8080" firestore:"port"`
}
type DB struct {
	Host     string `envconfig:"DB_HOST" default:"localhost" firestore:"db_host"`
	Port     int    `envconfig:"DB_PORT" default:"8080" firestore:"db_port"`
	Username string `envconfig:"DB_USERNAME" default:"user" firestore:"db_user"`
	Password string `envconfig:"DB_PASSWORD" default:"pass" firestore:"password"`
	DbName   string `envconfig:"DB_NAME" default:"db" firestore:"db_name"`
}
type redis struct {
	Host     string `envconfig:"REDIS_HOST" default:"localhost" firestore:"redis_host"`
	Port     int    `envconfig:"REDIS_PORT" default:"6379" firestore:"redis_port"`
	DB       int    `envconfig:"REDIS_DB" default:"0" firestore:"redis_db"`
	Username string `envconfig:"REDIS_USERNAME" firestore:"redis_username"`
	Password string `envconfig:"REDIS_PASSWORD" firestore:"redis_password"`
}

type secret struct {
	PasswordSecret []byte `envconfig:"PASSWORD_SECRET"`
	JWTSecret      []byte `envconfig:"JWT_SECRET"`
}

func Init() config {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		godotenv.Load("../.env")
	}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
	return cfg
}
