package main

import (
	"time"

	"github.com/harshitw/blockchainBeta/network"
)

// server - container with modules
// Transport layer - tcp/udp/websockets - maybe start with localTransport
// Block
// Transaction
// Keypairs

func main() {

	// We have our own machine (local node), transport for our own server
	// We have peers - remote peers/servers in the network
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello another node!"))
			time.Sleep(1 * time.Second)
		}

	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()

}
