package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const Prefix = "/list/"

type UserError string

func (e UserError) Error() string {
	return e.Message()
	//panic("implement me")
}

func (e UserError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println()
	if strings.Index(request.URL.Path, Prefix) != 0 {
		return UserError(fmt.Sprintf("%s %s", request.URL.Path, Prefix))
	}

	path := request.URL.Path[len(Prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil

}
