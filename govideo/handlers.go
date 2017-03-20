package govideo

import (
	"html/template"
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	err := a.auth.ClearUser(w, r)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

// sync syncs database with real files for details
func (a *App) sync(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	err := a.Sync()
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := models.GetSuccessResponse()
	response.Msg = "OK"
	easyjson.MarshalToHTTPResponseWriter(response, w)
	models.RecycleSuccessResponse(response)
}

// list returns json list of all available media
// in paths specified in configuration file
// all videos are added to dbs automatically
// this funciton just gets videos from dbs
// everytime video is requested, it returns fake path that lasts 24 hrs
func (a *App) list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// walk through a.config.App.Paths

	// identify media using extension
	// TODO: identify media by inspecting headers

	// get file database id

	// show updateAccess if user is admin

	// generate fakepath based on user ip & file db id

	// returns fakepath

	// TODO: export to json

}

// updateAccess updates access control for individual file
func (a *App) updateAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get file database id
	// query to display access control field
	// make sure the user has previlege (admin)
	// update access control field with post values
}

func (a *App) view(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// render page that shows the video
}

// serveFile serves actual video content in chunk based on filepath
func serveFile(w http.ResponseWriter, r *http.Request, filepath string) error {

	// query redis with fakepath to get real path
	// matches them against query ip
	fi, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fi.Close()
	http.ServeContent(w, r, "", time.Time{}, fi)
	return nil
}
