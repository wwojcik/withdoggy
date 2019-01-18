FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git tzdata && update-ca-certificates
RUN adduser -D -g '' withdoggy
RUN mkdir /withdoggy
WORKDIR /withdoggy
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN ./hack/build.sh

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder  /withdoggy/bin/* /withdoggy/bin/
USER withdoggy