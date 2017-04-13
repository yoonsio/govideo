package govideo

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/burntsushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/mailru/easyjson"
	"github.com/rakyll/magicmime"
	"github.com/sickyoon/govideo/govideo/models"
)

// App is GoVideo Web Application
// TODO: add securecookie
type App struct {
	*httprouter.Router
	handlers http.Handler
	config   models.Config
	db       *MongoClient
	auth     *AuthClient
	cache    *RedisClient
	store    *sessions.CookieStore
}

// NewApp creates new web application
func NewApp(configFile string) *App {

	log.Printf("Initializing web application with " + configFile)
	var err error

	// initialize app
	app := App{
		Router: httprouter.New(),
	}

	// load config file if exists
	if configFile != "" {
		if _, err := toml.DecodeFile(configFile, &app.config); err != nil {
			log.Fatal(err)
		}
	}

	// establish db connection
	app.db, err = NewMongoClient(app.config.Database.URI, app.config.Database.DBName)
	if err != nil {
		log.Fatal(err)
	}

	// establish redis connection
	app.cache, err = NewRedisClient(&app.config)
	if err != nil {
		log.Fatal(err)
	}

	// create session store
	app.store = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))
	app.store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // only over SSL - set 'true' for production
		MaxAge:   app.config.App.UserExpiry,
	}

	// create auth client
	app.auth, err = NewAuthClient(app.store, app.db, app.cache)
	if err != nil {
		log.Fatal(err)
	}

	// add handlers
	app.GET("/", app.index)
	app.GET("/login", app.index)
	app.GET("/profile", app.index)
	app.GET("/media", app.index)
	app.GET("/view", app.index)
	app.GET("/media/:encodedPath", app.index)
	app.POST("/login", app.loginPost)
	app.GET("/logout", app.logout)
	app.GET("/media/:encodedPath/info", app.auth.HttprouterMiddleware(app.infoFile))
	app.GET("/media/:encodedPath/data", app.auth.HttprouterMiddleware(app.serveFile(DATA)))
	app.GET("/media/:encodedPath/subtitle", app.auth.HttprouterMiddleware(app.serveFile(SUBTITLE)))
	app.Handler("GET", "/curuser", app.auth.Middleware(http.HandlerFunc(app.curUser)))
	app.Handler("GET", "/sync", app.auth.Middleware(http.HandlerFunc(app.sync)))
	app.Handler("GET", "/listMedia", app.auth.Middleware(http.HandlerFunc(app.list)))

	// add static resources handler
	staticPath := "static"
	if app.config.Server.StaticPath != "" {
		staticPath = app.config.Server.StaticPath
	}
	app.ServeFiles("/static/*filepath", fileOnlyFs{http.Dir(staticPath)})

	// add middlewares
	h := handlers.LoggingHandler(os.Stdout, app)
	h = handlers.ProxyHeaders(h)
	h = handlers.CompressHandler(h)
	h = handlers.RecoveryHandler()(h)
	app.handlers = h

	return &app
}

// Run starts server
func (a *App) Run() {
	log.Printf("Running server at port " + strconv.Itoa(a.config.Server.Port))
	log.Fatal(
		http.ListenAndServe(":"+strconv.Itoa(a.config.Server.Port), a.handlers),
	)
}

// Seed seeds database
func (a *App) Seed() error {
	log.Println("Creating test user")
	err := a.db.CreateUser(&models.User{
		Email:     "a",
		FirstName: "John",
		LastName:  "Doe",
		Hash:      []byte("a"),
	})
	if err != nil {
		return err
	}
	log.Println("test user successfully created")
	return nil
}

// Sync syncs database with media files
func (a *App) Sync() error {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		return err
	}
	for _, path := range a.config.App.Paths {
		err := filepath.Walk(path, a.registerFile)
		if err != nil {
			log.Printf("failed to process %s - %v", path, err)
			continue
		}
	}
	magicmime.Close()
	return nil
}

var videoFormats = []string{"video/mp4", "video/x-matroska", "text/plain"}
var subtitleExt = []string{"srt", "smi"}

func (a *App) registerFile(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {

		// check mimetype
		mimeType, err := magicmime.TypeByFile(path)
		if err != nil {
			// failed to check mimetype!
			log.Printf("failed to check mimetype for %s", path)
			return nil
		}

		// check if supported
		if !InSlice(videoFormats, mimeType) {
			// not supported!
			//log.Printf("media %s with %s is not supported", path, mimeType)
			return nil
		}

		// extract extension
		extension := filepath.Ext(path)[1:]
		name := info.Name()[0 : len(info.Name())-len(extension)-1]

		// update existing media if found
		if mimeType == "text/plain" && InSlice(subtitleExt, extension) {
			// search for media with same name
			err := a.db.UpdateSubtitle(name, path)
			if err != nil {
				log.Printf("failed to update subtitle for %s - %v", name, err)
			}
			return nil
		}

		// create new media
		media := models.GetMedia()
		media.Path = path
		media.Mimetype = mimeType
		media.Name = name
		media.Extension = extension
		media.Size = info.Size()
		media.Path = path
		// default acl is a?
		media.Access = []string{"a"}
		media.Added = time.Now().UTC()
		a.db.InsertMedia(media)
		models.RecycleMedia(media)
	}
	return nil
}

// ErrorHandler returns HTTP response with json error message
func ErrorHandler(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	errResponse := models.GetErrResponse()
	errResponse.Msg = error
	errResponse.Code = code
	easyjson.MarshalToHTTPResponseWriter(errResponse, w)
	models.RecycleErrResponse(errResponse)
}

type fileOnlyFs struct {
	fs http.FileSystem
}

func (fs fileOnlyFs) Open(name string) (http.File, error) {
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
