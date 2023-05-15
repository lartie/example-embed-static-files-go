//go:generate go run generator.go

package box

import (
	"os"
	"strings"
)

type FileInfo struct {
	Name    string      `json:"name"`
	Content []byte      `json:"content"`
	Perm    os.FileMode `json:"perm"`
}

type embedBox struct {
	storage map[string]FileInfo
}

// Create new box for embed files
func newEmbedBox() *embedBox {
	return &embedBox{storage: make(map[string]FileInfo)}
}

// Add a file to box
func (e *embedBox) Add(file string, fileInfo FileInfo) {
	e.storage[file] = fileInfo
}

// Get file's content
// Always use / for looking up
// For example: /init/README.md is actually configs/init/README.md
func (e *embedBox) Get(file string) []FileInfo {
	var files []FileInfo
	if strings.Contains(file, "*") {
		for s, info := range e.storage {
			if strings.HasPrefix(s, strings.TrimRight(file, "*")) {
				files = append(files, info)
			}
		}
	} else {
		if f, ok := e.storage[file]; ok {
			files = append(files, f)
		}
	}
	return files
}

// Find for a file
func (e *embedBox) Has(file string) bool {
	if _, ok := e.storage[file]; ok {
		return true
	}
	return false
}

// Embed box expose
var box = newEmbedBox()

// Add a file content to box
func Add(file string, fileInfo FileInfo) {
	box.Add(file, fileInfo)
}

// Get a file from box
func Get(file string) []FileInfo {
	return box.Get(file)
}

// Has a file in box
func Has(file string) bool {
	return box.Has(file)
}
