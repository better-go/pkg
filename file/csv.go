package file

import (
	"os"

	"github.com/gocarina/gocsv"
)

// 解析 csv:
func ParseCsv(filename string, dist interface{}) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return gocsv.UnmarshalFile(f, dist) // to dist
}
