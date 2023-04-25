package util

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func WriteMsg(w io.Writer, b []byte) error {
	var h = make([]byte, 4)

	binary.LittleEndian.PutUint32(h, uint32(len(b)))
	if _, err := w.Write(h); err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}

func ReadMsg(r io.Reader) ([]byte, error) {
	var h = make([]byte, 4)

	if _, err := r.Read(h); err != nil {
		return nil, err
	}
	l := binary.LittleEndian.Uint32(h)
	b := make([]byte, l)
	if _, err := r.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func EvalLatency(label string, fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("%s took %v ms\n", label, time.Since(start).Milliseconds())
}
