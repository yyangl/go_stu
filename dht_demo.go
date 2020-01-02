package main

import "github.com/yyangl/go_stu/dht"

func main() {
	d := dht.New(":13303")
	d.Run()
}
