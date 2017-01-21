package walter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSSHwithIP(t *testing.T) {
	ips := []string{"107.170.57.50"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 2222,
	}
	responses := SSH(config, "whoami")
	assert.Equal(t, len(responses), 1, "Should find just 1 response")
	if len(responses) != 1 {
		return
	}
	assert.Equal(t, responses[0].ip, "107.170.57.50", "IP should be returned")
	assert.Equal(t, responses[0].stdout, "root", "stdout should be result of command")
	assert.Equal(t, responses[0].stderr, "", "Make sure there are no errors")
}

func TestMultipleSSHwithIP(t *testing.T) {
	ips := []string{"107.170.57.50", "107.170.57.50", "107.170.57.50"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 2222,
	}
	responses := SSH(config, "whoami")
	assert.Equal(t, 3, len(responses), "Check if returns the correct number of responses")
	for i := 0; i < len(responses); i++ {

		assert.Equal(t, responses[i].ip, "107.170.57.50", "IP should be returned")
		assert.Equal(t, responses[i].stdout, "root", "stdout should be result of command")
		assert.Equal(t, responses[i].stderr, "", "Make sure there are no errors")
	}
}

func TestSSHwithDNS(t *testing.T) {
	ips := []string{"jackmuratore.com"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 2222,
	}
	responses := SSH(config, "whoami")
	assert.Equal(t, len(responses), 1, "Should find just 1 response")
	if len(responses) != 1 {
		return
	}
	assert.Equal(t, responses[0].ip, "jackmuratore.com", "IP should be returned")
	assert.Equal(t, responses[0].stdout, "root", "stdout should be result of command")
	assert.Equal(t, responses[0].stderr, "", "Make sure there are no errors")
}

func TestMultipleSSHwithDNS(t *testing.T) {
	ips := []string{"jackmuratore.com", "jackmuratore.com", "jackmuratore.com"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
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
	}
}

func TestAnError(t *testing.T) {
	ips := []string{"107.170.57.50"}
	config := &Config{
		Pem:  "/home/banjocat/.ssh/id_rsa",
		Ips:  ips,
		User: "root",
		Port: 2222,
	}
	responses := SSH(config, "oooo")
	assert.Equal(t, 1, len(responses))
	assert.Equal(t, responses[0].stderr, "oooo: command not found")
}
