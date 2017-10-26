package main

import (
	"encoding/json"
	"io"
)

func jsontidy(in io.Reader, out io.Writer) error {

	dec := json.NewDecoder(in)
	dec.UseNumber()

	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")

	var data interface{}
	for dec.More() {
		err := dec.Decode(&data)
		if err != nil {
			return err
		}

		err = enc.Encode(&data)
		if err != nil {
			return err
		}
	}
	return nil
}
