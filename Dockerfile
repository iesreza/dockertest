FROM golang:onbuild AS builder

# Add Maintainer Info
LABEL maintainer="IES ITALIA Reza <reza@ies-italia.it>"

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/iesreza/dockertest
COPY . .

RUN go get ./...

RUN go build -o /go/bin/dockertest

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/hello /go/bin/hello
# Run the hello binary.
ENTRYPOINT ["/go/bin/hello"]


EXPOSE 8080

