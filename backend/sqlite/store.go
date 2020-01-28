package sqlite

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/redbo/giphaux/shared"
	gormzap "github.com/wantedly/gorm-zap"
	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // anonymous import just makes the sqlite driver available to gorm
)

// This uses gorm, which I wanted to get some experience with.
// It turns out gorm has a lot of sharp edges and I can't enthusiastically recommend it.

// gifToGIF coerces a database GIF record into the data structure dealt with by the API.
func gifToGIF(src *GIF) *shared.GIF {
	dst := &shared.GIF{
		ID:      src.APIID,
		Caption: src.Caption,
		Tags:    strings.Split(src.Tags, ","),
		// TODO: so much
	}
	return dst
}

// userToUser coerces a database user object into the data structure used by the rest of the app.
func userToUser(src *User) *shared.User {
	categories := []string{}
	for _, cat := range src.Categories {
		categories = append(categories, cat.Title)
	}
	return &shared.User{
		Username:   src.Username,
		APIKey:     src.APIKey,
		Password:   src.Password,
		Cookie:     src.Cookie,
		Categories: categories,
	}
}

// sqlDataStore implements the giphaux/backend.DataStore interface for a sqlite database.
type sqlDataStore struct {
	db *gorm.DB
}

// RandomID returns a random hex string.  Enjoy.
func (s *sqlDataStore) RandomID() string {
	return fmt.Sprintf("%016x%016x", rand.Int63(), rand.Int63())
}

// NewUser creates a user with the given credentials
func (s *sqlDataStore) NewUser(username, password string) (*shared.User, error) {
	dbuser := new(User)
	if err := s.db.Create(&User{
		Username: username,
		Password: password,
		APIKey:   s.RandomID(),
		Cookie:   s.RandomID(),
	}).Scan(&dbuser).Error; err != nil {
		return nil, fmt.Errorf("Error creating user: %w", err)
	}
	return userToUser(dbuser), nil
}

// usernameToID is a utility function for grabbing a user's database ID by username
func usernameToID(db *gorm.DB, username string) (uint, error) {
	dbuser := &struct{ ID uint }{}
	if err := db.Table("users").Select("id").Where("username = ?", username).Scan(dbuser).Error; err != nil {
		return 0, fmt.Errorf("Unable to find user: %w", err)
	}
	return dbuser.ID, nil
}

// gifToID is a utility function for grabbing a gif's database ID by gifid
func gifToID(db *gorm.DB, gifid string) (uint, error) {
	dbgif := new(GIF)
	if err := db.Table("gifs").Select("id").Where("api_id = ?", gifid).Scan(dbgif).Error; err != nil {
		return 0, fmt.Errorf("Unable to find gif: %s", err)
	}
	return dbgif.ID, nil
}

// GetUser returns a user based on their username
func (s *sqlDataStore) GetUser(username string) (*shared.User, error) {
	dbuser := new(User)
	if err := s.db.Model(dbuser).Where("username = ?", username).Scan(dbuser).Error; err != nil {
		return nil, fmt.Errorf("Error finding user by username: %w", err)
	}
	if dbuser.Username == "" {
		return nil, fmt.Errorf("Unable to find user")
	}
	if err := s.db.Table("categories").Where("user_id = ?", dbuser.ID).Find(&dbuser.Categories).Error; err != nil {
		return nil, fmt.Errorf("Error fetching user categories: %w", err)
	}
	return userToUser(dbuser), nil
}

// GetUserByKey looks up a user based on their API key.
func (s *sqlDataStore) GetUserByKey(key string) (*shared.User, error) {
	dbuser := new(User)
	if err := s.db.Model(dbuser).Where("api_key = ?", key).Scan(dbuser).Error; err != nil {
		return nil, fmt.Errorf("Error fetching user by key: %w", err)
	}
	if dbuser.Username == "" {
		return nil, fmt.Errorf("Unable to find user")
	}
	if err := s.db.Table("categories").Where("user_id = ?", dbuser.ID).Find(&dbuser.Categories).Error; err != nil {
		return nil, fmt.Errorf("Error fetching user categories: %w", err)
	}
	return userToUser(dbuser), nil
}

