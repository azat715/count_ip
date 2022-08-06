package readfile

import (
	"bufio"
	"os"
	"strings"
)

const separator rune = '\n'

type reader struct {
	*bufio.Reader
}

func (r *reader) Readline() (string, error) {
	line, err := r.ReadString(byte(separator))
	if err != nil {
		return line, err
	}
	return strings.TrimSuffix(line, string(separator)), nil
}

func newReader(f *os.File) *reader {
	return &reader{
		bufio.NewReader(f),
	}
}

type TextFile struct {
	f *os.File
	*reader
}

func (t *TextFile) Close() error {
	return t.f.Close()
}

func New(path string) (*TextFile, func() error, error) {
	f, err := os.Open(path)
	if err != nil {
		return &TextFile{}, nil, err
	}
	t := TextFile{
		f,
		newReader(f),
	}
	return &t, t.Close, nil
}
