package walter

import (
	"fmt"
)

type Config struct {
	Pem  string
	Ips  []string
	Port int
	User string
}

func SSH(config *Config, command string) {
	fmt.Println("This is SSH!")
}
