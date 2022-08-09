package implement

import (
	"filemap/core"
	"fmt"
	"net/http"
	"path/filepath"
)

var path string

func Server(port, _path, filter string, depth int) {
	if _path == "" {
		_path, _ = filepath.Abs(".")
	}
	path = _path
	fmt.Println("server is http://127.0.0.1:" + port)
	http.HandleFunc("/", handleFunc(path, filter, depth))
	http.ListenAndServe("127.0.0.1:"+port, nil)
}

func handleFunc(path, filter string, depth int) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		folderTree, err := core.NewFolderTree(path, filter, depth)
		if err != nil {
			panic(err)
		}
		p, err := folderTree.BuildHtml("html")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(writer, p)
	}
}
