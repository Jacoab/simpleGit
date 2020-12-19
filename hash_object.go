package main

import (
	"crypto/sha1"
	"hash"
	"io"
)

func CmdHashObject(command string) (hash.Hash, error) {
	objHash := sha1.New()
	_, err := io.WriteString(objHash, command)
	return objHash, err
}