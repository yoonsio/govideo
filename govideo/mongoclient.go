package govideo

import (
	"log"

	"github.com/sickyoon/govideo/govideo/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	colMedia = "media"
	colUser  = "users"
)

// MongoClient holds master session and other db-related info
type MongoClient struct {
	session *mgo.Session // master session
	uri     string       // mongodb uri
	dbName  string       // database name
}

// NewMongoClient establishes connection to MongoDB database
// and returns new MongoClient object
func NewMongoClient(uri, dbName string) *MongoClient {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return &MongoClient{session, uri, dbName}
}

// GetSession returns mgo.Session copied from
// MongoClient's master session
// Be sure to close the session after done
func (mc *MongoClient) GetSession() *mgo.Session {
	return mc.session.Copy()
}

// CRUD

// CreateUser -
func (mc *MongoClient) CreateUser(user *models.User) error {
	s := mc.GetSession()
	err := s.DB(mc.dbName).C(colUser).Insert(user)
	s.Close()
	return err
}

// GetUserFromDB -
func (mc *MongoClient) GetUserFromDB(email, hash string) (*models.User, error) {
	s := mc.GetSession()
	user := models.User{}
	err := s.DB(mc.dbName).C(colUser).Find(bson.M{"email": email, "password": hash}).One(&user)
	if err != nil {
		return nil, err
	}
	s.Close()
	return &user, nil
}

// InsertMedia -
func (mc *MongoClient) InsertMedia(media *models.Media) error {
	s := mc.GetSession()
	err := s.DB(mc.dbName).C(colMedia).Insert(media)
	s.Close()
	return err
}

// FindMedia -
func (mc *MongoClient) FindMedia(mediaID string) (*models.Media, error) {
	s := mc.GetSession()
	var media models.Media
	err := s.DB(mc.dbName).C(colMedia).Find(mediaID).One(&media)
	s.Close()
	return &media, err
}

// UpdateMedia -
func (mc *MongoClient) UpdateMedia(media *models.Media) error {
	s := mc.GetSession()
	err := s.DB(mc.dbName).C(colMedia).Update(
		bson.M{"_id": media.ID},
		bson.M{"name": media.Name},
	)
	s.Close()
	return err
}

// RemoveMedia -
func (mc *MongoClient) RemoveMedia(mediaID string) error {
	s := mc.GetSession()
	err := s.DB(mc.dbName).C("media").Remove(mediaID)
	s.Close()
	return err
}
