# Build
FROM golang:1.14
WORKDIR $GOPATH/src/github.com/thenets/docker-svg-term
COPY . .
RUN go get -d -v ./...
RUN go build -o /tmp/svg-term-server
RUN chmod +x /tmp/svg-term-server

# Server
FROM node:12
WORKDIR /app
RUN npm install -g svg-term-cli
COPY --from=0 /tmp/svg-term-server /app/svg-term-server
COPY ./static/ /app/static/
EXPOSE 8080
ENTRYPOINT [ "/app/svg-term-server" ]
