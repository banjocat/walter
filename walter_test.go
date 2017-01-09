package walter

import (
	"github.com/stretchr/testify/assert"
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
	responses := SSH(config, "whoami")
	assert.Equal(t, responses[0].ip, "jackmuratore.com", "Should be jackmuratore.com")
	assert.Equal(t, responses[0].stdout, "root", "Should output root")
	assert.Equal(t, responses[0].stderr, "", "There should be no error")
	assert.Equal(t, responses[0].errorCode, 0, "Should exit with 0")

}
