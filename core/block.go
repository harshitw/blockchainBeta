package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/harshitw/blockchainBeta/types"
)

type Header struct {
	Version   uint32 // unsigned integer
	PrevBlock types.Hash
	Timestamp int64 // unix nano
	Height    uint32
	Nonce     uint64 // number used once - ensures same operation is
	// never repeated with same nonce within a context, adds randomness
	// to operations so replay attacks are prevented
}

// Encode header to byte slice, instead of returning a byte slice
// we are passing writer and this method is super generic and extendible
func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, &h.Version); err != nil {
		return err
	} // LittleEndian
	if err := binary.Write(w, binary.BigEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, &h.Height); err != nil {
		return err
	}
	return binary.Write(w, binary.BigEndian, &h.Nonce)
}

// Decode to header from byte slice
func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, &h.Version); err != nil {
		return err
	} // LittleEndian
	if err := binary.Read(r, binary.BigEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &h.Height); err != nil {
		return err
	}
	return binary.Read(r, binary.BigEndian, &h.Nonce)
}

// Block consists of header, if we want to hash the block
// we cannot hash everything and we want to space
// we want to keep track of blocks but don't want to keep
// track of transactions
type Block struct {
	Header       Header
	Transactions []Transaction

	// cached version of the block hash
	hash types.Hash // This field is used to improve performance
}

// If we want to sent this block over the network we have to use byte slice

// We want to cash the hash of block as it will be costly
func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{} // make buffer
	b.Header.EncodeBinary(buf)
	// use buffer to use it as a has

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	return b.hash
}

// Unmarshalling and Marshalling
func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil
}
