package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("File Renamer Problem")

	dir := "./sample"

	fileNames := make([]string, 0)
	count := 0
	type rename struct {
		filename string
		path     string
	}

	var toRename []struct {
		filename string
		path     string
	}

	var err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		must(err)
		if !info.IsDir() {
			// fmt.Println(filepath.Base(path))
			fileName := filepath.Base(path)
			fileNames = append(fileNames, filepath.Base(path))
			_, err := match(fileName, 4)
			if err == nil {
				count++
				toRename = append(toRename, rename{
					path:     filepath.Join(dir, fileName),
					filename: fileName,
				})
			}
		}
		return nil
	})
	must(err)

	for _, orignalPath := range toRename {
		originalPath := filepath.Join(dir, orignalPath.filename)
		newName, err := match(orignalPath.filename, count)
		must(err)
		newPath := filepath.Join(dir, newName)
		fmt.Printf("change fileName from %s to %s\n", originalPath, newPath)
		err = os.Rename(originalPath, newPath)
		must(err)

	}
}

func match(filneName string, total int) (string, error) {
	//split the string
	pieces := strings.Split(filneName, ".")
	fileExt := pieces[len(pieces)-1]
	file := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(file, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s did not match pattern", filneName)
	}
	return fmt.Sprintf("%s - %d of %d.%s", name, number, total, fileExt), nil
}

func must(err error) error {
	if err != nil {
		return err
	}
	return nil
}