// GetUserByCookie looks up a user based on their session id.
func (s *sqlDataStore) GetUserByCookie(cookie string) (*shared.User, error) {
	var dbuser User
	if err := s.db.Model(dbuser).Where("cookie = ?", cookie).First(&dbuser).Error; err != nil {
		return nil, fmt.Errorf("Error fetching user by cookie: %w", err)
	}
	if dbuser.Username == "" {
		return nil, fmt.Errorf("Unable to find user")
	}
	if err := s.db.Table("categories").Where("user_id = ?", dbuser.ID).Find(&dbuser.Categories).Error; err != nil {
		return nil, fmt.Errorf("Error fetching user categories: %w", err)
	}
	return userToUser(&dbuser), nil
}

// UserFrontpage returns the lists of gifs and things needed to show the user's home page.
func (s *sqlDataStore) UserFrontpage(username string) (*shared.FrontPageData, error) {
	fp := &shared.FrontPageData{Categories: make(map[string][]*shared.GIF)}
	userID, err := usernameToID(s.db, username)
	if err != nil {
		return nil, fmt.Errorf("Unable to find user: %w", err)
	}
	addResults := func(title string, results []GIF) {
		fp.Categories[title] = make([]*shared.GIF, 0)
		for _, gif := range results {
			fp.Categories[title] = append(fp.Categories[title], gifToGIF(&gif))
		}
	}
	gifs := []GIF{}
	if err := s.db.Table("gifs").Limit(8).Order("id desc").Where("user_id = ?", userID).Scan(&gifs).Error; err != nil {
		return nil, fmt.Errorf("Error getting uploads: %w", err)
	}
	addResults("Uploads", gifs)

	gifs = gifs[:]
	if err := s.db.Raw("SELECT g.* FROM gifs g JOIN favorites f ON f.gif_id=g.id WHERE g.user_id=? ORDER BY f.id DESC LIMIT 8", userID).Scan(&gifs).Error; err != nil {
		return nil, fmt.Errorf("Error getting favorites: %w", err)
	}
	addResults("Favorites", gifs)

	cats := []Category{}
	if err := s.db.Table("categories").Where("user_id = ?", userID).Scan(&cats).Error; err != nil {
		return nil, fmt.Errorf("Error fetching categories: %w", err)
	}
	for _, cat := range cats {
		gifs = gifs[:]
		if err := s.db.Raw(`SELECT g.* FROM (gifs g JOIN favorites f ON f.gif_id = g.id) JOIN categorized_favorites cf on f.id = cf.favorite_id
					WHERE cf.category_id = ? LIMIT 8`, cat.ID).Scan(&gifs).Error; err != nil {
			return nil, fmt.Errorf("Error fetching categorized favorites: %w", err)
		}
		addResults(cat.Title, gifs)
	}
	return fp, nil
}

// GifByID returns the gif with the given API ID (not numeric database row ID).
func (s *sqlDataStore) GIFByID(id string) (*shared.GIF, error) {
	dbgif := new(GIF)
	if r := s.db.Model(dbgif).Where("api_id = ?", id).Scan(dbgif); r.Error != nil {
		return nil, fmt.Errorf("Error fetching gif: %w", r.Error)
	}
	return gifToGIF(dbgif), nil
}

// RemoveFavorite removes a gif from the user's favorites.
func (s *sqlDataStore) RemoveFavorite(username string, gifid string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		userID, err := usernameToID(tx, username)
		if err != nil {
			return fmt.Errorf("Unable to find user: %w", err)
		}
		gifID, err := gifToID(tx, gifid)
		if err != nil {
			return fmt.Errorf("Unable to find gif: %w", err)
		}
		favorite := new(Favorite)
		if db := tx.Model(favorite).Where("user_id = ? AND gif_id = ?", userID, gifID).Scan(favorite); db.Error != nil {
			return fmt.Errorf("Unable to find category: %w", db.Error)
		}
		if db := tx.Delete(&CategorizedFavorite{}, "favorite_id = ?", favorite.ID); db.Error != nil {
			return fmt.Errorf("Unable to delete assocations: %w", db.Error)
		}
		return tx.Delete(favorite, "id = ?", favorite.ID).Error
	})
}

