package config

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	block       cipher.Block
	name        string
	configPlace string
}

func New(key []byte, configName string) (*Config, error) {
	configPlace := os.Getenv("HOME") + "/.config/" + configName + "/"
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if runtime.GOOS == "windows" {
		configPlace = fmt.Sprintf("%s\\%s\\", os.Getenv("APPDATA"), configName)
	}

	if err := os.MkdirAll(filepath.Dir(configPlace), os.ModePerm); err != nil {
		return nil, err
	}

	return &Config{block, configName, configPlace}, nil
}

func (c *Config) Save(fileName string, object interface{}) error {
	var err error
	var bytes []byte

	bytes, err = encodeGob(object)
	if err != nil {
		return err
	}

	bytes, err = c.encryptAES(bytes)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.configPlace+fileName, bytes, os.ModePerm)
}

func (c *Config) Load(fileName string, object interface{}) error {
	bytes, err := ioutil.ReadFile(c.configPlace + fileName)
	if err != nil {
		return err
	}
	bytes, err = c.decryptAES(bytes)
	if err != nil {
		return err
	}
	return decodeGob(bytes, object)
}

func (c *Config) Remove() error {
	return os.RemoveAll(c.configPlace)
}
