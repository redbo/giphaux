package giphaux

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"go.uber.org/zap"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/redbo/giphaux/shared"
	"github.com/stretchr/testify/require"
)

//go:generate mockgen -destination mockgen_test.go -package giphaux github.com/redbo/giphaux/shared DataStore

type templateCapture struct {
	user  *shared.User
	query string
	data  interface{}
}

func (t *templateCapture) ExecuteTemplate(w io.Writer, templateName string, data interface{}) error {
	if dm, ok := data.(map[string]interface{}); ok {
		if v := dm["User"]; v != nil {
			t.user = v.(*shared.User)
		}
		if v := dm["Query"]; v != nil {
			t.query = v.(string)
		}
		t.data = dm["Data"]
	}
	return nil
}

type testAPIResponse struct {
	Meta struct{ Status int }
	Data map[string]interface{}
}

func getAPIResponse(f http.HandlerFunc, r *http.Request) *testAPIResponse {
	w := httptest.NewRecorder()
	rsp := testAPIResponse{}
	f(w, r)
	if err := json.Unmarshal(w.Body.Bytes(), &rsp); err != nil {
		panic(err.Error())
	}
	return &rsp
}

func TestApiRandomID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds}
	r := httptest.NewRequest("GET", "/v1/random", bytes.NewBuffer([]byte{}))
	ds.EXPECT().RandomID().Return("THIS IS A RANDOM ID")
	rsp := getAPIResponse(srv.apiRandomID, r)
	require.Equal(t, 200, rsp.Meta.Status)
	require.Equal(t, "THIS IS A RANDOM ID", rsp.Data["random_id"])
}

func TestApiSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("GET", "/v1/gifs/search?q=cats", bytes.NewBuffer([]byte{}))
	ds.EXPECT().Search(gomock.Eq("cats"), gomock.Eq(srv.queryLimit), gomock.Eq(0), gomock.Eq("g"))
	srv.apiSearch(httptest.NewRecorder(), r)

	r = httptest.NewRequest("GET", "/v1/gifs/search?q=cats&limit=5&offset=2&rating=pg", bytes.NewBuffer([]byte{}))
	ds.EXPECT().Search(gomock.Eq("cats"), gomock.Eq(5), gomock.Eq(2), gomock.Eq("pg"))
	rsp := getAPIResponse(srv.apiSearch, r)
	require.Equal(t, 200, rsp.Meta.Status)
}

func TestAPIGifID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("GET", "/v1/gifs/123", bytes.NewBuffer([]byte{}))
	r = mux.SetURLVars(r, map[string]string{"id": "1234567890abcdef"})
	ds.EXPECT().GIFByID(gomock.Eq("1234567890abcdef")).Return(&shared.GIF{}, nil)
	rsp := getAPIResponse(srv.apiGifID, r)
	require.Equal(t, 200, rsp.Meta.Status)
}

func TestAPIGifs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("GET", "/v1/gifs?ids=1234567890abcdef,ghijklmnoprstuvwxyz", bytes.NewBuffer([]byte{}))
	ds.EXPECT().GIFsByID(gomock.Eq([]string{"1234567890abcdef", "ghijklmnoprstuvwxyz"}), gomock.Eq(10), gomock.Eq(0))
	rsp := getAPIResponse(srv.apiGifs, r)
	require.Equal(t, 200, rsp.Meta.Status)
}

func TestAPITrending(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("GET", "/v1/gifs/trending", bytes.NewBuffer([]byte{}))
	ds.EXPECT().Trending(gomock.Eq(10), gomock.Eq(0), gomock.Eq("g"))
	rsp := getAPIResponse(srv.apiTrending, r)
	require.Equal(t, 200, rsp.Meta.Status)
}

func TestAPITranslate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("GET", "/v1/gifs/translate?s=cats&weirdness=7", bytes.NewBuffer([]byte{}))
	ds.EXPECT().RandomSearch(gomock.Eq("cats"), gomock.Eq(7))
	rsp := getAPIResponse(srv.apiTranslate, r)
	require.Equal(t, 200, rsp.Meta.Status)
}

func TestAPIRandomSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("GET", "/v1/gifs/random?tag=cat", bytes.NewBuffer([]byte{}))
	ds.EXPECT().RandomByTag(gomock.Eq("cat"), gomock.Eq("g"))
	rsp := getAPIResponse(srv.apiRandomSearch, r)
	require.Equal(t, 200, rsp.Meta.Status)
}

