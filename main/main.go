package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// program cannot go to other folders when it finds a folder and renames files in it
// what should i do
// maybe seperate file path from folder path
// walk function have a problem
// idk how to fix it
// i can only rename files in one folder
// program does not walk properly
// ToDo if string found in directory name rename it then copy the renamed directory to path and go on renameing
// ToDo if wanted to rename files in a directorys copy directory name somewhere after renaming files in the directory
// ToDo copy directory name back to path
// ToDo or rename files with a seperate path than directorys
// ToDo make exception and try for when program changes a folder name except it ant try to run the program again
func main() {

	//if path is a directory {
	//	rename the directory name and copy the new path to path
	//}

	var fileName string
	OrgPath := "E:\\ Your Path \\"

	Recurse := true
	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if stat.IsDir() && path != OrgPath && !Recurse {
			fmt.Println("skipping dir:", path)
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}


		fileName = strings.Replace(path, " Find This String ", " Replace With This One", -1)


		if fileName != path {
			err = os.Rename(path, fileName)
			fmt.Println("\n", path, "\n\t\t\tRenamed to : \n", fileName)
		}
		return nil
	}
	err := filepath.Walk(OrgPath, walkFn)
	if err != nil {

		log.Fatal(err)
	}
}
