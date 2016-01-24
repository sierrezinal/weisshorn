all: build install

install:
	go install -x github.com/gophergala2016/weisshorn/cmd/weisshorn
	go install -x github.com/gophergala2016/weisshorn/cmd/weisshorn-server

build:
	go build -x github.com/gophergala2016/weisshorn
	go build -x github.com/gophergala2016/weisshorn/cmd/weisshorn
	go build -x github.com/gophergala2016/weisshorn/cmd/weisshorn-server

format:
	go fmt github.com/gophergala2016/weisshorn/cmd/weisshorn
	go fmt github.com/gophergala2016/weisshorn/cmd/weisshorn-server

platform:
	gox github.com/gophergala2016/weisshorn/cmd/weisshorn
