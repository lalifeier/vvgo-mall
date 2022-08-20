package util

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func WriteFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	wt := bufio.NewWriter(out)

	defer out.Close()

	_, err = io.Copy(wt, resp.Body)
	return err
}
