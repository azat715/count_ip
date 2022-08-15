package readfile

import (
	"bufio"
	"os"
)

type reader struct {
	*bufio.Scanner
}

func newReader(f *os.File) *reader {
	return &reader{
		bufio.NewScanner(f),
	}
}

type TextFile struct {
	f *os.File
	r *reader
}

func (t *TextFile) Close() error {
	return t.f.Close()
}

func (t *TextFile) Bytes() []byte {
	return t.r.Bytes()
}

func (t *TextFile) Scan() bool {
	return t.r.Scan()
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
