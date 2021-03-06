FROM golang:1.16.0 AS movies_builder
WORKDIR /opt/movies
COPY . /opt/movies
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make -B depends
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make -B storage
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make -B parser
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make -B bot



FROM alpine:3.13.2 as movies_storage
WORKDIR /opt/movies
COPY --from=movies_builder /opt/movies/build/storage .
CMD ["./storage"]



FROM alpine:3.13.2 as movies_parser
WORKDIR /opt/movies
COPY --from=movies_builder /opt/movies/build/parser .
CMD ["./parser"]



FROM alpine:3.13.2 as movies_bot
WORKDIR /opt/movies
COPY --from=movies_builder /opt/movies/build/bot .
CMD ["./bot"]
