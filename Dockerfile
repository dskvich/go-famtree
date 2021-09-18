# Build client side
FROM node:13.12.0-alpine as ui-build

WORKDIR /app

COPY *.json *js ./
RUN npm ci
COPY web web
RUN npm run build


# This image builds server side
FROM golang:alpine as server-build
RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . ./
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o go-famtree .

# Create user for the scratch image
RUN adduser -S -u 10001 scratchuser


# The final image
FROM scratch

WORKDIR /app

COPY --from=ui-build /app/build ./build
COPY --from=server-build /app/go-famtree ./

# Run under non-privileged user with minimal write permissions
USER 10001

CMD ["./go-famtree"]

# Heroku redefines exposed port
ENV PORT=8080
EXPOSE $PORT