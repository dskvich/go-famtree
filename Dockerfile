FROM node:13.12.0-alpine as frontend_builder
RUN mkdir /app
ADD web /app
WORKDIR /app
RUN npm i
RUN npm run build

FROM golang:alpine as backend_builder
RUN apk update && apk add --no-cache git
RUN mkdir /app 
ADD . /app/
WORKDIR /app
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o go-famtree .
RUN adduser -S -D -H -u 10001 appuser

FROM scratch
COPY --from=frontend_builder /app/build /app/web/build
COPY --from=backend_builder /app/go-famtree /app
USER 10001
WORKDIR /app
EXPOSE 8080
CMD ["./go-famtree"]