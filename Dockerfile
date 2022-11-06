###########
# builder #
###########

FROM golang:1.19-buster AS builder
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    upx-ucl

WORKDIR /build
COPY . .
RUN go install github.com/cosmtrek/air@latest

RUN GO111MODULE=on CGO_ENABLED=0 go build \
      -ldflags='-w -s -extldflags "-static"' \
      -o ./bin/golife \
 && upx-ucl --best --ultra-brute ./bin/golife

###########
# release #
###########
# CLI実行用

FROM golang:1.19-buster AS release
COPY --from=builder /build/bin/golife /bin/
COPY --from=builder /build/world.txt /workdir/world.txt
WORKDIR /workdir
CMD ["golife", "cli"]

##########
# Heroku #
##########
# Server実行用

FROM golang:1.19-buster as heroku
COPY --from=builder /build/bin/golife /bin/
 # まだ初期ワールドをファイルから読み取ってるので必要
COPY --from=builder /build/world.txt /workdir/world.txt
WORKDIR /workdir
# herokuではportが変動するため、合わせる必要がある
CMD PORT=$PORT golife server
