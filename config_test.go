package config

import (
	"testing"
)

var key = []byte("E3A15379B331C825")

func TestConfig(t *testing.T) {
	var input struct {
		TestString string
	}

	input.TestString = "HelloWorld"
	c, err := New(key, "config_test")
	if err != nil {
		t.Fatal(err)
	}
	err = c.Save("test.config", &input)
	if err != nil {
		t.Fatal(err)
	}

	err = c.Load("test.config", &input)
	if err != nil {
		t.Fatal(err)
	}

	if input.TestString != "HelloWorld" {
		t.Fail()
	}

	err = c.Remove()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("config test : done")
}
