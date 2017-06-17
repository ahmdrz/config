package main

import (
	"fmt"

	"github.com/ahmdrz/config"
)

func main() {
	var testStruct struct {
		Arr []string
	}

	testStruct.Arr = []string{"ahmdrz"}

	cfg, err := config.New([]byte("E3A15379B331C825"), "save_load")
	if err != nil {
		panic(err)
	}
	err = cfg.Save("test.cfg", testStruct)
	if err != nil {
		panic(err)
	}

	err = cfg.Load("test.cfg", &testStruct)
	if err != nil {
		panic(err)
	}

	fmt.Println(testStruct)

	err = cfg.Remove()
	if err != nil {
		panic(err)
	}
}