func TestAPIUpload(t *testing.T) {
	gifdata, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	gifsrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(gifdata)
	}))
	defer gifsrv.Close()
	body := []byte("tags=cat,sunglasses&source_image_url=" + url.QueryEscape(gifsrv.URL))
	r := httptest.NewRequest("POST", "/v1/gifs", bytes.NewBuffer(body))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	ds.EXPECT().AddGIF(
		gomock.Eq("redbo"), gomock.Eq(""), gomock.Eq([]string{"cat", "sunglasses"}),
		gomock.Eq([]string{}), gomock.Eq(gifsrv.URL), gomock.Eq("g"), gomock.Eq(1),
		gomock.Eq(1), gomock.Eq(42), gomock.Eq(1), gomock.Eq(gifdata))
	ctx := context.WithValue(r.Context(), userKey, &user)
	rsp := getAPIResponse(srv.apiUploadGif, r.WithContext(ctx))
	require.Equal(t, 200, rsp.Meta.Status)
}

func TestUserIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	tc := &templateCapture{}
	srv := server{
		ds:         ds,
		queryLimit: 10,
		templates:  tc,
	}
	r := httptest.NewRequest("GET", "/user/", bytes.NewBuffer([]byte{}))
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	ctx := context.WithValue(r.Context(), userKey, &user)
	w := httptest.NewRecorder()
	fp := &shared.FrontPageData{}
	ds.EXPECT().UserFrontpage(gomock.Eq("redbo")).Return(fp, nil)
	srv.userIndex(w, r.WithContext(ctx))
	require.Equal(t, fp, tc.data)
}

func TestUserCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("POST", "/user/addcategory", bytes.NewBuffer([]byte("category=cats%20with%20hats")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	w := httptest.NewRecorder()
	ds.EXPECT().AddCategory(gomock.Eq("redbo"), gomock.Eq("cats-with-hats")).Return(nil)
	srv.userAddCategory(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusSeeOther, w.Code)

	r = httptest.NewRequest("POST", "/user/removecategory", bytes.NewBuffer([]byte("category=cats%20with%20hats")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	ds.EXPECT().RemoveCategory(gomock.Eq("redbo"), gomock.Eq("cats-with-hats")).Return(nil)
	srv.userRemoveCategory(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusSeeOther, w.Code)

	r = httptest.NewRequest("POST", "/user/updatecategories", bytes.NewBuffer([]byte("gifid=0123456789abcdef&cat-cats=true&cat-dogs=true")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	ds.EXPECT().UpdateCategories(gomock.Eq("redbo"), gomock.Eq("0123456789abcdef"), gomock.Len(2))
	srv.userUpdateCategories(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusSeeOther, w.Code)
}

func TestUserFavorites(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("POST", "/user/favorite", bytes.NewBuffer([]byte("gifid=0123456789abcdef&cat-cats=true&cat-dogs=true")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	w := httptest.NewRecorder()
	ds.EXPECT().AddFavorite(gomock.Eq("redbo"), gomock.Eq("0123456789abcdef"), gomock.Len(2)).Return(nil)
	srv.userFavorite(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusSeeOther, w.Code)

	r = httptest.NewRequest("POST", "/user/unfavorite", bytes.NewBuffer([]byte("gifid=0123456789abcdef")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	ds.EXPECT().RemoveFavorite(gomock.Eq("redbo"), gomock.Eq("0123456789abcdef")).Return(nil)
	srv.userUnfavorite(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusSeeOther, w.Code)
}

func TestUserUploads(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10, templates: &templateCapture{}}
	r := httptest.NewRequest("GET", "/user/uploads", bytes.NewBuffer([]byte("")))
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	w := httptest.NewRecorder()
	ds.EXPECT().UserUploads(gomock.Eq("redbo"), gomock.Eq(12), gomock.Eq(0)).Return([]*shared.GIF{}, 0, nil)
	srv.userUploads(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusOK, w.Code)
}

func TestUserFavoritesList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10, templates: &templateCapture{}}
	r := httptest.NewRequest("GET", "/user/favorites", bytes.NewBuffer([]byte("")))
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	w := httptest.NewRecorder()
	ds.EXPECT().UserFavorites(gomock.Eq("redbo"), gomock.Eq(12), gomock.Eq(0)).Return([]*shared.GIF{}, 0, nil)
	srv.userFavorites(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusOK, w.Code)
}

func TestUserDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}
	r := httptest.NewRequest("POST", "/user/delete", bytes.NewBuffer([]byte("gifid=0123456789abcdef")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	w := httptest.NewRecorder()
	ds.EXPECT().RemoveGIF(gomock.Eq("redbo"), gomock.Eq("0123456789abcdef")).Return(nil)
	srv.userDelete(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusSeeOther, w.Code)
}

func TestUserUpload(t *testing.T) {
	gifdata, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{ds: ds, queryLimit: 10}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("caption", "deal with it")
	writer.WriteField("tags", "monkey,lips,funny")
	writer.WriteField("rating", "PG-13") // it is a sassy 1 pixel gif
	writer.WriteField("cat-monkey", "")
	writer.WriteField("cat-lips", "")
	part, _ := writer.CreateFormFile("uploadFile", "1.gif")
	part.Write(gifdata)
	writer.Close()
	r := httptest.NewRequest("POST", "/user/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	w := httptest.NewRecorder()
	ds.EXPECT().AddGIF(gomock.Eq("redbo"), gomock.Eq("deal with it"), gomock.Len(3), gomock.Len(2),
		gomock.Eq(""), gomock.Eq("pg-13"), gomock.Eq(1), gomock.Eq(1), gomock.Eq(1), gomock.Eq(42), gomock.Eq(gifdata)).Return(&shared.GIF{ID: "success"}, nil)
	srv.userUpload(w, r.WithContext(context.WithValue(r.Context(), userKey, &user)))
	require.Equal(t, http.StatusSeeOther, w.Code)
}

func TestAuthorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	tc := &templateCapture{}
	logger, _ := zap.NewDevelopment()
	srv := NewServer(&shared.Configuration{}, logger, ds, tc)

	r := httptest.NewRequest("GET", "/user/", bytes.NewBuffer([]byte{}))
	r.Header.Set("Cookie", "session=12345")
	user := shared.User{Username: "redbo", APIKey: "something", Password: "something", Cookie: "something", Categories: nil}
	w := httptest.NewRecorder()
	ds.EXPECT().GetUserByCookie(gomock.Eq("12345")).Return(&user, nil)
	fp := &shared.FrontPageData{}
	ds.EXPECT().UserFrontpage(gomock.Eq("redbo")).Return(fp, nil)
	srv.ServeHTTP(w, r)
	require.Equal(t, fp, tc.data)

	r = httptest.NewRequest("GET", "/user/", bytes.NewBuffer([]byte{}))
	r.Header.Set("Cookie", "session=12345")
	w = httptest.NewRecorder()
	ds.EXPECT().GetUserByCookie(gomock.Eq("12345")).Return(nil, fmt.Errorf("User not found"))
	srv.ServeHTTP(w, r)
	require.Equal(t, http.StatusUnauthorized, w.Code)

	r = httptest.NewRequest("GET", "/user/?api_key=12345", bytes.NewBuffer([]byte{}))
	w = httptest.NewRecorder()
	ds.EXPECT().GetUserByKey(gomock.Eq("12345")).Return(&user, nil)
	ds.EXPECT().UserFrontpage(gomock.Eq("redbo")).Return(fp, nil)
	srv.ServeHTTP(w, r)
	require.Equal(t, fp, tc.data)

	r = httptest.NewRequest("GET", "/user/?api_key=12345", bytes.NewBuffer([]byte{}))
	w = httptest.NewRecorder()
	ds.EXPECT().GetUserByKey(gomock.Eq("12345")).Return(nil, fmt.Errorf("User not found"))
	srv.ServeHTTP(w, r)
	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestWebIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	srv := server{templates: &templateCapture{}, ds: ds}
	r := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	ds.EXPECT().Frontpage().Return(&shared.FrontPageData{}, nil)
	srv.frontPage(w, r)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestWebLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	logger, _ := zap.NewDevelopment()
	srv := server{ds: ds, queryLimit: 10, logger: logger, templates: &templateCapture{}}
	r := httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte("username=redbo&password=test")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	user := shared.User{Username: "redbo", APIKey: "something", Password: "$2a$14$Gxs5vNqQIBLRLbra4fJBl.4QXLdXn6huraIgmyZoSpfJGULWreyaG",
		Cookie: "something", Categories: nil}
	ds.EXPECT().GetUser(gomock.Eq("redbo")).Return(&user, nil)
	srv.login(w, r)
	require.Equal(t, http.StatusSeeOther, w.Code)

	// bad password
	r = httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte("username=redbo&password=test")))
	w = httptest.NewRecorder()
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	user = shared.User{Username: "redbo", APIKey: "something", Password: "XXXXXX", Cookie: "something", Categories: nil}
	ds.EXPECT().GetUser(gomock.Eq("redbo")).Return(&user, nil)
	srv.login(w, r)
	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestWebRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	logger, _ := zap.NewDevelopment()
	srv := server{ds: ds, queryLimit: 10, logger: logger, templates: &templateCapture{}}
	r := httptest.NewRequest("POST", "/register", bytes.NewBuffer([]byte("username=redbo&password=test")))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	user := shared.User{Username: "redbo", APIKey: "something", Password: "XXXXXX", Cookie: "something", Categories: nil}
	ds.EXPECT().NewUser(gomock.Eq("redbo"), gomock.Len(60)).Return(&user, nil)
	srv.register(w, r)
	require.Equal(t, http.StatusSeeOther, w.Code)
}

func TestWebLogout(t *testing.T) {
	srv := server{}
	r := httptest.NewRequest("GET", "/logout", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	srv.logout(w, r)
	require.True(t, strings.Contains(w.Header().Get("Set-Cookie"), "session"))
}

func TestWebGifPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	logger, _ := zap.NewDevelopment()
	srv := server{ds: ds, queryLimit: 10, logger: logger, templates: &templateCapture{}}
	r := httptest.NewRequest("GET", "/gifs/1234567890abcdef", bytes.NewBuffer([]byte{}))
	r = mux.SetURLVars(r, map[string]string{"id": "1234567890abcdef"})
	user := shared.User{Username: "redbo", APIKey: "something", Password: "XXXXXX", Cookie: "something", Categories: nil}
	r = r.WithContext(context.WithValue(r.Context(), userKey, &user))
	ds.EXPECT().GIFByID(gomock.Eq("1234567890abcdef")).Return(&shared.GIF{}, nil)
	ds.EXPECT().UserGIFInfo(gomock.Eq("redbo"), gomock.Eq("1234567890abcdef")).Return(&shared.UserGIFInfo{}, nil)
	w := httptest.NewRecorder()
	srv.gifPage(w, r)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestWebRawGif(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	logger, _ := zap.NewDevelopment()
	srv := server{ds: ds, queryLimit: 10, logger: logger, templates: &templateCapture{}}
	r := httptest.NewRequest("GET", "/gif/1234567890abcdef", bytes.NewBuffer([]byte{}))
	r = mux.SetURLVars(r, map[string]string{"id": "1234567890abcdef"})
	user := shared.User{Username: "redbo", APIKey: "something", Password: "XXXXXX", Cookie: "something", Categories: nil}
	r = r.WithContext(context.WithValue(r.Context(), userKey, &user))
	ds.EXPECT().GIFData(gomock.Eq("1234567890abcdef")).Return([]byte("GIFDATA"), nil)
	w := httptest.NewRecorder()
	srv.rawGif(w, r)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "image/gif", w.Header().Get("Content-Type"))
	require.Equal(t, "GIFDATA", w.Body.String())
}

func TestWebSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	logger, _ := zap.NewDevelopment()
	srv := server{ds: ds, queryLimit: 10, logger: logger, templates: &templateCapture{}}
	r := httptest.NewRequest("GET", "/search?q=cats", bytes.NewBuffer([]byte{}))
	w := httptest.NewRecorder()
	ds.EXPECT().Search(gomock.Eq("cats"), gomock.Eq(12), gomock.Eq(0), gomock.Eq("g")).Return([]*shared.GIF{}, 14, nil)
	srv.search(w, r)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestStillGif(t *testing.T) {
	gifdata, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := NewMockDataStore(ctrl)
	logger, _ := zap.NewDevelopment()
	srv := server{ds: ds, queryLimit: 10, logger: logger, templates: &templateCapture{}}
	r := httptest.NewRequest("GET", "/still/1234567890abcdef", bytes.NewBuffer([]byte{}))
	r = mux.SetURLVars(r, map[string]string{"id": "1234567890abcdef"})
	w := httptest.NewRecorder()
	ds.EXPECT().GIFData(gomock.Eq("1234567890abcdef")).Return(gifdata, nil)
	srv.stillGif(w, r)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "image/gif", w.Header().Get("Content-Type"))
	require.Greater(t, w.Body.Len(), 40) // presumption of innocence
}
