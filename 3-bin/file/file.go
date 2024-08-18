package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	if ok, err := isJson(path); !ok {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, path string) {
	fileData, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer fileData.Close()

	bytes, err := fileData.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Written %d bytes", bytes)
}

func isJson(path string) (bool, error) {
	pathSlice := strings.Split(path, "/")
	fmt.Println(pathSlice)
	fmt.Println(path)

	if !strings.Contains(pathSlice[len(pathSlice)-1], ".json") {
		fmt.Println("Not a JSON file")
		return false, errors.New("Incorrect file format received")
	}

	return true, nil
}
