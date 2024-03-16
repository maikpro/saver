package services

import (
	"fmt"
	"io"
	"net/http"
)

func GetFileData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}
