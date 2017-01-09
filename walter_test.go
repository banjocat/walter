package walter

import (
	"testing"
)

func TestSSH(t *testing.T) {
	ips := []string{"jackmuratore.com"}
	config := &Config{
		Pem:  "~/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 2222,
	}
	SSH(config, "uptime")
}
