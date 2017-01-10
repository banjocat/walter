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
	assert.Equal(t, responses[0].ip, "jackmuratore.com", "IP should be returned")
	assert.Equal(t, responses[0].stdout, "root", "stdout should be result of command")
	assert.Equal(t, responses[0].stderr, "", "Make sure there are no errors")
	assert.Equal(t, responses[0].errorCode, 0, "Returns 0 if no error")
}

func TestMultipleSSH(t *testing.T) {
	ips := []string{"jackmuratore.com", "jackmuratore.com", "jackmuratore.com"}
	config := &Config{
		Pem:  "~/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 2222,
	}
	responses := SSH(config, "whoami")
	assert.Equal(t, 3, len(responses), "Check if returns the correct number of responses")
	for i := 0; i < len(responses); i++ {

		assert.Equal(t, responses[i].ip, "jackmuratore.com", "IP should be returned")
		assert.Equal(t, responses[i].stdout, "root", "stdout should be result of command")
		assert.Equal(t, responses[i].stderr, "", "Make sure there are no errors")
		assert.Equal(t, responses[i].errorCode, 0, "Returns 0 if no error")
	}
}
