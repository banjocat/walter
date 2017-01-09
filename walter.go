package walter

type Config struct {
	Pem  string
	Ips  []string
	Port int
	User string
}

type Response struct {
	ip        string
	errorCode int
	stderr    string
	stdout    string
}

func SSH(config *Config, command string) (responses []*Response) {
	response := &Response{
		ip:        "fixme",
		errorCode: 130,
		stderr:    "fixme",
		stdout:    "fixme",
	}
	return []*Response{response}
}
