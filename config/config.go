package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type AppConfig struct{
	AppPort string
}

type SecurityConfig struct {
	Key string
	Durasi time.Duration
	Issuer string
}

type Config struct {
	DbConfig
	AppConfig
	SecurityConfig
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error mang env na")
	}

	lifeTime, err := strconv.Atoi(os.Getenv("JWT_LIFE_TIME"))
	if err != nil {
		return err
	}
	
	c.SecurityConfig = SecurityConfig{
		Key: os.Getenv("JWT_KEY"),
		Durasi: time.Duration(lifeTime),
		Issuer: os.Getenv("JWT_ISSUER_NAME"),
	}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
	c.AppConfig = AppConfig {
		AppPort: os.Getenv("PORT_APP"),
	}
	
	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" || c.AppConfig.AppPort == ""|| c.SecurityConfig.Key == ""|| c.SecurityConfig.Durasi == 0 || c.SecurityConfig.Issuer == ""{
		return errors.New("environment is empty")
	}
	return nil
}

func NewConfig() (*Config, error){
	config := &Config{}
	if err := config.readConfig(); err != nil {
		return nil, err		
	}
	return config, nil
}
