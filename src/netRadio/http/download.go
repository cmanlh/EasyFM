package http

import (
	"bufio"
	"errors"
	"net/http"
	"os"
)

const (
	netBufferSize  = 1048576
	fileBufferSize = 10485760
)

func Download(url string, path string) error {
	resp, err := http.Get(url)

	if nil != err {
		return err
	}

	if http.StatusNotFound == resp.StatusCode {
		return errors.New("Not find the mp3!")
	}

	mp3File, fileErr := os.Create(path)
	if fileErr != nil {
		return fileErr
	}

	netBuffer := make([]byte, netBufferSize)
	len, readErr := resp.Body.Read(netBuffer)

	buffredFile := bufio.NewWriterSize(mp3File, fileBufferSize)
	for readErr == nil {
		buffredFile.Write(netBuffer[0:len])
		len, readErr = resp.Body.Read(netBuffer)
	}
	buffredFile.Write(netBuffer[0:len])
	buffredFile.Flush()
	mp3File.Close()

	return nil
}
