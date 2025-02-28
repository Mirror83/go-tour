package main

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	count, ok := rot13.r.Read(b)
	if ok != nil {
		return count, ok
	}

	for i, char := range b {
		if char >= 'A' && char <= 'Z' {
			b[i] = 'A' + (char-'A'+13)%26
		} else if char >= 'a' && char <= 'z' {
			b[i] = 'a' + (char-'a'+13)%26
		} else {
			b[i] = char
		}
	}

	return count, nil
}
