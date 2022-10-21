FROM golang:1.18

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

EXPOSE 80

CMD ["go", "run", "main.go"]