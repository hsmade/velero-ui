FROM node:21.3.0-alpine as node
COPY web /app/web
WORKDIR /app/web
RUN yarn install
RUN yarn build

FROM golang:1.21-alpine as go
COPY api /app/api
COPY cmd /app/cmd
COPY internal /app/internal
COPY go.mod go.sum /app/
WORKDIR /app
RUN go build cmd/velero-ui/main.go

FROM scratch
COPY --from=node /app/web/dist /web/dist
COPY --from=go /app/main /velero-ui
ENTRYPOINT ["/velero-ui"]
