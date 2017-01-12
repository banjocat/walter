/*
Package walter implements a simple way to ssh multiple hosts at once
*/
package walter

import (
    "golang.org/x/crypto/ssh"
    "io/ioutil"
    "bytes"
    "sync"
    "log"
    "strings"
    "fmt"
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
    stderr    string
    stdout    string
}

/*
SSH is the main endpoint to send ssh commands
*/
func SSH(config *Config, command string) ([]*Response) {
    clientConfig, err := walterConfigToCyrptoConfig(config)
    if err != nil {
	return nil
    }
    var wg sync.WaitGroup
    responses := make(chan *Response, len(config.Ips))
    for _, ip := range config.Ips {
	wg.Add(1)
	go runOneSSH(clientConfig, ip, config.Port, command, responses, &wg)
    }
    wg.Wait()
    close(responses)
    var response_slice []*Response
    for elem := range responses {
	response_slice = append(response_slice, elem)
    }
    return response_slice
}

func runOneSSH(clientConfig *ssh.ClientConfig, host string, port int, command string, responses chan <- *Response, wg *sync.WaitGroup) {
    defer wg.Done()
    hostname := fmt.Sprintf("%s:%d", host, port)
    client, err := ssh.Dial("tcp", hostname, clientConfig)
    if err != nil {
	log.Fatalf("Fatal client: %s", err)
	return
    }
    session, err := client.NewSession()
    if err != nil {
	log.Fatalf("Fatal session: %s", err)
	return
    }
    defer session.Close()
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    session.Stdout = &stdout
    session.Stderr = &stderr
    if err := session.Run(command); err != nil {
	log.Fatalf("Fatal run: %s", err)
	return
    }
    response := &Response {
	ip: host,
	stderr: strings.TrimSpace(stderr.String()),
	stdout: strings.TrimSpace(stdout.String()),
    }
    responses <- response
}

func walterConfigToCyrptoConfig(config *Config) (*ssh.ClientConfig, error) {
    key, err := ioutil.ReadFile(config.Pem)
    if err != nil {
	log.Fatal("Fatal: %s", err)
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
    return clientConfig, nil
}

