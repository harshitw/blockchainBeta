package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/harshitw/blockchainBeta/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10, // Height of block is index, if three blocks we have one genesis, height 2 or 3
		Nonce:     348432,
	}
	buf := &bytes.Buffer{} // we could encode binary and put a connection into it
	// we put in buffer, buffer is writer
	// we could stream the bytes of the block
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)

}

func TestBock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     48484,
		},
		Transactions: nil, //make([]Transaction, 1),
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))

	assert.Equal(t, b, bDecode)

	fmt.Printf("%+v", bDecode)
}
