package main

import (
	"bytes"
	"strings"
	"testing"
)

var tests = []struct {
	Name        string
	Input       string
	Expected    string
	ExpectError bool
}{
	{
		Name:  "Simple object",
		Input: `{"foo": "bar"}`,
		Expected: `{
  "foo": "bar"
}
`,
	},
	{
		Name:        "Invalid json",
		Input:       `this isn't JSON`,
		ExpectError: true,
	},
}

func TestJSONTidy(t *testing.T) {
	var output bytes.Buffer
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			output.Reset()
			err := jsontidy(strings.NewReader(tc.Input), &output)
			if tc.ExpectError && err == nil {
				t.Fatalf("expected an error, got none")
			}
			if !tc.ExpectError && err != nil {
				t.Fatalf("expected no err, got error %s", err.Error())
			}
			if output.String() != tc.Expected {
				t.Fatalf("got:\n%s\nwant:\n%s", output.String(), tc.Expected)
			}
		})

	}
}
