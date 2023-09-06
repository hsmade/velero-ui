FROM node:20.6.0-alpine as node
COPY web /app/web
WORKDIR /app/web
RUN yarn install
RUN yarn build

FROM golang:1.21-alpine as go
COPY api /app/api
COPY cmd /app/cmd
COPY internal /app/internal
COPY util /app/util
COPY go.mod go.sum /app/
WORKDIR /app
RUN go build cmd/velero-ui

FROM scratch
COPY --from=node /app/web/dist /web/dist
COPY --from=go /app/velero-ui /
ENTRYPOINT ["/velero-ui"]
