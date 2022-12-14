# builder image
FROM golang:latest as stage-1-builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o golang-memtest .

# Generate clean and final image for the users
FROM alpine:latest
COPY --from=stage-1-builder /build/golang-memtest .
ENTRYPOINT [ "./main" ]
CMD [ "3", "300" ]
