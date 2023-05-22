# Build client side
FROM node:13.12.0-alpine as ui-build

WORKDIR /app

COPY *.json *js ./
RUN npm ci
COPY web web
RUN npm run build


# This image builds server side
FROM golang:1.20-alpine as server-build
RUN apk update && apk add --no-cache git
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest && make swagger

WORKDIR /app

COPY . ./
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o go-famtree ./cmd


# The final image
FROM alpine

WORKDIR /app

COPY --from=ui-build /app/build ./build
COPY --from=server-build /app/go-famtree ./

# Run under non-privileged user with minimal write permissions
RUN adduser -S -D -H user
USER user

CMD ["./go-famtree"]

# Heroku redefines exposed port
ENV PORT=8080
EXPOSE $PORT