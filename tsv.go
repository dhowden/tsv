package tsv // import "github.com/dhowden/tsv"

import (
	"bufio"
	"bytes"
	"io"
)

// NewReader creates a new Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		Null: []byte{'\\', 'N'},
		Tab:  []byte{'\t'},

		s: bufio.NewScanner(r),
	}
}

// Reader reads records from a tsv file.  Must be created by
// NewReader, but properties can be set before use.
type Reader struct {
	// Null is the value denoting an empty column.
	// It is set to \N by NewReader.
	Null []byte

	// Tab is the tab character used to denote a column boundary.
	Tab []byte

	s *bufio.Scanner
}

// ReadRow extracts a single row from the tsv file, returning nil, io.EOF
// if no more records are available.  The underlying arrays returned will
// be invalidated on subsequent calls to ReadRow.
func (r *Reader) ReadRow() ([][]byte, error) {
	if !r.s.Scan() {
		return nil, io.EOF
	}

	b := r.s.Bytes()
	cols := bytes.Split(b, r.Tab)

	if r.Null == nil {
		return cols, nil
	}
	for i, c := range cols {
		if bytes.Equal(r.Null, c) {
			cols[i] = nil
		}
	}
	return cols, nil
}
