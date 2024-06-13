package core

import "io"

// Make adapters so I can plug in etherum EVM into our own blockchain
// In that case data will be byte code/solidity code

type Transaction struct {
	Data []byte // Generic data not blockchain specific

}

func (tx *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}

func (tx *Transaction) EncodeBinary(r io.Writer) error {
	return nil
}
