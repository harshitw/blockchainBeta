package network

import (
	"fmt"
	"time"
)

// Multiple transport layers
type ServerOpts struct {
	Transports []Transport
}

// Container with multiple modules
type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}

	// It will also hold our transaction mem pool - every time trans comes it will be validated
	// It will also have generic blockchain handler that handles blocks and maybe we need to persist blocks
}

// Implementing first transport of our blockchain server object
func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	// For every transport we need to make them listen for messages
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	// keep looping and check
	for {
		select {
		// is there something to consume from rpc channel
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		// do we need to quit this
		case <-s.quitCh:
			break free
		// if there is nothing else then it freezes up
		// default will helps as it will go here and keep looping
		// But it's very cpu intensive,
		// default :
		case <-ticker.C: // Each 5 seconds
			fmt.Println("Do stuff every x seconds")
		}
	}

	fmt.Println("Server shutdown ")

}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		// start a go rountine for each transport
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				// Is the message going to be thread safe?
				// use locks OR pipe every rpc coming for go routines
				// from each transport directly into server its own rpc channel
				s.rpcCh <- rpc // pipeing
			}
		}(tr)
	}
}
