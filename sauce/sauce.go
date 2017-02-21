// %N0 CARR13R█

// Package sauce implements a ACiD's Standard Architecture for Universal Comment Extensions
package sauce

import (
	"os"
)

const (
	dateFormat = "%Y%m%d"
)

type Record struct {
	ID       string `sauce:"required"`
	Version  string `sauce:"required"`
	Title    string
	Author   string
	Group    string
	Date     string
	FileSize string
	DataType string `sauce:"required"`
	FileType string `sauce:"required"`
	TInfo1   string
	TInfo2   string
	TInfo3   string
	TInfo4   string
	Comments string
	TFlags   string
	TInfoS   string
}

// Exists returns true if a SAUCE record is found.
func Exists(filename string) bool {
	file, _ := os.Open(filename)
	defer file.Close()

	buffer := make([]byte, 1024)

	_, _ = file.Seek(-128, 2)

	_, err := file.Read(buffer)

	if err != nil {
		panic(err)
	}

	if string(buffer[0:5]) == "SAUCE" {
		return true
	}
	return false
}

/*
func example() {



	file, err := os.Open("n0-signature.ans")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

s := sauce.Get(file)


	buffer := make([]byte, 1024)

	newoffset, _ := file.Seek(-128, 2)

	n, err := file.Read(buffer)

	if err != nil {
		panic(err)
	}

	// out the buffer content
	fmt.Println(string(buffer[:n]))
	fmt.Println(string(buffer[0:5]))

	fmt.Println("New off set : ", newoffset)
}


var s sauceRecord
*/
