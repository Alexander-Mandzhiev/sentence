package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env        string         `yaml:"env" env:"ENV" env-default:"development"`
	HTTPServer HTTPServer     `yaml:"http_server"`
	Services   ServicesConfig `yaml:"services"`
	Frontend   Frontend       `yaml:"frontend"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"0.0.0.0"`       // Адрес сервера
	Port        int           `yaml:"port" env:"PORT" env-default:"6000"`                // Порт сервера
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`            // Таймаут запроса
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"` // Idle таймаут
}

type ServicesConfig struct {
	AttachmentTypesAddr      string `yaml:"attachment_types_addr" env:"ATTACHMENT_TYPES_ADDR" env-default:"0.0.0.0:7001"`
	DirectionsAddr           string `yaml:"directions_addr" env:"DIRECTIONS_ADDR" env-default:"0.0.0.0:7002"`
	HistoryAddr              string `yaml:"history_addr" env:"HISTORY_ADDR" env-default:"0.0.0.0:7003"`
	AttachmentsAddr          string `yaml:"attachments_addr" env:"ATTACHMENTS_ADDR" env-default:"0.0.0.0:7004"`
	DepartmentsAddr          string `yaml:"departments_addr" env:"DEPARTMENTS_ADDR" env-default:"0.0.0.0:7005"`
	ImplementorsAddr         string `yaml:"implementors_addr" env:"IMPLEMENTORS_ADDR" env-default:"0.0.0.0:7006"`
	PrioritiesAddr           string `yaml:"priorities_addr" env:"PRIORITIES_ADDR" env-default:"0.0.0.0:7007"`
	StatusesAddr             string `yaml:"statuses_addr" env:"STATUSES_ADDR" env-default:"0.0.0.0:7008"`
	SentencesAddr            string `yaml:"sentences_addr" env:"SENTENCES_ADDR" env-default:"0.0.0.0:7009"`
	SentencesAttachmentsAddr string `yaml:"sentences_attachments_addr" env:"SENTENCES_ATTACHMENTS_ADDR" env-default:"0.0.0.0:7010"`
}

type Frontend struct {
	Addr string `yaml:"addr" env:"FRONTEND_ADDR" env-default:"0.0.0.0:8080"`
}

func MustLoad() *Config {
	configPath := fetchConfigFlag()
	if configPath == "" {
		return loadingDataInEnv()
	}
	return MustLoadByPath(configPath)
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}

func fetchConfigFlag() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

func loadingDataInEnv() *Config {
	loadEnv()

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		log.Printf("Warning: Invalid PORT value in environment variables, using default value %d.", 6000)
		port = 6000
	}

	return &Config{
		HTTPServer: HTTPServer{
			Address:     os.Getenv("ADDRESS"),
			Port:        port,
			Timeout:     parseDuration(os.Getenv("TIMEOUT"), 5*time.Second),
			IdleTimeout: parseDuration(os.Getenv("IDLE_TIMEOUT"), 60*time.Second),
		},
		Services: ServicesConfig{
			AttachmentTypesAddr:      os.Getenv("ATTACHMENT_TYPES_ADDR"),
			DirectionsAddr:           os.Getenv("DIRECTIONS_ADDR"),
			HistoryAddr:              os.Getenv("HISTORY_ADDR"),
			AttachmentsAddr:          os.Getenv("ATTACHMENTS_ADDR"),
			DepartmentsAddr:          os.Getenv("DEPARTMENTS_ADDR"),
			ImplementorsAddr:         os.Getenv("IMPLEMENTORS_ADDR"),
			PrioritiesAddr:           os.Getenv("PRIORITIES_ADDR"),
			StatusesAddr:             os.Getenv("STATUSES_ADDR"),
			SentencesAddr:            os.Getenv("SENTENCES_ADDR"),
			SentencesAttachmentsAddr: os.Getenv("SENTENCES_ATTACHMENTS_ADDR"),
		},
		Frontend: Frontend{
			Addr: os.Getenv("FRONTEND_ADDR"),
		},
	}
}

func loadEnv() {
	if err := godotenv.Load(".gateway.env"); err != nil {
		log.Println("Warning: .gateway.env file not found, using default values.")
	}
}

func parseDuration(value string, defaultValue time.Duration) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil || duration <= 0 {
		log.Printf("Warning: Invalid TIMEOUT or IDLE_TIMEOUT value in environment variables, using default value %v.", defaultValue)
		return defaultValue
	}
	return duration
}
