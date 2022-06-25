FROM golang:1.16 as gobuild
ARG VERSION=latest

WORKDIR /go/src/github.com/myk4040okothogodo/EvolveAPI
ADD go.mod go.sum main.go  ./
ADD vendor      ./vendor
ADD pkg         ./pkg
ADD docs        ./docs


RUN CGO_ENABLED=0 GOOS=linux  GO111MODULE=on go build -mod=vendor -o EvolveAPI -ldflags "-X main.version=$VERSION" main.go

FROM gcr.io/distroless/base

COPY --from=gobuild /go/src/github.com/myk4040okothogodo/EvolveAPI/EvolveAPI   /bin

ENTRYPOINT["/bin/EvolveAPI"]
