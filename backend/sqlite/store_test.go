package sqlite

import (
	"testing"

	"github.com/redbo/giphaux/shared"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestRandomID(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	id := db.RandomID()
	require.Equal(t, 32, len(id))
}

func TestNewUser(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	user, err := db.NewUser("redbo", "test")
	require.Nil(t, err)
	require.Equal(t, "redbo", user.Username)
}

func TestGetUser(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("redbo", "test")
	require.Nil(t, err)
	user, err := db.GetUser("redbo")
	require.Nil(t, err)
	require.Equal(t, "redbo", user.Username)
}

func TestGetUserByKey(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	user, err := db.NewUser("redbo", "test")
	require.Nil(t, err)
	user, err = db.GetUserByKey(user.APIKey)
	require.Nil(t, err)
	require.Equal(t, "redbo", user.Username)
}

func TestGetUserByCookie(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	user, err := db.NewUser("redbo", "test")
	require.Nil(t, err)
	user, err = db.GetUserByCookie(user.Cookie)
	require.Nil(t, err)
	require.Equal(t, "redbo", user.Username)
}

func TestFavorites(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	_, err = db.NewUser("user2", "test")
	require.Nil(t, err)
	gif, err := db.AddGIF("user1", "caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, db.AddCategory("user2", "funny"))
	require.Nil(t, db.AddFavorite("user2", gif.ID, []string{"funny"}))
	require.Nil(t, db.AddCategory("user2", "cats"))
	require.Nil(t, db.UpdateCategories("user2", gif.ID, []string{"funny", "cats"}))
	info, err := db.UserGIFInfo("user2", gif.ID)
	require.Nil(t, err)
	require.Equal(t, map[string]bool{"cats": true, "funny": true}, info.Categories)
	require.True(t, info.Favorited)
	require.False(t, info.Owner)
	db.RemoveFavorite("user2", gif.ID)
	info, err = db.UserGIFInfo("user2", gif.ID)
	require.Nil(t, err)
	require.Equal(t, map[string]bool{}, info.Categories)
	require.False(t, info.Favorited)
	require.False(t, info.Owner)
	require.Nil(t, db.RemoveGIF("user1", gif.ID))
	info, err = db.UserGIFInfo("user1", gif.ID)
	require.NotNil(t, err)
}

func TestSearch(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	_, err = db.AddGIF("user1", "caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, err)
	_, err = db.AddGIF("user1", "tag1", []string{}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, err)
	gifs, totalcount, err := db.Search("tag1", 10, 0, "g")
	require.Nil(t, err)
	require.Equal(t, 2, totalcount)
	require.Equal(t, 2, len(gifs))
	gifs, totalcount, err = db.Search("tag1", 1, 0, "g")
	require.Nil(t, err)
	require.Equal(t, 2, totalcount)
	require.Equal(t, 1, len(gifs))
	gifs, totalcount, err = db.Search("tag1", 10, 2, "g")
	require.Nil(t, err)
	require.Equal(t, 2, totalcount)
	require.Equal(t, 0, len(gifs))
	gifs, totalcount, err = db.Search("tag:tag1", 10, 0, "g")
	require.Nil(t, err)
	require.Equal(t, 1, totalcount)
	require.Equal(t, 1, len(gifs))
}

func TestUserFrontpage(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	require.Nil(t, db.AddCategory("user1", "cats"))
	require.Nil(t, db.AddCategory("user1", "funny"))
	gif, err := db.AddGIF("user1", "caption", []string{"tag1", "tag2"}, []string{"cats"}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, err)
	_, err = db.AddGIF("user1", "tag1", []string{}, []string{"funny"}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, err)
	fp, err := db.UserFrontpage("user1")
	require.Nil(t, err)
	require.Equal(t, 1, len(fp.Categories["funny"]))
	require.Equal(t, 1, len(fp.Categories["cats"]))
	require.Equal(t, 2, len(fp.Categories["Favorites"]))
	require.Equal(t, 2, len(fp.Categories["Uploads"]))
	require.Nil(t, db.RemoveFavorite("user1", gif.ID))
	fp, err = db.UserFrontpage("user1")
	require.Nil(t, err)
	require.Equal(t, 1, len(fp.Categories["funny"]))
	require.Equal(t, 0, len(fp.Categories["cats"]))
	require.Equal(t, 1, len(fp.Categories["Favorites"]))
	require.Equal(t, 2, len(fp.Categories["Uploads"]))
}
