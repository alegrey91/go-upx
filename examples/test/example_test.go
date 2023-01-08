package main

import (
	"fmt"

	goupx "github.com/alegrey91/go-upx"
)

func main() {

	upx := goupx.NewUPX()
	_, err := upx.TestCompressedFile("/usr/bin/upx")
	if err != nil {
		fmt.Println(err)
	}
}
