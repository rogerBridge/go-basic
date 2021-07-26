package main

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathRand "math/rand"
)

func init() {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
	for i:=0; i<4; i++ {
		fmt.Println(RandomIntGen())
	}
}

func RandomIntGen() int {
	s := make([]int, 0, 11)
	for i:=1; i<12; i++ {
		s = append(s, i)
	}
	n := mathRand.Int() % len(s)
	return s[n]
}