package debug

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PrintFSContent(fs http.FileSystem, maxDepth int) {
	log.Println("Printing FS content:")
	printDir(fs, "/", 0, "", maxDepth)
}

func printDir(fs http.FileSystem, path string, level int, prefix string, maxDepth int) {
	if maxDepth >= 0 && level >= maxDepth {
		return
	}

	dir, err := fs.Open(path)
	if err != nil {
		fmt.Printf("%sError opening directory %s: %v\n", prefix, path, err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("%sError reading directory %s: %v\n", prefix, path, err)
		return
	}

	for i, file := range files {
		isLast := i == len(files)-1
		newPrefix := prefix
		if isLast {
			fmt.Printf("%s└── %s\n", prefix, file.Name())
			newPrefix += "    "
		} else {
			fmt.Printf("%s├── %s\n", prefix, file.Name())
			newPrefix += "│   "
		}

		if file.IsDir() {
			printDir(fs, strings.TrimSuffix(path, "/")+"/"+file.Name(), level+1, newPrefix, maxDepth)
		}
	}
}
