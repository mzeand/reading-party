package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chank '%v' (%d bytes)\n", string(buffer), length)
}

func rearChanks(file *os.File) []io.Reader {
	var chanks []io.Reader

	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chanks = append(chanks, io.NewSectionReader(file, offset, int64(length)+12))
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chanks
}

func main() {
	file, err := os.Open("bumper480x270.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := rearChanks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
