package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

func gzipCompress(data string) ([]byte, error) {
	var buf bytes.Buffer 
	gzipWriter := gzip.NewWriter(&buf)
	_, err := gzipWriter.Write([]byte(data))
	if err != nil {
		return nil, err
	}
	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	data := "hello, world"
	compressedData, err := gzipCompress(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Compressed data: %v", compressedData)

	// decompress the data
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		panic(err)
	}
	decompressedData, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decompressed data: %v", string(decompressedData))

}