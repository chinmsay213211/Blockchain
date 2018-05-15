package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
)
type SimpleChaincode struct {
}

// this is simple program to read input file and sort to out file
//file name : input.txt
//file name 2 : output.txt

func main() {
	// open input file
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

//	for i := 0; i < 256; i++ {
//		dictionary[i] = string(i)
//	}

	var lines []string
	scanner := bufio.NewScanner("input.txt")
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()


	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}


