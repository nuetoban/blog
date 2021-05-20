# Build image
FROM golang:1.16.4-alpine3.13 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk add --update --no-cache git && \
    mkdir -p /build

WORKDIR /build
COPY . .
RUN go build -o /build/server .


# Copy binary to alpine container
FROM alpine:3.11 AS final

RUN apk update --no-cache && \
    apk add --no-cache ca-certificates \
                       tzdata \
                       curl && \
    cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime && \
    echo "Europe/Moscow" > /etc/timezone

COPY --from=build /build/server /bin/server

EXPOSE 8090

ENTRYPOINT ["/bin/server"]

