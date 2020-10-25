package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main () {
	if len(os.Args) == 1 {
		handleErr(errors.New("No arguments provided for compressing or decompressing"))
		printHelp()
		return
	}
	if os.Args[1] == "--decompress" || os.Args[1] == "-d" {
		Decompress(os.Args[2])
		return
	}
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}
	Compress(os.Args[1])
}

func printHelp() {
	fmt.Println(`Tarzan - Compress files
Usage: tarzan --help/-h | --decompress/d <path> | <path>
Note: <path> is compressed if no options are used
Examples:
   tarzan file.txt
   tarzan -d file.txt.tarzan
   tarzan --help`)
}

func decompressBytes (toDecompress []byte) (decompressed []byte) {
	// TODO: add the actual algo
	return toDecompress
}

func Decompress (path string) {
	toDecompress, err := ioutil.ReadFile(path)
	handleErr(err)
	decompressed := decompressBytes(toDecompress)
	out, err := os.Create(strings.TrimSuffix(path, ".tarzan")) // TODO: maybe add a condition to check if file ends with .tarzan
	handleErr(err)
	defer out.Close()
	_, err = out.Write(decompressed)
	handleErr(err)
}

func compressBytes (toCompress []byte) (compressed []byte) {
	// TODO: add the actual algo
	return toCompress
}

func Compress (path string) {
	toCompress, err := ioutil.ReadFile(path)
	handleErr(err)
	compressed := compressBytes(toCompress)
	out, err := os.Create(path + ".tarzan")
	handleErr(err)
	defer out.Close()
	_, err = out.Write(compressed)
	handleErr(err)
}

func handleErr (err error) {
	if err != nil {
		fmt.Println("tarzan: error:", err)
	}
}


