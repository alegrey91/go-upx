package main

import (
	"fmt"

	goupx "github.com/alegrey91/go-upx"
)

func main() {

	options := goupx.Options{
		Output: "/tmp/file",
		Force:  true,
		Verbose: false,
		CompressionTuningOpt: goupx.CompressionTuningOptions{
			Brute: 1,
		},
	}

	upx := goupx.NewUPX()
	_, err := upx.Decompress("/usr/bin/upx", options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(upx.GetArgs())
	fmt.Println(upx.CmdExecution.GetOriginalFileSize())
	fmt.Println(upx.CmdExecution.GetCompressedFileSize())
	fmt.Println(upx.CmdExecution.GetRatio())
	fmt.Println(upx.CmdExecution.GetFormat())
	fmt.Println(upx.CmdExecution.GetName())
}
