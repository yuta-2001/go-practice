package functionaloptionspattern

import (
	"errors"
	"net/http"
)

type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == nil {
		port = 8080
	} else {
		if *options.port == 0 {
			port = 8080
		} else {
			port = *options.port
		}
	}

	return &http.Server{
		Addr: addr + ":" + string(port),
	}, nil
}

func main() {
	server, err := NewServer("localhost", WithPort(8081))
	if err != nil {
		panic(err)
	}
	server.ListenAndServe()
}
