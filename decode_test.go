package tsv_test

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/dhowden/tsv"
)

type errorReader struct {
	err error
}

func (e errorReader) Read(_ []byte) (int, error) {
	return 0, e.err
}

func TestDecodeErrors(t *testing.T) {
	v, err := tsv.Decode(strings.NewReader(""))
	if err == nil {
		t.Errorf("expected error, got %v", v)
	}

	v, err = tsv.Decode(strings.NewReader("A\t\\N\tB"))
	if err == nil {
		t.Errorf("expected error, got %v", v)
	}

	v, err = tsv.Decode(strings.NewReader("A\tB\tC\n1\n"))
	if err == nil {
		t.Errorf("expected error, got %v", v)
	}

	v, err = tsv.Decode(
		io.MultiReader(strings.NewReader("A\tB\n"), errorReader{
			err: errors.New("something stupid"),
		}),
	)
	if err == nil {
		t.Errorf("expected error, got %v", v)
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		in  string
		out []map[string]string
	}{
		{
			in:  "A",
			out: nil,
		},
		{
			in:  "A\tB",
			out: nil,
		},
		{
			in: "A\nA",
			out: []map[string]string{
				map[string]string{"A": "A"},
			},
		},
		{
			in: "A\tB\nA\tB",
			out: []map[string]string{
				map[string]string{"A": "A", "B": "B"},
			},
		},
		{
			in: "A\tB\n\\N\tB",
			out: []map[string]string{
				map[string]string{"B": "B"},
			},
		},
		{
			in: "A\tB\n\\N\t\\N",
			out: []map[string]string{
				map[string]string{},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			out, err := tsv.Decode(strings.NewReader(test.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(test.out, out) {
				t.Errorf("expected %v, got %v", test.out, out)
			}
		})
	}
}
