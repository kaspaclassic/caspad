ARG CASPAD_IMAGE
ARG CASPADMINER_IMAGE

FROM ${CASPAD_IMAGE} as caspad
FROM ${CASPADMINER_IMAGE} as casminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/kaspaclassic/caspad

WORKDIR /go/src/github.com/kaspaclassic/caspad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=caspad /app/ /app/
COPY --from=casminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/kaspaclassic/caspad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
