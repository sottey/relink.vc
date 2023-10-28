package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

var relinkvcPath string = filepath.Join(os.Getenv("HOME"), ".rvc/relinkvc")
var bookmarksFilename string = "bookmarks"
var bookmarksFullFilename string = filepath.Join(relinkvcPath, bookmarksFilename)
var archiveFilename string = "archive"
var archiveFullFilename string = filepath.Join(relinkvcPath, archiveFilename)

func writeBookmarksToFile(filename string, bookmarks []Bookmark) {
	text, err := json.Marshal(bookmarks)
	if err != nil {
		log.Fatal(err)
	}

	if !FileExists(relinkvcPath) {
		// Create the data dir
		err := os.Mkdir(relinkvcPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

	}

	err = os.WriteFile(filename, text, 0660)
	if err != nil {
		log.Fatal(err)
	}
}

func readBookmarksFromFile(filename string) (bookmarks []Bookmark) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return []Bookmark{}
	}

	err = json.Unmarshal(text, &bookmarks)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func GetBookmarks() (bookmarks []Bookmark) {
	return readBookmarksFromFile(bookmarksFullFilename)
}

func GetArchivedBookmarks() (archive []Bookmark) {
	return readBookmarksFromFile(archiveFullFilename)
}

func GetBookmarkByUUID(UUID string) (bookmark Bookmark) {
	bookmarks := GetBookmarks()

	for _, bookmark := range bookmarks {
		if bookmark.UUID == UUID {
			return bookmark
		}
	}

	log.Fatalf("No bookmark with UUID %s exits.", UUID)

	return
}

func getHostFromURL(inputURL string) (hostname string) {
	u, err := url.Parse(inputURL)
	if err != nil {
		log.Fatal(err)
	}

	return u.Hostname()
}

func AddBookmark(url string) (title string) {
	bookmarks := GetBookmarks()

	new_uuid := uuid.New().String()
	added_ts := time.Now().Unix()
	title = GetPageTitle(url)
	hostname := getHostFromURL(url)

	bookmarks = append(bookmarks, Bookmark{
		UUID:           new_uuid,
		AddedTimestamp: added_ts,
		URL:            url,
		Title:          title,
		Hostname:       hostname,
	})

	writeBookmarksToFile(bookmarksFullFilename, bookmarks)

	return title
}

func EditBookmark(bmUuid string, url string) (title string) {
	bookmarks := GetBookmarks()

	for _, bookmark := range bookmarks {
		if bookmark.UUID == bmUuid {
			new_uuid := uuid.New().String()
			added_ts := time.Now().Unix()
			title = GetPageTitle(url)
			hostname := getHostFromURL(url)

			bookmarks = append(bookmarks, Bookmark{
				UUID:           new_uuid,
				AddedTimestamp: added_ts,
				URL:            url,
				Title:          title,
				Hostname:       hostname,
			})
			DeleteBookmark(bookmark.UUID)
		}
	}

	writeBookmarksToFile(bookmarksFullFilename, bookmarks)

	return title
}

func ArchiveBookmark(uuid string) {
	bookmarks := GetBookmarks()
	archive := GetArchivedBookmarks()

	for i, bookmark := range bookmarks {
		if bookmark.UUID == uuid {
			bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
			archive = append(archive, bookmark)
			break
		}
	}

	writeBookmarksToFile(bookmarksFullFilename, bookmarks)
	writeBookmarksToFile(archiveFullFilename, archive)
}

func DeleteBookmark(uuid string) {
	bookmarks := GetBookmarks()

	for i, bookmark := range bookmarks {
		if bookmark.UUID == uuid {
			bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
			break
		}
	}

	writeBookmarksToFile(bookmarksFullFilename, bookmarks)
}
