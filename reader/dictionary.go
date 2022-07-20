package reader

import (
	"io/ioutil"
	"strings"
)

func Read(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	strData := string(data)
	strData = strings.ReplaceAll(strData, "\r", "")
	return strings.Split(strData, "\n"), nil
}
