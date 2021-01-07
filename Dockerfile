# Golang image for building the project
FROM golang:alpine as builder

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

RUN apk --no-cache add git dep ca-certificates
RUN apk --no-cache add tzdata

RUN mkdir -p $GOPATH/src/github.com/pzentenoe/onboarding-golang

WORKDIR $GOPATH/src/github.com/pzentenoe/onboarding-golang

COPY go.mod .
COPY app/ app/
COPY vendor vendor/

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/main ./app/main.go


# Runtime image with scratch container
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/bin/ /app/

ENV TZ America/Santiago

ENTRYPOINT ["/app/main"]