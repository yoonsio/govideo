package govideo

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (a *App) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func (a *App) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	t.Execute(w, nil)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO
}

// list returns json list of all available media
// in paths specified in configuration file
func (a *App) list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// walk through a.config.App.Paths

	// identify media using extension
	// TODO: identify media by inspecting headers

	// TODO: export to json

}

func downloadFile(filepath string, url string) error {
	// create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// get data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil

}
