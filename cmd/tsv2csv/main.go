package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/dhowden/tsv"
)

var (
	input  = flag.String("input", "", "input TSV file")
	output = flag.String("output", "", "output CSV file")
)

func main() {
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer f.Close()

	fout, err := os.Create(*output)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	defer fout.Close()

	cw := csv.NewWriter(fout)
	defer cw.Flush()

	tr := tsv.NewReader(f)
	for {
		bs, err := tr.ReadRow()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not read row: %v", err)
		}

		out := make([]string, len(bs))
		for i := range bs {
			out[i] = string(bs[i])
		}
		cw.Write(out)
	}
}
