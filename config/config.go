package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type TokenConfig struct {
	IssuerName       string `json:"IssuerName"`
	JwtSignatureKy   []byte `json:"JwtSignatureKy"`
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtExpiresTime   time.Duration
}

type Config struct {
	TokenConfig
}

func (c *Config) Configuration() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}
	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	c.TokenConfig = TokenConfig{
		IssuerName:       os.ExpandEnv("TOKEN_ISSUE"),
		JwtSignatureKy:   []byte(os.Getenv("TOKEN_SECRET")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtExpiresTime:   time.Duration(tokenExpire) * time.Minute,
	}

	if c.IssuerName == "" || c.JwtExpiresTime < 0 || len(c.JwtSignatureKy) == 0 {
		return fmt.Errorf("missing required environment")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.Configuration(); err != nil {
		return nil, err
	}
	return cfg, nil
}
