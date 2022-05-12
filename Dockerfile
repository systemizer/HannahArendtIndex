FROM golang:alpine as builder
ADD . /app
WORKDIR /app/backend
RUN apk add --update gcc musl-dev
RUN go mod download
RUN CGO_ENABLED=1 go build --tags "fts5" -ldflags "-w" -a -o /main .

# Build React
FROM node:16-alpine AS node_builder
COPY --from=builder /app/frontend ./
RUN npm install
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=builder /app/backend/gorm.db ./
COPY --from=node_builder /build ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main server