all: build install

install:
	go install -x github.com/sierrezinal/weisshorn/cmd/weisshorn
	go install -x github.com/sierrezinal/weisshorn/cmd/weisshorn-server

build:
	go build -x github.com/sierrezinal/weisshorn
	go build -x github.com/sierrezinal/weisshorn/cmd/weisshorn
	go build -x github.com/sierrezinal/weisshorn/cmd/weisshorn-server

format:
	go fmt github.com/sierrezinal/weisshorn/cmd/weisshorn
	go fmt github.com/sierrezinal/weisshorn/cmd/weisshorn-server

platform:
	gox github.com/sierrezinal/weisshorn/cmd/weisshorn
