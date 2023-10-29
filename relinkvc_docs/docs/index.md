# relink.vc
---
relink.vc is a tool for command line fans that allows you to store and organize bookmarks. It is a full featured bookmark manager all stored in a JSON file so it is super portable and tooling new apps for the data format is super easy.

## Getting Started
To learn about available relink.vc commands, see the [usage page](./usage/)

## Features
* Save links with fetched site titles
* Links can be archived without deleting
* Edit URLs
* Delete URLs
* Built in web server to serve up web site of links

## Future Plans
* Store custom link tags
* Store custom link categories
* Store custom link metadata
* Ability to change ip address and port of web server
* Ability to specify alternate rowser from default
* Theming


## Downloading
Find the installation for your device in the releases section of the GitHub project [here](https://github.com/sottey/relink.vc/releases)


## Building

Clone locally and run the following to create a local binary
```
go build
```

OR (using rake) you can build a release using
```
rake build
```

## Shoutouts

* [bm](https://github.com/kcwebapply/bm) - HUGE thanks to everyone who created [bm](https://github.com/kcwebapply/bm), a great CLI bookmark tool!


## License

Relink.vc uses the MIT License. Please see the [relink.vc license](https://github.com/sottey/relink.vc/blob/main/LICENSE) for details