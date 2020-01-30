package giphaux

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/redbo/giphaux/shared"
	"github.com/stretchr/testify/require"
)

//go:generate mockgen -destination mockgen_test.go -package giphaux github.com/redbo/giphaux/shared DataStore

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
	gifdata, err := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
	if err != nil {
		t.Fail()
	}
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
