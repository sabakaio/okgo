dev:
	go get -u github.com/nsf/gocode
	go get -u github.com/rogpeppe/godef

deps:
	go get -u github.com/constabulary/gb/...
	go get -u github.com/smartystreets/goconvey
	go get github.com/codegangsta/cli
	go get github.com/gin-gonic/gin
	go get github.com/docker/libkv/store/boltdb

test:
	go test -v ./models
