# builder image
FROM golang:1.20-alpine as builder
WORKDIR /build
COPY . /build
RUN go build -o api-gateway .

# generate clean, final image for end users
FROM alpine:3.18.0 as hoster
COPY --from=builder /build/api-gateway ./api-gateway
COPY --from=builder /build/.env ./.env
# executable
ENTRYPOINT [ "./api-gateway" ]