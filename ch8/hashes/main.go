package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// Hash function takes a set of data and reduces it to a smaller fixed size.
// It's frequently used in programming for everything from looking up data to easily detecting changes.

// In Go hash function are broken into two categories: cryptographic and non-cryptographic.
// Non-cryptographic hash functions can be found underneath the hash package and include
// adler32, crc32, crc64, and fnv.

// A common use for crc32 is to compare two files.
// If the Sum32 value for both files is the same,
// it's highly likely (though not 100% certain) that the files are the same.
// If the values are different, then the files are definitely not hte same.

// getHash generate the hash based on file content
func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer f.Close()

	// create a hash
	h := crc32.NewIEEE()
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func main() {

	// Non-cryptographic

	// create a hash
	h := crc32.NewIEEE()
	// write out data to it
	text := "Hello, World!"
	_, _ = h.Write([]byte(text))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println("Checksum from hash crc32 for text", text, "is", v)

	file1 := "ch8/hashes/test1.txt"
	h1, err := getHash(file1)
	if err != nil {
		return
	}
	file2 := "ch8/hashes/test2.txt"
	h2, err := getHash(file2)
	if err != nil {
		return
	}
	fmt.Println("File 1:", file1, "hash1:", h1, "File 2:", file2, "hash2:", h2, "hash1 == hash2?", h1 == h2)

	fmt.Println()

	// cryptographic has the added property of being hard to reverse.
	// Given the cryptographic hash of a set of data, it's extremely difficult
	// to determine what made the hash.
	// These hashes are often used in security applications.
	// common hash SHA-1
	hashSha1 := sha1.New()
	hashSha1.Write([]byte(text))
	bs := hashSha1.Sum([]byte{})
	fmt.Println("Hash SHA-1 for the text:", text, "is", bs)
}
