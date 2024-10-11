[![Build Status](https://travis-ci.org/theosiemensrhodes/go-bktree.svg?branch=master)](https://travis-ci.org/theosiemensrhodes/go-bktree)
# Go-bktree
https://travis-ci.org/theosiemensrhodes/go-bktree.svg?branch=master
Go-bktree provides an implementation of [BK-tree](http://en.wikipedia.org/wiki/BK-tree).

# Fork Info
Forked from github.com/bahadrix/go-bktree which was forked from github.com/hjr265/go-bktree.
They both had a breaking bug which the Find method incorrectly calculates the range of child nodes to explore during a search. 
This miscalculation leads to incomplete search results, as valid words within the specified distance are missed.

This repo simply fixes that bug while the pull requests to both repos get merged.

## Installation

Install Go-bktree using the go get command:

    $ go get github.com/theosiemensrhodes/go-bktree

The only dependency is the Go distribution itself.

## Documentation

- [Reference](http://godoc.org/github.com/theosiemensrhodes/go-bktree)

## Contributing

Contributions are welcome.

## License

Go-bktree is available under the [BSD (3-Clause) License](http://opensource.org/licenses/BSD-3-Clause).