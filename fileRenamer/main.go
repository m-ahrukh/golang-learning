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

	// f, err := os.Open(dir)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// files, err := f.ReadDir(-1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	if !file.IsDir() {
	// 		fmt.Println(file.Name())
	// 	}
	// }

	///Returns the presend working directory
	// wd, err := os.Getwd()

	fileNames := make([]string, 0)
	var err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		must(err)
		if !info.IsDir() {
			fmt.Println(filepath.Base(path))
			fileNames = append(fileNames, filepath.Base(path))
		}
		return nil
	})

	must(err)

	newName, err := match(fileNames[0])

	if err != nil {
		fmt.Println("no match")
		os.Exit(1)
	}

	fmt.Println("New Name:", newName)
}

func match(filneName string) (string, error) {
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
	return fmt.Sprintf("%s - %d.%s", name, number, fileExt), nil
}

func must(err error) error {
	if err != nil {
		return err
	}
	return nil
}
