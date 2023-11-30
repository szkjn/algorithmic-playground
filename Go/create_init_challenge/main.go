package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run init_challenge.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	goFolder := fmt.Sprintf("Go/%s", filename)
	pythonFolder := "Python/scripts"
	typescriptFolder := "TypeScript/src"

	// Create Go folder and files
	os.MkdirAll(goFolder, os.ModePerm)
	createFile(fmt.Sprintf("%s/main.go", goFolder))
	createFile(fmt.Sprintf("%s/go.mod", goFolder))
	writeToFile(fmt.Sprintf("%s/go.mod", goFolder), fmt.Sprintf("module github.com/szkjn/algorithmic-playground/Go/%s\n\ngo 1.21.4\n", filename))

	// Create Python file
	createFile(fmt.Sprintf("%s/%s.py", pythonFolder, filename))

	// Create TypeScript file
	typescriptFilename := toCamelCase(filename)
	createFile(fmt.Sprintf("%s/%s.ts", typescriptFolder, typescriptFilename))

	fmt.Printf("Initialization of challenge %s complete.\n", filename)
}

func createFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Failed to create file %s: %s\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()
}

func writeToFile(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Failed to write to file %s: %s\n", filePath, err)
		os.Exit(1)
	}
}

func toCamelCase(str string) string {
	words := strings.Split(str, "_")
	for i := range words {
		if i > 0 {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}
