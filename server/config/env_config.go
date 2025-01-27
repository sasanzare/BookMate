package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"github.com/joho/godotenv" 
)

var env ConfigDto

func init() {
	LoadEnvironmentVariable() 
	ConfigEnv()              
}

type ConfigDto struct {
	Port        string
	SecretKey   string
	DatabaseUrl string
}

func ConfigEnv() {
	env = ConfigDto{
		Port:        os.Getenv("PORT"),
		SecretKey:   os.Getenv("SECRET_KEY"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}

func LoadEnvironmentVariable() {
	err := godotenv.Load() 
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}

func AccessField(key string) (string, error) {
	v := reflect.ValueOf(env)
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected struct")
	}

	_, ok := t.FieldByName(key)
	if !ok {
		return "", fmt.Errorf("property '%s' could not be found", key)
	}

	f := v.FieldByName(key)
	if !f.IsValid() || f.Kind() != reflect.String {
		return "", fmt.Errorf("invalid field type or field '%s' not found", key)
	}

	return f.String(), nil
}

func GetEnvProperty(key string) (string, error) {
	if env.Port == "" {
		ConfigEnv()
	}
	return AccessField(key)
}