package main

import (
	// "fmt"
	// "time"
	"encoding/json"
	"os"
	"path/filepath"
)

func main() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "0-Repo/TEST-2/MyApp", "MyTempFile.json")
		
	}
	Printfln("Full path: %v", path)
	err = os.MkdirAll(filepath.Dir(path), 0766)
	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}
}
