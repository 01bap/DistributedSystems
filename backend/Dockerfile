ARG CODE_VERSION=tip-alpine3.20
FROM golang:${CODE_VERSION} AS builder

ENV PORT=3004

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/itemStore


EXPOSE $PORT
CMD ["./itemStore"]


# ENTRYPOINT ["./app"]