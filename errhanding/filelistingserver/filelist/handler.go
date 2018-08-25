package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {

	urlPath := request.URL.Path

	if index := strings.Index(urlPath, prefix); index != 0 {
		return userError("path must start with " + prefix)
	}
	path := urlPath[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(bytes)
	return nil
}
