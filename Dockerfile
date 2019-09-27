# go
# --- local ---
FROM golang:1.12-alpine as local
WORKDIR /go/server
COPY . .
RUN apk add --no-cache git make && \
  go get gopkg.in/urfave/cli.v2@master && \
  go get github.com/oxequa/realize && \
  make build

# --- production ---
FROM alpine as production
WORKDIR /server
COPY --from=local /go/server/bin/bot .
COPY --from=local /go/server/_tools ./_tools
RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /server/bot
EXPOSE 8080
CMD ["./bot","-c","_tools/production/config.toml"]
