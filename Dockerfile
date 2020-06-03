FROM golang:1.14 as builder

WORKDIR /hortbot

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN go run github.com/markbates/pkger/cmd/pkger

ARG version
RUN go build -ldflags="-X github.com/hortbot/hortbot/internal/version.version=${version}" .

# TODO: Use distroless/static and statically compile above. (https://golang.org/issue/26492)
FROM gcr.io/distroless/base:nonroot
COPY --from=builder /hortbot/hortbot /hortbot
ENTRYPOINT [ "/hortbot" ]
