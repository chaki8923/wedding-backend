package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"golang.org/x/xerrors"
)

type Config struct {
	AppDomain  string `env:"API_APP_DOMAIN" envDefault:""`
	Env        string `env:"API_ENV" envDefault:"dev"`
	Port       string `env:"API_PORT" envDefault:"8080"`
	ProFrontURL   string `env:"PRO_FRONT_URL" envDefault:"https://front.wedding-hackathon.com/"`
	FrontURL   string `env:"API_FRONT_URL" envDefault:"https://localhost:3443/"`
	LocalFrontURL   string `env:"API_FRONT_URL" envDefault:"http://localhost:3000/"`
	DBHost     string `env:"API_DB_HOST" envDefault:"172.30.0.3"`
	DBName     string `env:"API_DB_NAME" envDefault:"golang"`
	DBUser     string `env:"API_DB_USER" envDefault:"root"`
	DBPass     string `env:"API_DB_PASS" envDefault:"pass"`
	DBPort     string `env:"API_DB_PORT" envDefault:"3306"`
	EncryptKey string `env:"API_ENCRYPT_KEY" envDefault:"passwordpassword"`
	JwtSecret  string `env:"API_JWT_SECRET" envDefault:"secret"`
	SentryDsn  string `env:"API_SENTRY_DSN" envDefault:"https://xxxxxxxx"`
	AwsAccessKey  string `env:"AWS_ACCESS_KEY" envDefault:"aaaaaaa"`
	AwsSecretKey  string `env:"AWS_SECRET_KEY" envDefault:"bbbbbb"`
	GoogleAccount  string `env:"GOOGLE_ACCOUNT" envDefault:"bbbbbb"`
	GoogleApiKey  string `env:"GOOGLE_API_KEY" envDefault:"bbbbbb"`
}

func New() (*Config, error) {
	// .envファイルをロード
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerrors.Errorf("fail to parse cfg: %w", err)
	}
	return cfg, nil
}
