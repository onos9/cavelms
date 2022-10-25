###############################
# STEP 2 Build services image
###############################
FROM golang:1.19-alpine AS builder

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV GOARCH=amd64

WORKDIR /src

# System dependencies
RUN apk update && apk upgrade \
    && apk add --no-cache ca-certificates

COPY go.* /src/

# Fetch dependencies.
RUN  go mod tidy \
    && go mod download \
    && go mod verify

COPY ./ .

# Buid for production
RUN go run github.com/99designs/gqlgen && go run graph/plugin/custom_tags.go && go mod tidy
RUN cd cmd && go build -gcflags "all=-N -l" -o ../cavelms

################################
# STEP 3 build a small image for backend
################################
FROM scratch AS server

COPY --from=builder /src/cavelms ./cavelms
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./cavelms"]