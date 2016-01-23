all: build install

install:
	go install -x github.com/gophergala2016/weisshorn/cmd/weisshorn
	go install -x github.com/gophergala2016/weisshorn/cmd/weisshorn-server

build:
	go build -x github.com/gophergala2016/weisshorn/cmd/weisshorn
	go build -x github.com/gophergala2016/weisshorn/cmd/weisshorn-server
