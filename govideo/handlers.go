package govideo

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/julienschmidt/httprouter"
)

func (a *App) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func (a *App) loginPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	user, err := a.auth.Authenticate(w, r, username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	marshaller := jsonpb.Marshaler{
		EmitDefaults: false,
		OrigName:     true,
	}
	err = marshaller.Marshal(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := a.auth.ClearUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("OK"))
}

func (a *App) curUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := a.auth.CurUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	marshaller := jsonpb.Marshaler{
		EmitDefaults: false,
		OrigName:     true,
	}
	err = marshaller.Marshal(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
