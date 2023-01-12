# go-upx

![go-upx](go-upx.jpg)

![Test](https://github.com/alegrey91/go-upx/actions/workflows/test.yml/badge.svg)  [![Go Reference](https://pkg.go.dev/badge/github.com/alegrey91/go-upx.svg)](https://pkg.go.dev/github.com/alegrey91/go-upx)

**go-upx** is a command-line wrapper for `upx` utility.

For package reference, visit [https://pkg.go.dev/github.com/alegrey91/go-upx](https://pkg.go.dev/github.com/alegrey91/go-upx).

## What is UPX

UPX is a free, secure, portable, extendable, high-performance executable packer for several executable formats. More info at [https://upx.github.io/](https://upx.github.io/).

## Install

To install the module, run the following command:

```sh
go get github.com/alegrey91/go-upx
```

## Example

Let's see how you can use `go-upx` to compress a file from a Go program:

```golang
import (
    ...
    "github.com/alegrey91/go-upx"
)

...

    // UPX options definition
    options := goupx.Options{
    	Output: "/tmp/file",
    	Force:  true,
    	Verbose: false,
    	CompressionTuningOpt: goupx.CompressionTuningOptions{
    		Brute: 1,
    	},
    }

    // command execution
    upx := goupx.NewUPX()
    _, err := upx.Compress("/path/to/file/to_be_compressed", 9, options)
    if err != nil {
    	fmt.Println(err)
    }

    // display the results
    fmt.Println(upx.CmdExecution.GetFormat())
    fmt.Println(upx.CmdExecution.GetOriginalFileSize())
    fmt.Println(upx.CmdExecution.GetCompressedFileSize())
    fmt.Println(upx.CmdExecution.GetRatio())
    fmt.Println(upx.CmdExecution.GetFormat())
    fmt.Println(upx.CmdExecution.GetName())

...
```

Here you can find more code examples here: [examples](https://github.com/alegrey91/go-upx/tree/main/examples)

## Test

To test the code locally, run the following command:

```sh
go test -v .
```

## License

**go-upx** is available under [MIT](https://github.com/alegrey91/go-upx/blob/main/LICENSE) license.