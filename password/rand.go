package password

import (
	"bufio"
	crand "crypto/rand"
	"encoding/binary"
	"io"
	mrand "math/rand/v2"
)

type rand struct {
	*mrand.Rand
}

func newRand() *rand {
	return &rand{Rand: mrand.New(newSource(crand.Reader))} //nolint:gosec
}

// source wraps crypto/rand.Reader with bufio.Reader to minimize Read calls from /dev/urandom.
type source struct {
	r *bufio.Reader
}

func newSource(r io.Reader) *source {
	return &source{r: bufio.NewReader(r)}
}

func (s *source) Uint64() uint64 {
	var v uint64
	if err := binary.Read(s.r, binary.NativeEndian, &v); err != nil {
		panic(err)
	}
	return v
}