// AddFavorite adds a gif to the user's favorites.
func (s *sqlDataStore) AddFavorite(username string, gifid string, categories []string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		userID, err := usernameToID(tx, username)
		if err != nil {
			return fmt.Errorf("Unable to find user: %w", err)
		}
		gifID, err := gifToID(tx, gifid)
		if err != nil {
			return fmt.Errorf("Unable to find gif: %w", err)
		}
		fav := new(Favorite)
		if err := tx.Create(&Favorite{UserID: userID, GIFID: gifID}).Scan(&fav).Error; err != nil {
			return fmt.Errorf("Error saving favorite: %w", err)
		}
		return tx.Exec(`INSERT INTO categorized_favorites (favorite_id, category_id)
						SELECT ?, id FROM categories WHERE title IN (?) AND user_id=?`, fav.ID, categories, userID).Error
	})
}

// UserGIFInfo returns the user's info for a given gif - whether it's favorited and what categories they have it in etc.
func (s *sqlDataStore) UserGIFInfo(username string, gifid string) (*shared.UserGIFInfo, error) {
	userID, err := usernameToID(s.db, username)
	if err != nil {
		return nil, fmt.Errorf("Unable to find user: %w", err)
	}
	dbgif := new(GIF)
	dbfav := new(Favorite)
	if err := s.db.Model(dbgif).Where("api_id = ?", gifid).Scan(dbgif).Error; err != nil {
		return nil, fmt.Errorf("Error fetching gif: %w", err)
	}
	favorited := false
	categories := []struct {
		Title     string
		Favorited bool
	}{}
	if r := s.db.Model(dbfav).Where("user_id = ? AND gif_id = ?", userID, dbgif.ID).Scan(dbfav); r.Error == nil {
		favorited = true
		s.db.Raw(`SELECT c.Title as title,cf.ID,cf.ID IS NOT NULL as favorited
					FROM categories c LEFT JOIN categorized_favorites cf ON (cf.category_id=c.id AND cf.favorite_id = ?)
					WHERE c.user_id = ?`,
			dbfav.ID, userID).Scan(&categories)
	}
	cats := make(map[string]bool)
	for _, c := range categories {
		cats[c.Title] = c.Favorited
	}
	return &shared.UserGIFInfo{
		Owner:      userID == dbgif.UserID,
		Favorited:  favorited,
		Categories: cats,
	}, nil
}

// GIFsByID returns a list of gifs with the given IDs.
func (s *sqlDataStore) GIFsByID(ids []string) ([]*shared.GIF, error) {
	dbgifs := []GIF{}
	if err := s.db.Model(&GIF{}).Where("api_id IN ?", ids).Scan(&dbgifs).Error; err != nil {
		return nil, fmt.Errorf("Error fetching gif: %w", err)
	}
	gifs := []*shared.GIF{}
	for _, gif := range dbgifs {
		gifs = append(gifs, gifToGIF(&gif))
	}
	return gifs, nil
}

// Search searches gifs for the given query.
func (s *sqlDataStore) Search(query string, limit int, offset int, rating string) ([]*shared.GIF, int, error) {
	dbgifs := []GIF{}
	documentCount := struct{ Count int }{0}
	if err := s.db.Table("gifsearch").Select("COUNT(*) as count").
		Where("gifsearch MATCH ? AND rating MATCH ?", query, rating).Scan(&documentCount).Error; err != nil {
		return nil, 0, fmt.Errorf("Error getting search result count: %w", err)
	}
	if err := s.db.Table("gifsearch").Select("gifs.*").Joins("JOIN gifs").
		Where("gifsearch MATCH ? AND rating MATCH ?", query, rating).Where("gifs.id = gifsearch.docid").
		Limit(limit).Offset(offset).Scan(&dbgifs).Error; err != nil {
		return nil, 0, fmt.Errorf("Error getting search results: %w", err)

	}
	gifs := []*shared.GIF{}
	for _, gif := range dbgifs {
		gifs = append(gifs, gifToGIF(&gif))
	}
	return gifs, documentCount.Count, nil
}

// Trending returns the most recent trending gifs.
func (s *sqlDataStore) Trending(limit int, off int, rating string) ([]*shared.GIF, int, error) {
	documentCount := struct{ Count int }{0}
	if err := s.db.Table("gifs").Select("COUNT(*) as count").Where("trending IS NOT NULL").Scan(&documentCount).Error; err != nil {
		return nil, 0, fmt.Errorf("Error getting trending count: %w", err)
	}
	gifs := []GIF{}
	if err := s.db.Table("gifs").Limit(limit).Offset(off).
		Where("trending IS NOT NULL").Order("trending desc").Scan(&gifs).Error; err != nil {
		return nil, 0, fmt.Errorf("Error fetching gifs: %w", err)
	}
	rgifs := make([]*shared.GIF, 0)
	for _, gif := range gifs {
		rgifs = append(rgifs, gifToGIF(&gif))
	}
	return rgifs, documentCount.Count, nil
}

