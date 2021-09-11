package tsv // import "github.com/dhowden/tsv"

import (
	"fmt"
	"io"
)

// Decode reads the values from the TSV encoded data provided by the
// reader.
func Decode(r io.Reader) ([]map[string]string, error) {
	tr := NewReader(r)

	bs, err := tr.ReadRow()
	if err != nil {
		return nil, fmt.Errorf("could not read title row: %v", err)
	}

	titles := make([]string, 0, len(bs))
	for _, b := range bs {
		if b == nil {
			return nil, fmt.Errorf("titles must be non-null")
		}
		titles = append(titles, string(b))
	}

	var values []map[string]string
	for {
		bs, err := tr.ReadRow()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(bs) != len(titles) {
			return nil, fmt.Errorf("invalid number of columns %d (expected %d)", len(bs), len(titles))
		}

		m := make(map[string]string, len(titles))
		for i, title := range titles {
			if bs[i] != nil {
				m[title] = string(bs[i])
			}
		}
		values = append(values, m)
	}
	return values, nil
}
