package network

type NetAddr string

// <Message sent over transport layer>
type RPC struct {
	From    NetAddr
	Payload []byte
}

// Transport is module in server and server needs access to all messages
// sent over transport layer with consume method
type Transport interface {
	// returns channel of RPC
	Consume() <-chan RPC
	// connect with another transport
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
