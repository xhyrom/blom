package reader

import (
	"errors"
	"io"
)

type Reader struct {
	runes  []rune
	length int
	cursor int
}

func New(content string) *Reader {
	runes := []rune(content)

	return &Reader{
		runes:  runes,
		length: len(runes),
		cursor: 0,
	}
}

func (r *Reader) Read() (rune, error) {
	if r.cursor == r.length {
		return 0, io.EOF
	}

	rune := r.runes[r.cursor]
	r.cursor++

	return rune, nil
}

func (r *Reader) Peek() (rune, error) {
	if r.cursor == r.length {
		return 0, io.EOF
	}

	return r.runes[r.cursor], nil
}

func (r *Reader) Current() (rune, error) {
	if r.cursor < 1 {
		return 0, io.EOF
	}

	return r.runes[r.cursor-1], nil
}

func (r *Reader) Rewind() error {
	r.cursor--

	if r.cursor < 0 {
		return errors.New("cannot rewind past the beginning of the reader")
	}

	return nil
}
