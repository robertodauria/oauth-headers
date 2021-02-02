FROM golang:1.15 as build
ENV CGO_ENABLED 0
ADD . /go/src/github.com/robertodauria/oauth-headers
WORKDIR /go/src/github.com/robertodauria/oauth-headers
RUN go get \
    -v \
    -ldflags "-X github.com/m-lab/go/prometheusx.GitShortCommit=$(git log -1 --format=%h)" \
    github.com/robertodauria/oauth-headers/cmd/oauth-headers

# Now copy the built image into the minimal base image
FROM alpine:3.12
RUN apk add ca-certificates
COPY --from=build /go/bin/oauth-headers /
WORKDIR /
ENTRYPOINT ["/oauth-headers"]