ARG PYIPAD_IMAGE
ARG PYIPADMINER_IMAGE

FROM ${PYIPAD_IMAGE} as caspad
FROM ${PYIPADMINER_IMAGE} as casminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/casklas/caspad

WORKDIR /go/src/github.com/casklas/caspad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=caspad /app/ /app/
COPY --from=casminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/casklas/caspad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
