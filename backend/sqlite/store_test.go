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

func TestFrontpage(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	_, err = db.AddGIF("user1", "caption", []string{"tag1", "tag2"}, []string{"cats"}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, err)
	_, err = db.AddGIF("user1", "tag1", []string{}, []string{"funny"}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, err)
	fp, err := db.Frontpage()
	require.Nil(t, err)
	require.Equal(t, 2, len(fp.Categories["Uploads"]))
	require.Equal(t, 0, len(fp.Categories["Trending"]))
}

func TestGIFByID(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	_, err = db.GIFByID("INCORRECTID")
	require.NotNil(t, err)
	gif, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, err)
	gif2, err := db.GIFByID(gif.ID)
	require.Nil(t, err)
	require.Equal(t, "some caption", gif2.Caption)
}

func TestGIFData(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	gif, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte("THIS IS SOME GIF DATA"))
	require.Nil(t, err)
	data, err := db.GIFData(gif.ID)
	require.Equal(t, []byte("THIS IS SOME GIF DATA"), data)
}

func TestGIFsByID(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	ids := []string{}
	for i := 0; i < 10; i++ {
		gif, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
		require.Nil(t, err)
		ids = append(ids, gif.ID)
	}
	gifs, totalcount, err := db.GIFsByID(ids, 5, 0)
	require.Nil(t, err)
	require.Equal(t, 10, totalcount)
	require.Equal(t, 5, len(gifs))
}

func TestUserUploads(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	ids := []string{}
	for i := 0; i < 10; i++ {
		gif, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
		require.Nil(t, err)
		ids = append(ids, gif.ID)
	}
	gifs, totalcount, err := db.UserUploads("user1", 5, 0)
	require.Nil(t, err)
	require.Equal(t, 10, totalcount)
	require.Equal(t, 5, len(gifs))
}

func TestUserFavorites(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	ids := []string{}
	for i := 0; i < 10; i++ {
		gif, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
		require.Nil(t, err)
		ids = append(ids, gif.ID)
	}
	db.RemoveFavorite("user1", ids[0])
	gifs, totalcount, err := db.UserFavorites("user1", 5, 0)
	require.Nil(t, err)
	require.Equal(t, 9, totalcount)
	require.Equal(t, 5, len(gifs))
}

func TestRandomSearch(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	for i := 0; i < 10; i++ {
		_, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
		require.Nil(t, err)
	}
	// pull a bunch of results, make sure we get all 10 uploaded images.
	seen := make(map[string]bool)
	for i := 0; i < 150; i++ {
		gif, err := db.RandomSearch("caption", 10)
		require.Nil(t, err)
		seen[gif.ID] = true
	}
	require.Equal(t, 10, len(seen))
}

func TestRandomByTag(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	for i := 0; i < 10; i++ {
		_, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
		require.Nil(t, err)
	}
	_, err = db.RandomByTag("caption", "r") // this shouldn't find anything
	require.NotNil(t, err)
	// pull a bunch of results, make sure we get all 10 uploaded images.
	seen := make(map[string]bool)
	for i := 0; i < 150; i++ {
		gif, err := db.RandomByTag("tag1", "g")
		require.Nil(t, err)
		seen[gif.ID] = true
	}
	require.Equal(t, 10, len(seen))
}

func TestRemoveCategory(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	gif, err := db.AddGIF("user1", "caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
	require.Nil(t, db.AddCategory("user1", "funny"))
	require.Nil(t, db.AddCategory("user1", "cats"))
	require.Nil(t, db.UpdateCategories("user1", gif.ID, []string{"funny", "cats"}))
	info, err := db.UserGIFInfo("user1", gif.ID)
	require.Equal(t, map[string]bool{"cats": true, "funny": true}, info.Categories)
	require.Nil(t, db.RemoveCategory("user1", "funny"))
	info, err = db.UserGIFInfo("user1", gif.ID)
	require.Equal(t, map[string]bool{"cats": true}, info.Categories)
}

func TestTrending(t *testing.T) {
	// I don't have anything that can actually mark an image as trending yet.
	// But this should at least make sure the queries compile.
	logger, _ := zap.NewDevelopment()
	db, _ := InitDatabase(&shared.Configuration{Database: ":memory:"}, logger)
	_, err := db.NewUser("user1", "test")
	require.Nil(t, err)
	ids := []string{}
	for i := 0; i < 10; i++ {
		gif, err := db.AddGIF("user1", "some caption", []string{"tag1", "tag2"}, []string{}, "", "g", 1, 1, 42, 1, []byte{})
		require.Nil(t, err)
		ids = append(ids, gif.ID)
	}
	gifs, totalcount, err := db.Trending(10, 0, "g")
	require.Nil(t, err)
	require.Equal(t, 0, totalcount)
	require.Equal(t, 0, len(gifs))
}
