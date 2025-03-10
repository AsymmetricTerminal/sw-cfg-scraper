package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Function to write a string to a file.
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

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
	// Array to store paths of .cfg files
	var cfgFiles []string

	// Walk through the current directory
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file extension is .cfg
		if !info.IsDir() && filepath.Ext(path) == ".cfg" {
			cfgFiles = append(cfgFiles, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Array to store contents of .cfg files
	var cfgContents []string

	// Read the contents of each .cfg file
	for _, path := range cfgFiles {
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", path, err)
			continue
		}
		defer file.Close()
		info, infoErr := file.Stat()
		if infoErr != nil {
			fmt.Printf("Error opening file %s: %v\n", path, err)
			continue
		}
		var maxSize int
		scanner := bufio.NewScanner(file)
		maxSize = int(info.Size())
		buffer := make([]byte, 0, maxSize)
		scanner.Buffer(buffer, maxSize)
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

	// Print the contents of each .cfg file
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
