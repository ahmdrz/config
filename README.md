## ahmdrz/config
> In-App simple configuration manager with AES encryption for Go applications.

## Usage

```go
    // Create configuration manager
    // with AES key and application name
    cfg, err := config.New([]byte("E3A15379B331C825"), "my_application")
    if err != nil {
	// handle error , occurred because key size is invalid.
    }

    // Let's save a struct , for example testStruct
    err = cfg.Save("test.cfg", testStruct)
    if err != nil {
    	panic(err)
    }

    // Now take a look to your .config or %APPDATA% ...
    // you will see a directory with your application name.

    // It's time to load our encrypted configuration into testStruct
    err = cfg.Load("test.cfg", &testStruct)
    if err != nil {
	panic(err)
    }

    // Do you want to remove all of configurations ?
    err = cfg.Remove()
```

## TODO

- [ ] Create another New method for RSA with public/private key.
- [ ] Using alexflint/go-memdump instead of GOB encoding method.