// RandomSearch returns a random GIF that matches the query. I don't know what to do with weirdness.
func (s *sqlDataStore) RandomSearch(q string, weirdness int) (*shared.GIF, error) {
	dbgif := GIF{}
	documentCount := struct{ Count int64 }{0}
	if err := s.db.Table("gifsearch").Select("COUNT(*) as count").
		Where("gifsearch MATCH ?", q).Scan(&documentCount).Error; err != nil {
		return nil, fmt.Errorf("Error getting random search count: %w", err)
	}
	if err := s.db.Table("gifsearch").Limit(1).Offset(rand.Int63()%documentCount.Count).
		Select("gifs.*").Joins("JOIN gifs").
		Where("gifsearch MATCH ?", q).
		Where("gifs.id = gifsearch.docid").Scan(&dbgif).Error; err != nil {
		return nil, fmt.Errorf("Error getting random search results: %w", err)
	}
	return gifToGIF(&dbgif), nil
}

// RandomByTag returns a random image with the given tag.
func (s *sqlDataStore) RandomByTag(tag string, rating string) (*shared.GIF, error) {
	return s.RandomSearch("tag:"+tag, 1)
}

// AddCategory adds a category to the user's profile.
func (s *sqlDataStore) AddCategory(username string, title string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		userID, err := usernameToID(tx, username)
		if err != nil {
			return fmt.Errorf("Unable to find user: %w", err)
		}
		return tx.Create(&Category{
			UserID: uint(userID),
			Title:  title,
		}).Error
	})
}

// RemoveCategory removes a user's category.
func (s *sqlDataStore) RemoveCategory(username string, title string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		userID, err := usernameToID(tx, username)
		if err != nil {
			return fmt.Errorf("Unable to find user: %w", err)
		}
		category := new(Category)
		if err := tx.Model(category).Where("user_id = ? AND title = ?", userID, title).Scan(category).Error; err != nil {
			return fmt.Errorf("Unable to find category: %w", err)
		}
		if err := tx.Delete(&CategorizedFavorite{}, "category_id = ?", category.ID).Error; err != nil {
			return fmt.Errorf("Unable to delete assocations: %w", err)
		}
		return tx.Delete(&Category{}, "id = ?", category.ID).Error
	})
}

// AddGIF adds a GIF to the database.
func (s *sqlDataStore) AddGIF(username, caption string, tags, cats []string, sourceURL, rating string) (*shared.GIF, error) {
	t := time.Now()
	fav := new(Favorite)
	gif := new(GIF)
	err := s.db.Transaction(func(tx *gorm.DB) error {
		userID, err := usernameToID(tx, username)
		if err != nil {
			return fmt.Errorf("Unable to find user: %w", err)
		}
		if err := tx.Create(&GIF{
			Type:             "gif",
			APIID:            fmt.Sprintf("%016x", rand.Int63()),
			ImportDatetime:   &t,
			TrendingDatetime: nil,
			Rating:           rating,
			Source:           sourceURL,
			Caption:          caption,
			Tags:             strings.Join(tags, ","),
			UserID:           userID,
			ContentURL:       "",
		}).Scan(&gif).Error; err != nil {
			return fmt.Errorf("Error creating gif: %w", err)
		}
		// automatically favorite the GIF on upload.
		if err := tx.Create(&Favorite{UserID: userID, GIFID: gif.ID}).Scan(&fav).Error; err != nil {
			return fmt.Errorf("Error creating favorite: %w", err)
		}
		if len(cats) != 0 {
			categories := []Category{}
			if err := s.db.Table("categories").Where("user_id = ?", userID).Where("Title in (?)", cats).Scan(&categories).Error; err != nil {
				return fmt.Errorf("Unable to load categories: %w", err)
			}
			for _, cat := range categories {
				if err := tx.Create(&CategorizedFavorite{
					FavoriteID: fav.ID,
					CategoryID: cat.ID,
				}).Error; err != nil {
					return fmt.Errorf("Unable to save new gif to categories: %w", err)
				}
			}
		}
		return nil
	})
	return gifToGIF(gif), err
}

