# go-plugin-cross-compile

A minimal example for cross-compiling go plugin apps

## Prerequisites

- go1.18.3

## Compile the plugin

> :bulb: This step is optional. Pre-compiled plugins for _linux/arm64_ & _linux/amd64_ target platforms are available in bin/ already.

Compile the plugin on the right platform and move it to the bin directory.

```bash
go build -buildmode plugin -trimpath -o bin/plugin-$(go env GOOS)-$(go env GOARCH).so plugin/main.go
```

## Compile and run the binary

Compile the binary and test it.

```bash
go build -trimpath -o bin/main entry/main.go
echo hello | ./bin/main
```

## Cross-compile the binary

Cross-compile the binary.

```bash
GOOS=linux GOARCH=arm64 go build -trimpath -o bin/main entry/main.go
```

## Issues when cross-compiling

Steps to reproduce:

1. Compile plugin on _linux/arm64_ (already located in bin/)
2. Cross-compile binary on _linux/amd64_ (or any other non-_linux/arm64_) targeting _linux/arm64_
3. Execute binary on _linux/arm64_
4. It raises the following error:

    ```bash
    $ echo hello | ./main
    error loading plugin "/home/ec2-user/tmp/testplugin/plugin-linux-arm64.so"
    panic: plugin: not implemented

    goroutine 1 [running]:
    main.main()
            ./main.go:30 +0x110
    ```
