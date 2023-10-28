# relinkvc [![build status](https://travis-ci.org/jamesbvaughan/relinkvc.svg)](https://travis-ci.org/jamesbvaughan/relinkvc)

relinkvc is a bookmarking tool, similar to
[Pocket](https://getpocket.com/).
It has a simple command line interface as well as a web interface.

## Installation

The best way to install relinkvc right now is to

1. Download a precompiled binary from [the releases page](https://github.com/jamesbvaughan/relinkvc/releases).
2. Place it in your `PATH`.
3. Make it executable.

Once you've done that you're ready to go!
Just run `relinkvc add https://github.com/jamesbvaughan/relinkvc`
to add your first bookmark!

Alternatively, you can use relinkvc from the web interface, which looks like this:
![web interface screenshot](https://raw.githubusercontent.com/jamesbvaughan/relinkvc/master/misc/web-ui.png "relinkvc web interface")
and can be served locally by running `relinkvc serve`.

## Usage

### Add a bookmark

```sh
relinkvc add <URL>
```

### List bookmarks

```sh
relinkvc list
```

### Open a bookmark

```sh
relinkvc open <ID>
```

### Archive a bookmark

```sh
relinkvc archive <ID>
```

### Start the webserver for the web interface

```sh
relinkvc serve
```

### Permanently delete a bookmark

```sh
relinkvc delete <ID>
```

### Get help

```sh
relinkvc help
```

### Get help for a specific command

```sh
relinkvc help <COMMAND>
```

## [Licence](LICENSE)

The MIT License (MIT)

Copyright (c) 2018 James Vaughan
