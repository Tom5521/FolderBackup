package main

import (
	"os"

	"github.com/Tom5521/VSCodeBackup/src"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "restore":
		src.Rclone("restore")
	case "save":
		src.Rclone("save")
	}
}
