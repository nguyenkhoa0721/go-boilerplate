package generator

import (
	"bufio"
	"github.com/iancoleman/strcase"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GetModuleName() string {
	file, err := os.Open("go.mod")
	if err != nil {
		return ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	moduleNames := strings.Split(scanner.Text(), " ")
	if len(moduleNames) > 1 {
		return moduleNames[1]
	}

	return ""
}

func CopyDir(source, destination string, appName string, serviceName string, domain string) error {
	// Create destination directory if it doesn't exist
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		err = os.MkdirAll(destination, 0755)
		if err != nil {
			return err
		}
	}

	// Read source directory
	files, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	// Loop through files and directories
	for _, file := range files {
		sourcePath := filepath.Join(source, file.Name())
		destinationPath := filepath.Join(destination, file.Name())

		if file.IsDir() {
			// Recursively copy directories
			err = CopyDir(sourcePath, destinationPath, appName, serviceName, domain)
			if err != nil {
				return err
			}
		} else {
			// Copy and process template files
			err = copyAndProcessTemplate(sourcePath, destinationPath, appName, serviceName, domain)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyAndProcessTemplate(source, destination string, appName string, serviceName string, domain string) error {
	sourceContent, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	// Parse and execute template
	tmpl, err := template.New("file").Parse(string(sourceContent))
	if err != nil {
		return err
	}

	data := struct {
		App          string
		Service      string
		LowerService string
		Domain       string
	}{
		App:          appName,
		Service:      strcase.ToCamel(serviceName),
		LowerService: strcase.ToSnake(serviceName),
		Domain:       domain,
	}

	// Create the destination file
	destination = strings.TrimSuffix(destination, filepath.Ext(destination))
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Execute template and write to destination file
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}

	return nil
}