// RemoveGIF removes a gif from the database entirely.
func (s *sqlDataStore) RemoveGIF(username string, gifid string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		userID, err := usernameToID(tx, username)
		if err != nil {
			return fmt.Errorf("Unable to find user: %w", err)
		}
		dbgif := new(GIF)
		if r := tx.Model(dbgif).Where("api_id = ? AND user_id = ?", gifid, userID).Scan(dbgif); r.Error != nil {
			return fmt.Errorf("Error finding gif: %w", err)
		}
		if err := tx.Exec(`DELETE FROM categorized_favorites
							WHERE favorite_id IN (SELECT id from favorites WHERE gif_id = ?)`, dbgif.ID).Error; err != nil {
			return fmt.Errorf("Unable to remove categorizations for gif: %w", err)
		}
		if err := tx.Exec(`DELETE FROM favorites WHERE gif_id = ?`, dbgif.ID).Error; err != nil {
			return fmt.Errorf("Unable to remove favorites for gif: %w", err)
		}
		if err := tx.Exec(`DELETE FROM gifs WHERE id = ?`, dbgif.ID).Error; err != nil {
			return fmt.Errorf("Unable to delete GIF: %w", err)
		}
		return nil
	})
}

// UpdateCategories updates the user's categorizations for a favorited gif.
func (s *sqlDataStore) UpdateCategories(username string, gifid string, categories []string) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		userID, err := usernameToID(tx, username)
		if err != nil {
			return fmt.Errorf("Unable to find user: %w", err)
		}
		gifID, err := gifToID(tx, gifid)
		if err != nil {
			return fmt.Errorf("Unable to find gif: %w", err)
		}
		favorite := new(Favorite)
		if err := tx.Model(favorite).Where("user_id = ? AND gif_id = ?", userID, gifID).Scan(favorite).Error; err != nil {
			return fmt.Errorf("Unable to find favorite: %w", err)
		}
		if err := tx.Delete(&CategorizedFavorite{}, "favorite_id = ?", favorite.ID).Error; err != nil {
			return fmt.Errorf("Unable to delete categorized favorites: %w", err)
		}
		return tx.Exec(`INSERT INTO categorized_favorites (favorite_id, category_id)
						SELECT ?, id FROM categories WHERE title IN (?) AND user_id=?`, favorite.ID, categories, userID).Error
	})
	return err
}

// make sure that our implementation satisfies the interface
var _ shared.DataStore = &sqlDataStore{}

// OpenStore returns a giphaux.DataStore connected to the named database file.
func OpenStore(settings *shared.Configuration, logger *zap.Logger) (shared.DataStore, error) {
	db, err := gorm.Open("sqlite3", settings.Database)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %w", err)
	}
	if settings.Verbose {
		db.LogMode(true)
		db.SetLogger(gormzap.New(logger))
	}

	// Have gorm automatically create these tables.
	// This would probably be a terrible idea in a production setting.
	db.AutoMigrate(&User{}, &GIF{}, &Category{}, &Favorite{}, &CategorizedFavorite{})

	// Manually create the text search table and triggers to update it because
	// gorm doesn't know how to do any of that.
	// fts4 was giving me "logic errors" when I fiddle with docid but fts4 works. shruggy guy emoji.
	db.Exec(`CREATE VIRTUAL TABLE IF NOT EXISTS gifsearch USING fts3(content="gifs",
				caption, tag, rating)`)
	db.Exec(`CREATE TRIGGER IF NOT EXISTS gifs_bu BEFORE UPDATE ON gifs BEGIN
				DELETE FROM gifsearch WHERE docid=old.rowid;
			END;`)
	db.Exec(`CREATE TRIGGER IF NOT EXISTS gifs_bd BEFORE DELETE ON gifs BEGIN
				DELETE FROM gifsearch WHERE docid=old.rowid;
			END;`)
	db.Exec(`CREATE TRIGGER IF NOT EXISTS gifs_au AFTER UPDATE ON gifs
			BEGIN
				INSERT INTO gifsearch(docid, caption, tag, rating)
				VALUES(new.rowid, new.caption, new.tags, new.rating);
			END;`)
	db.Exec(`CREATE TRIGGER IF NOT EXISTS gifs_ai AFTER INSERT ON gifs
			BEGIN
				INSERT INTO gifsearch(docid, caption, tag, rating)
				VALUES(new.rowid, new.caption, new.tags, new.rating);
			END;`)
	return &sqlDataStore{db}, nil
}
