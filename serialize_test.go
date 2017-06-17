package config

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	var input struct {
		TestArray []string
	}

	input.TestArray = []string{"ahmdrz"}

	bytes, err := encodeGob(&input)
	if err != nil {
		t.Fatal(err)
	}

	err = decodeGob(bytes, &input)
	if err != nil {
		t.Fatal(err)
	}

	if len(input.TestArray) != 1 {
		t.Fail()
	}
	t.Log("serialize test : done")
}
