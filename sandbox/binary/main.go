package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
)
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// This decompresses a zlib compressed file and parses the content
func parseBinaryFile(filePath string) string{
	// check if the file exists
	if !fileExists(filePath) {
		fmt.Fprintf(os.Stderr, "File %s not found\n", filePath)
		os.Exit(1)
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		os.Exit(1)
	}
	bytes := bytes.NewReader(data)
	reader, err := zlib.NewReader(bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating zlib reader: %s\n", err)
		os.Exit(1)
	}
	defer reader.Close()
	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading decompressed data: %s\n", err)
		os.Exit(1)
	}
	return string(decompressedData)
}

// basic problem:
// parse the content of the file
// in this case: a git tree object  
// which has a peculiar but well-defined structure
// https://stackoverflow.com/questions/14790681/what-is-the-internal-format-of-a-git-tree-object
func parseContent(content string){
	fmt.Printf("%s", content)
	var parts = strings.Split(content, "\x00")
	// iterate over the parts
	for i, part := range parts {
		fmt.Printf("Part %d: %s\n", i, part)
	}
}
func main() {
	var path = "/Users/allantabilog/dev/temp/git-test-4/.git/objects/d7/e609f3481db74d6bbf1378afd1e8f898aa193c"
	var content = parseBinaryFile(path)
	parseContent(content)
}
