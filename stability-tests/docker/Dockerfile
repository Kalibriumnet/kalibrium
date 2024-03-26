ARG KALIPAD_IMAGE
ARG KALIPAMINER_IMAGE

FROM ${KALIPAD_IMAGE} as Kalibrium
FROM ${KALIPAMINER_IMAGE} as kalibriumminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/kalibriumnet/Kalibrium

WORKDIR /go/src/github.com/kalibriumnet/Kalibrium

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=Kalibrium /app/ /app/
COPY --from=kalibriumminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/Kalibriumnet/Kalibrium/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]