# See: https://blog.jetbrains.com/go/2018/04/30/debugging-containerized-go-applications/
# Compile stage
FROM golang:1.10.1-alpine3.7 AS build-env
ENV CGO_ENABLED 0
ADD . /go/src/app

WORKDIR /go/src/app
ENV GOPATH /go

RUN go get -u github.com/golang/dep/...
RUN dep ensure

# The -gcflags "all=-N -l" flag helps us get a better debug experience
RUN go build -gcflags "all=-N -l" -o /server app

# Compile Delve
RUN apk add --no-cache git
RUN go get github.com/derekparker/delve/cmd/dlv

# Final stage
FROM alpine:3.7

# Port 3000 belongs to our application, 40000 belongs to Delve
EXPOSE 3000 40000

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat

WORKDIR /

COPY --from=build-env /server /
COPY --from=build-env /go/bin/dlv /

# Run delve
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/server"]