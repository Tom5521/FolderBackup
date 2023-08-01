package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Tom5521/VSCodeBackup/src"
)

var sh = src.Sh{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enought arguments")
		return
	}
	if os.Args[1] == "test" {
		fmt.Println(runtime.GOOS)
		fmt.Println(sh.Getprefix())
		return
	}
	src.Rclone()
}
