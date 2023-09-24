FROM golang:1.21

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV PORT 8000

WORKDIR $GOPATH/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o $GOPATH/bin/app . && ls -la $GOPATH/bin

EXPOSE $PORT

CMD ["app"]