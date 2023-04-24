package util

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

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
	defer fmt.Printf("%s took %v ms\n", label, time.Since(start).Milliseconds())
	fn()
}
