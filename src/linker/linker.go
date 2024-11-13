package linker

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Node string

func (node *Node) linkNode() string {
	lines := strings.Split(string(*node), "\n")

	var outputOfMyLink strings.Builder

	for _, line := range lines {
		if strings.HasPrefix(line, "#exec") {
			linkableFileName := strings.Split(line, " ")[1]
			fmt.Println("Found line with link, linking to " + linkableFileName)

			content, err := os.ReadFile(linkableFileName + ".bf")
			if err != nil {
				fmt.Println("Could not link, exiting")
				os.Exit(0)
			}

			node := Node(content[:])
			line = node.linkNode()
		}

		outputOfMyLink.WriteString(line)
	}

	return outputOfMyLink.String()
}

func Link(files []*os.File) (*os.File, error) {
	// no linking needed
	if len(files) == 1 {
		return files[0], nil
	}

	fileContents, err := filesToNodes(files)
	if err != nil {
		fmt.Println("Could not link")
		os.Exit(1)
	}

	mainFile := fileContents[0]

	linkedFileContent := mainFile.linkNode()

	return WriteStringToFile(linkedFileContent)
}

func WriteStringToFile(content string) (*os.File, error) {

	filename := "objectFile.o"

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	if _, err := file.WriteString(content); err != nil {
		file.Close()
		return nil, err
	}

	if _, err := file.Seek(0, 0); err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}

// defer this
func CleanObjectFiles() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Could not clean files")
		}
		// rm all object files
		if !info.IsDir() && filepath.Ext(path) == ".o" {
			fmt.Println("Removing:", path)
			if err := os.Remove(path); err != nil {
				return err
			}
		}
		return nil
	})
}

func filesToNodes(files []*os.File) ([]Node, error) {
	var contents []Node
	for _, file := range files {
		data, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", file.Name(), err)
		}
		contents = append(contents, Node(data))
	}
	return contents, nil
}
