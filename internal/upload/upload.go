package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

func SendFile(filePath, url string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	totalSize := fileInfo.Size()

	bar := progressbar.NewOptions64(
		totalSize,
		progressbar.OptionSetDescription("Загрузка"),
		progressbar.OptionSetWidth(40),
		progressbar.OptionShowBytes(true),
		progressbar.OptionClearOnFinish(),
	)

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	done := make(chan error, 1)

	go func() {
		defer pw.Close()
		defer writer.Close()

		part, err := writer.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			done <- pw.CloseWithError(err)
			return
		}

		progressReader := io.TeeReader(file, bar)

		if _, err := io.Copy(part, progressReader); err != nil {
			done <- pw.CloseWithError(err)
			return
		}

		done <- nil
	}()

	req, err := http.NewRequest("POST", url, pr)
	if err != nil {
		pw.CloseWithError(err)
		<-done
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Content-Length", strconv.FormatInt(totalSize, 10))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		pw.CloseWithError(err)
		return err
	}
	defer resp.Body.Close()

	if err := <-done; err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("\nОтвет сервера: %s\n", string(body))
	return nil
}
