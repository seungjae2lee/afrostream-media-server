package main

import (
        "os"
	"mp4"
//        "log"
)

func main() {
  mp4.Debug(true)
  mp4.ParseFile(os.Args[1], "eng")
}
