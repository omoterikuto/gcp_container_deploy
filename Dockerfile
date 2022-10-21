FROM golang:1.18

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0

WORKDIR ${ROOT}
COPY ./ ./

EXPOSE 8080

CMD ["go", "run", "main.go"]