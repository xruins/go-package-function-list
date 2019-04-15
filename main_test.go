package main

import (
	"testing"
)

const (
	testDataPath = "./testdata/hoge"
)

func TestDo(t *testing.T) {
	var cases = []struct {
		description string
		cmdOptions  *cmdOptions
		output      string
	}{
		{
			description: "parses directory without any options",
			cmdOptions: &cmdOptions{
				Dir:   testDataPath,
				Bound: " ",
			},
			output: "FooTest BarTest Foo Bar hoge fuga",
		},
		{
			description: "parses directory with public-only option",
			cmdOptions: &cmdOptions{
				Dir:        testDataPath,
				PublicOnly: true,
				Bound:      " ",
			},
			output: "FooTest BarTest Foo Bar",
		},
		{
			description: "parses directory with suffix option",
			cmdOptions: &cmdOptions{
				Dir:    testDataPath,
				Suffix: "Test",
				Bound:  " ",
			},
			output: "FooTest BarTest",
		},
		{
			description: "parses directory with regex option",
			cmdOptions: &cmdOptions{
				Dir:   testDataPath,
				Regex: "Test$",
				Bound: " ",
			},
			output: "FooTest BarTest",
		},
		{
			description: "parses directory with regex option and suffix option",
			cmdOptions: &cmdOptions{
				Dir:    testDataPath,
				Regex:  "^Foo",
				Suffix: "Test",
				Bound:  " ",
			},
			output: "FooTest",
		},
		{
			description: "parses directory with regex option and suffix option",
			cmdOptions: &cmdOptions{
				Dir:    testDataPath,
				Regex:  "^Foo",
				Suffix: "Test",
				Bound:  " ",
			},
			output: "FooTest",
		},
		{
			description: "parses directory with regex option and suffix option, public-only option",
			cmdOptions: &cmdOptions{
				Dir:        testDataPath,
				Regex:      "^[A-Za-z]",
				PublicOnly: true,
				Suffix:     "r",
				Bound:      " ",
			},
			output: "Bar",
		},
	}

	for _, c := range cases {
		out, err := do(c.cmdOptions)
		if err != nil {
			t.Fatalf("unexpected error occured, case: %s, err: %s", c.description, err)
		}
		if out != c.output {
			t.Errorf("unexpected outout, expected: %s, actual: %s", c.output, out)
		}
	}
}
