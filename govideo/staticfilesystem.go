package govideo

import (
	"log"
	"net/http"
	"os"
)

// StaticFileSystem that disables directory listing
// https://groups.google.com/forum/#!topic/golang-nuts/bStLPdIVM6w
type StaticFileSystem struct {
	fs http.FileSystem
}

// Open only serves non-directory files
func (fs StaticFileSystem) Open(name string) (http.File, error) {
	log.Println("opening file " + name)
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}
