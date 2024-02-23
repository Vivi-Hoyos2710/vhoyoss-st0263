package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ListenIPServer    string
	ListenPortServer  string
	URLPeer           string
	IPCentralServer   string
	PortCentralServer string
}

// Bootstrap is an initial configuration function that gets called before the application starts.It takes environment variables and configures the application.
func Bootstrap(path string) (Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		fmt.Errorf("Error loading .env file: %w", err)
		return Config{}, err
	}
	config := Config{
		ListenIPServer:    os.Getenv("LISTEN_IP_SERVER"),
		ListenPortServer:  os.Getenv("LISTEN_PORT_SERVER"),
		URLPeer:           os.Getenv("URL_PEER"),
		IPCentralServer:   os.Getenv("IP_CENTRAL_SERVER"),
		PortCentralServer: os.Getenv("PORT_CENTRAL_SERVER"),
	}
	return config, nil
}

// que directorio voy a exportar?
