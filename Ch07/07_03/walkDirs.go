package main

import (
	"os"
	"path/filepath"
)

func main() {
	println("********************************************************************")
	root, _ := filepath.Abs(".")
	println("Path:",root)

	err := filepath.Walk(root, processPath)
	checkError(err)

}

func processPath(path string, info os.FileInfo, err error) error {
	checkError(err)
	//Check if not in root folder
	if path != "." {
		if(info.IsDir()){
			println("Directory:", path)
		} else {
			println("File:", path)
		}
	}
	return nil
}

func checkError(e error) error {
	//If there's an error
	// or If error is not null: Meaning there's an error
	if(e != nil){
		println("Error:", e)
		return e
	}
	return nil
}