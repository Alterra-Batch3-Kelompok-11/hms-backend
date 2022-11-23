FROM golang:1.19 as blender

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app .

FROM alpine:latest

COPY --from=blender /app/main.app .

CMD ["/main.app"]