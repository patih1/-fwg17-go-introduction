package main

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generatePassword(input string, strength string) string {
	// charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charset := ""
	var password []byte

	input = strings.ToLower(input)
	var sliceStr []byte = []byte{}
	var resStr []byte

	for i := 0; i < len(input); i++ {
		sliceStr = nil
		sliceStr = append(sliceStr, input[i])
		sliceStr = append(sliceStr, input[i]-32)
		randStr := rand.Intn(len(sliceStr))
		resStr = append(resStr, sliceStr[randStr])
	}

	password = append(password, resStr...)

	if strength == "med" || strength == "strong" {
		charset += "0123456789"
	}
	if strength == "strong" {
		charset += "!@#$%^&*()_+=-"
	}

	for i := 0; i < len(input); i++ {
		randNum := rand.Intn(len(charset))
		password = append(password, charset[randNum])
	}
	return string(password)
}
