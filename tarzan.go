package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main () {
	if len(os.Args) == 1 { // no args (go sets os.Args[1] to the executable's path)
		toCompress, _ := ioutil.ReadAll(os.Stdin) // compress stdin and send to stdout
		fmt.Printf("%s", compressBytes(toCompress))
		return
	}
	if os.Args[1] == "--decompress" || os.Args[1] == "-d" {
		if len(os.Args) == 2 { // only the option has been given so decompress stdin and send to stdout
			toDecompress, _ := ioutil.ReadAll(os.Stdin)
			fmt.Printf("%s", decompressBytes(toDecompress))
		} else {
			Decompress(os.Args[2])
		}
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
	if handleErr(err) {return}
	decompressed := decompressBytes(toDecompress)
	out, err := os.Create(strings.TrimSuffix(path, ".tarzan"))
	if handleErr(err) {return}
	defer out.Close()
	_, err = out.Write(decompressed)
	if handleErr(err) {return}
}

func compressBytes (toCompress []byte) (compressed []byte) {
	//var rowDelim byte = 500 // 111110100 in binary
    table := make([]byte, 1, len(toCompress)/3)
	table[0] = 0
	//repeatedByteSlices := make([][]byte, 1, (len(toCompress)/3) + 1) // TODO: Make algo to find repeated byte slices using the subset algo

	//fmt.Println(table)
	// TODO: use info at https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-wusp/2f6ddb6a-9026-43a3-b1d9-d8a19af3f03f to implement a "bitmask" so there are no edge cases
	return toCompress
}

func byteSliceContainsSubset(targetSlice, candidateSlice []byte) bool {
	if targetSlice == nil || candidateSlice == nil {return false}
	if len(targetSlice) < len(candidateSlice) {return false}
	if len(targetSlice) == len(candidateSlice) {
		return byteSlicesAreEqual(targetSlice, candidateSlice)
	}
	for i := 0; i < (len(targetSlice) - len(candidateSlice) + 1); i++ {
		if byteSlicesAreEqual(targetSlice[i:(i+len(candidateSlice))], candidateSlice) {
			return true
		}
	}
	return false
}

func byteSlicesAreEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Compress (path string) {
	toCompress, err := ioutil.ReadFile(path)
	if handleErr(err) {return}
	compressed := compressBytes(toCompress)
	out, err := os.Create(path + ".tarzan")
	if handleErr(err) {return}
	defer out.Close()
	_, err = out.Write(compressed)
	if handleErr(err) {return}
}

func handleErr (err error) bool {
	if err != nil {
		fmt.Println("tarzan: error:", err)
		return true
	}
	return false
}


