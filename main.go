package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "relinkvc"
	app.Version = "0.5.0"
	app.Compiled = time.Now()
	app.HelpName = "relinkvc"
	app.Usage = "bookmark things like you mean it"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"new"},
			Usage:   "add a bookmark",
			Action: func(c *cli.Context) {
				bookmarks := GetBookmarks()
				url := c.Args().First()
				title := AddBookmark(url)

				fmt.Printf("added bookmark %d: %s\n", len(bookmarks)+1, title)
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l", "ls"},
			Usage:   "list bookmarks",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "urls, u",
					Usage: "print bookmark URLs",
				},
				cli.BoolFlag{
					Name:  "archive, a",
					Usage: "print archived bookmarks",
				},
			},
			Action: func(c *cli.Context) {
				var bookmarks []Bookmark

				if c.Bool("archive") {
					bookmarks = GetArchivedBookmarks()
				} else {
					bookmarks = GetBookmarks()
				}

				PrintBookmarkTable(bookmarks, c.Bool("urls"), !c.Bool("archive"))
			},
		},
		{
			Name:    "open",
			Aliases: []string{"o"},
			Usage:   "open a bookmark",
			Action: func(c *cli.Context) {
				index := c.Args().First()
				bookmark := GetBookmarkAtIndex(index)

				fmt.Printf("opening \"%s\"...\n", bookmark.Title)

				err := open.Run(bookmark.URL)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:    "archive",
			Aliases: []string{"a"},
			Usage:   "archive a bookmark",
			Action: func(c *cli.Context) {
				index := c.Args().First()
				bookmark := GetBookmarkAtIndex(index)

				ArchiveBookmark(bookmark.UUID)

				fmt.Printf("archived bookmark: \"%s\"\n", bookmark.Title)
			},
		},
		{
			Name:    "edit",
			Usage:   "edit a bookmark",
			Aliases: []string{"e"},
			Action: func(c *cli.Context) {
				index := c.Args().Get(0)
				url := c.Args().Get(1)
				bookmark := GetBookmarkAtIndex(index)

				EditBookmark(bookmark.UUID, url)

			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d", "del", "rm"},
			Usage:   "permanently delete a bookmark",
			Action: func(c *cli.Context) {
				index := c.Args().First()
				bookmark := GetBookmarkAtIndex(index)

				DeleteBookmark(bookmark.UUID)

				fmt.Printf("deleted bookmark: \"%s\"\n", bookmark.Title)
			},
		},
		{
			Name:    "serve",
			Aliases: []string{"sv"},
			Usage:   "start the webserver",
			Action: func(c *cli.Context) {
				Serve()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
