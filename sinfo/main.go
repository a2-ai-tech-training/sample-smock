// https://slurm.schedmd.com/sinfo.html
package main

import (
	"bufio"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

//go:embed outputs/*
var f embed.FS

func main() {

	arguments := os.Args[1:]
	hasher := md5.New()

	for _, v := range arguments {
		hasher.Write([]byte(v))
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	hash_path := fmt.Sprintf("outputs/%s.txt", hash)

	file, err := f.Open(hash_path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skips first line that contains command
	scanner.Scan()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
