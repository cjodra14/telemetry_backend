FROM golang:alpine as builder

RUN apk update ; apk add -U --no-cache tzdata bash upx ca-certificates musl-dev gcc

ARG PKG=telemetry_backend
ARG GITHUB_TOKEN
ARG GITHUB_USER

RUN apk update \
 && apk add git

RUN mkdir -p /go/src \
 && mkdir -p /go/bin \
 && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

WORKDIR $GOPATH/src/app
RUN go install github.com/ahmetb/govvv@latest
COPY . .

RUN rm -f go.sum

# go get uses git to pull lib dependencies
RUN git config --global url."https://$GITHUB_USER:$GITHUB_TOKEN@github.com".insteadOf "https://github.com"

RUN env GO111MODULE=on GOPRIVATE=github.com/cjodra14 go get ./...
# RUN git describe --tags --always > VERSION
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 govvv build -a -installsuffix cgo -ldflags " -s -w" -o /go/bin/$PKG main.go

RUN upx /go/bin/$PKG

FROM scratch

WORKDIR /

EXPOSE 8000
EXPOSE 3000
USER 1001

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/$PKG /$PKG

ENTRYPOINT ["/telemetry_backend"]
