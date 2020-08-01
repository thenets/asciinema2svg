IMAGE_TAG=thenets/svg-term

dev:
	@go run ./main.go

install-build-dependencies:
	go tool dist install -v pkg/runtime
	go install -v -a std

build: build-linux build-windows

build-windows:
	mkdir -p ./dist/windows/
	GOARCH=amd64 \
	GOOS=windows \
	go build -o ./dist/windows/svg-term-server.exe


build-linux:
	mkdir -p ./dist/linux/
	GOARCH=amd64 \
	GOOS=linux \
	go build -o ./dist/linux/svg-term-server

docker-build:
	docker build -t $(IMAGE_TAG) .

docker-run:
	docker run -it --rm \
		--name svg-term-server \
		-p 8080:8080 \
		$(IMAGE_TAG)
