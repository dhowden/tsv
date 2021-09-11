package tsv_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dhowden/tsv"
)

func TestReader(t *testing.T) {
	tests := []struct {
		in  string
		out [][]byte
	}{
		{
			in:  "A",
			out: [][]byte{[]byte("A")},
		},
		{
			in:  "A\tB",
			out: [][]byte{[]byte("A"), []byte("B")},
		},
		{
			in:  "A\t\\N\tC",
			out: [][]byte{[]byte("A"), nil, []byte("C")},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			r := tsv.NewReader(strings.NewReader(test.in))
			out, err := r.ReadRow()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.out, out) {
				t.Errorf("expected %v, got %v", test.out, out)
			}
		})
	}
}

func TestReaderCustomNull(t *testing.T) {
	tests := []struct {
		in   string
		null []byte
		out  [][]byte
	}{
		{
			in:   "A\t\\N",
			null: nil,
			out:  [][]byte{[]byte("A"), []byte("\\N")},
		},
		{
			in:   "A\t\\N",
			null: []byte("A"),
			out:  [][]byte{nil, []byte("\\N")},
		},
		{
			in:   "A\t\\N",
			null: []byte("\t"),
			out:  [][]byte{[]byte("A"), []byte("\\N")},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			r := tsv.NewReader(strings.NewReader(test.in))
			r.Null = test.null

			out, err := r.ReadRow()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.out, out) {
				t.Errorf("expected %v, got %v", test.out, out)
			}
		})
	}
}
