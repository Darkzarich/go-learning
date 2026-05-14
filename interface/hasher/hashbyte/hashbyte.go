package hashbyte

import "io"

type Hasher interface {
	// Compose interface
	io.Writer
	Hash() byte
}

type hash struct {
	result byte
}

func New(_init byte) Hasher {
	return &hash{
		result: _init,
	}
}

// Write — here we can pass an array of bytes of any length, hash each byte and update the hash value
func (h *hash) Write(bytes []byte) (n int, err error) {
	// going over each byte of the array and updating the hash value
	// so hash.result is just one value
	for _, b := range bytes {
		h.result = (h.result^b)<<1 + b%2
	}
	return len(bytes), nil
}

func (h hash) Hash() byte {
	return h.result
}
