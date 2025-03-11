package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func decodeBase64(encodedStr string) (string, error) {
	// Decode the Base64 encoded string
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string and return it
	decodedStr := string(decodedBytes)
	return decodedStr, nil
}

func main() {
	// Array to store paths of .exp files
	var cfgFiles []string

	// Walk through the current directory
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file extension is .exp
		if !info.IsDir() && filepath.Ext(path) == ".exp" {
			cfgFiles = append(cfgFiles, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Array to store contents of .exp files
	var cfgContents []string

	// Read the contents of each .exp file
	for _, path := range cfgFiles {
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", path, err)
			continue
		}
		defer file.Close()

		var maxBufferSize = 32 * 1024 * 1024
		scanner := bufio.NewScanner(file)
		buffer := make([]byte, 0, maxBufferSize)
		scanner.Buffer(buffer, maxBufferSize)
		var content string
		for scanner.Scan() {
			content += scanner.Text() + "\n"
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			continue
		}

		cfgContents = append(cfgContents, content)
	}

	// Print the contents of each .exp file
	for i, content := range cfgContents {
		//fmt.Printf("Content of %s:\n%s\n", cfgFiles[i], content)
		var cleansw string
		cleansw = strings.ReplaceAll(content, "&&", "")
		decodedStr, swerr := decodeBase64(cleansw)
		output := strings.ReplaceAll(decodedStr, "&", "\n")
		file, fileerr := os.Create("sonicwall_config_export" + cfgFiles[i] + ".txt")
		if fileerr != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		if swerr != nil {
			fmt.Println("Error decoding Base64 string:", err)
			return
		}
		defer file.Close() // Ensure the file is closed when done
		_, writeErr := file.WriteString(output)
		if writeErr != nil {
			fmt.Println("Error writing to file:", writeErr)
			return
		}

		fmt.Println("File written successfully.")
	}
}
