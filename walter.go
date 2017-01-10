/*
Package walter implements a simple way to ssh multiple hosts at once
*/
package walter

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

/*
Config is the struct used to setup an SSH request
*/
type Config struct {
	Pem  string
	Ips  []string
	Port int
	User string
}

/*
Response is the slice returned from SSH
*/
type Response struct {
	ip        string
	errorCode int
	stderr    string
	stdout    string
}

/*
SSH is the main endpoint to send ssh commands
*/
func SSH(config *Config, command string) (responses []*Response) {
	response := &Response{
		ip:        "fixme",
		errorCode: 130,
		stderr:    "fixme",
		stdout:    "fixme",
	}
	return []*Response{response}
}

func walterConfigToCyrptoConfig(config *Config) (*ssh.ClientConfig, error) {
	key, err := ioutil.ReadFile(config.Pem)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	clientConfig := &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	return clientConfig, err
}
