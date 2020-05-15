package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	//var strSource, strDest string
	strSource := flag.String("source", "n/a", "specify source file")
	strDest := flag.String("dest", "n/a", "specify destination")
	flag.Parse()
	//fmt.Println(*strSource, *strDest)
	copy(*strSource, *strDest)

}

func copy(src, dst string) (int64, error) {
	fmt.Println(src, dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
