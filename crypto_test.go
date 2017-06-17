package config

import (
	"testing"
)

func TestCrypto(t *testing.T) {
	c, err := New(key, "crypto_test")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := c.encryptAES([]byte("helloworld"))
	if err != nil {
		t.Fatal(err)
	}
	bytes, err = c.decryptAES(bytes)
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != "helloworld" {
		t.Fail()
	}

	err = c.Remove()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("crypto test : done")
}
