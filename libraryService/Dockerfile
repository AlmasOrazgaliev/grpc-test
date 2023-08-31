# builder image
FROM golang:1.20-alpine as builder
WORKDIR /build
COPY . /build
RUN go build -o libraryService .

# generate clean, final image for end users
FROM alpine:3.18.0 as hoster
COPY --from=builder /build/libraryService ./libraryService
COPY --from=builder /build/.env* ./.env
# executable
ENTRYPOINT [ "./libraryService" ]