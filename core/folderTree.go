package core

import (
	"bytes"
	"html/template"
	"markdownFolderMap/static"
	"os"
	"path/filepath"
)

var folderTreeMap = map[string]FolderTree{}
var Template = map[string]string{"html": static.Temple}

type FolderTree struct {
	path   string
	filter string
	tree   map[string]interface{}
}

func NewFolderTree(path, filter string) (FolderTree, error) {
	if value, OK := folderTreeMap[path+filter]; OK {
		return value, nil
	}
	tree, err := walkMd(path, filter)
	if err != nil {
		return FolderTree{}, err
	}
	out := FolderTree{path: path, tree: tree}
	folderTreeMap[path+filter] = out
	return out, nil
}

func (f FolderTree) BuildHtml(temple string) (string, error) {
	var outBytes bytes.Buffer
	t1 := template.New("filemap")
	t1 = template.Must(t1.Parse(Template[temple]))
	err := t1.Execute(&outBytes, map[string]map[string]interface{}{"payload": f.tree})
	return outBytes.String(), err
}

func walkMd(path, filter string) (map[string]interface{}, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return walkMdOperator(path, filter, info)
}

func walkMdOperator(path, filter string, info os.FileInfo) (map[string]interface{}, error) {
	name := info.Name()
	if name[0] == '_' || name[0] == '.' || name[0] == ' ' {
		return nil, nil
	}
	if !info.IsDir() {
		return parseMd(path, info)
	} else {
		tem, _ := os.ReadDir(path)
		outList := make([]map[string]interface{}, 0, len(tem))
		for _, name := range tem {
			localPath := filepath.Join(path, name.Name())
			localInfo, _ := name.Info()
			localOut, _ := walkMdOperator(localPath, filter, localInfo)
			if localOut != nil {
				outList = append(outList, localOut)
			}
		}
		return map[string]interface{}{"content": info.Name(), "children": outList}, nil
	}
}

func parseMd(path string, info os.FileInfo) (map[string]interface{}, error) {
	return map[string]interface{}{"content": info.Name()}, nil
}
