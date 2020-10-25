package main

import "os"

func main () {
	if os.Args[1] == "--decompress" || os.Args[1] == "-d" {
		decompress(os.Args[2])
		return
	}
	compress(os.Args[1])
}

func decompress (path string) {

}
func compress (path string) {

}
