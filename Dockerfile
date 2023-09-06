FROM node:20.6.0-alpine as node
COPY web /app/web
WORKDIR /app/web
RUN yarn install
RUN yarn build

FROM golang:1.21-alpine as go
COPY api /app/api
COPY internal /app/internal
COPY util /app/util
COPY main.go go.mod go.sum /app/
WORKDIR /app
RUN go build main.go

FROM scratch
COPY --from=node /app/web/dist /web/dist
COPY --from=go /app/main /
ENTRYPOINT ["/main"]
