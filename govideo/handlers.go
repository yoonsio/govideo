package govideo

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"os"

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
func (a *App) list(w http.ResponseWriter, r *http.Request) {
	user, err := a.auth.CurUser(r)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusForbidden)
		return
	}
	mediaList, err := a.db.GetAllMedia(user.Email)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get client ip
	ipAddr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// encode media path
	for i := 0; i < len(mediaList.Data); i++ {
		media := &mediaList.Data[i]
		encodedPath, err := a.cache.GetEncodedPath(media.Path, ipAddr)
		if err != nil {
			ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		}
		log.Println(media.Path + " -> " + encodedPath)
		media.Path = encodedPath
	}
	easyjson.MarshalToHTTPResponseWriter(mediaList, w)
	models.RecycleMediaList(mediaList)
}

// updateAccess updates access control for individual file
func (a *App) updateAccess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get file database id
	// query to display access control field
	// make sure the user has previlege (admin)
	// update access control field with post values
}

// serveFile serves actual video content in chunk based on encoded filepath
func (a *App) serveFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	filepath := ps.ByName("encodedPath")

	// get client ip
	ipAddr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// query redis with fakepath to get real path
	realPath, err := a.cache.GetRealPath(filepath, ipAddr)
	if err != nil {
		ErrorHandler(w, "Invalid encoded path", http.StatusBadRequest)
		return
	}

	// matches them against query ip
	info, err := os.Stat(string(realPath))
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fi, err := os.Open(string(realPath))
	if err != nil {
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer fi.Close()
	http.ServeContent(w, r, "", info.ModTime(), fi)
}
