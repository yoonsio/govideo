package govideo

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/julienschmidt/httprouter"
	"github.com/mailru/easyjson"
	"github.com/sickyoon/govideo/govideo/models"
)

func (a *App) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func (a *App) loginPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: validate empty values
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := a.auth.Authenticate(w, r, username, password)
	if err != nil {
		// Clean up
		a.auth.ClearUser(w, r)
		ErrorHandler(w, err.Error(), http.StatusUnauthorized)
		return
	}
	marshaller := jsonpb.Marshaler{
		EmitDefaults: false,
		OrigName:     true,
	}
	err = marshaller.Marshal(w, user)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := a.auth.ClearUser(w, r)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	response := models.GetSuccessResponse()
	response.Msg = "OK"
	easyjson.MarshalToHTTPResponseWriter(response, w)
	models.RecycleSuccessResponse(response)
}

func (a *App) curUser(w http.ResponseWriter, r *http.Request) {
	user, err := a.auth.CurUser(r)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusNoContent)
		return
	}
	marshaller := jsonpb.Marshaler{
		EmitDefaults: false,
		OrigName:     true,
	}
	err = marshaller.Marshal(w, user)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func serveFile(w http.ResponseWriter, r *http.Request, filepath string) error {
	fi, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fi.Close()
	http.ServeContent(w, r, "", time.Time{}, fi)
	return nil
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
