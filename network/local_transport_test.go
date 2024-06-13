package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Normally use SLS third party packages
func TestConnect(t *testing.T) {
	//assert.Equal(t, 1, 1)
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.peers[trb.Addr()], trb)
	assert.Equal(t, trb.peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Hello another!")
	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	// Instead of method use function in real world, local test file could access the price
	// non public variable. Outside world will have no aceess to consumeCh as its private
	//We need to consume the channel sending from tra to trb
	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.Addr())

}
