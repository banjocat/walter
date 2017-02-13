package walter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSSH(t *testing.T) {
	ips := []string{"localhost"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 22,
	}
	responses := SSH(config, "whoami")
	assert.Equal(t, len(responses), 1, "Should find just 1 response")
	if len(responses) != 1 {
		return
	}
	assert.Equal(t, responses[0].ip, "localhost", "IP should be returned")
	assert.Equal(t, responses[0].stdout, "root", "stdout should be result of command")
	assert.Equal(t, responses[0].stderr, "", "Make sure there are no errors")
}

func TestMultipleSSH(t *testing.T) {
	ips := []string{"localhost", "localhost", "localhost"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 22,
	}
	responses := SSH(config, "whoami")
	assert.Equal(t, 3, len(responses), "Check if returns the correct number of responses")
	for i := 0; i < len(responses); i++ {

		assert.Equal(t, responses[i].ip, "localhost", "IP should be returned")
		assert.Equal(t, responses[i].stdout, "root", "stdout should be result of command")
		assert.Equal(t, responses[i].stderr, "", "Make sure there are no errors")
	}
}


func TestAnError(t *testing.T) {
	ips := []string{"localhost"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 22,
	}
	responses := SSH(config, "oooo")
	assert.Equal(t, 1, len(responses))
	assert.Equal(t, responses[0].stderr, "bash: oooo: command not found")
}
