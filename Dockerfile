FROM golang:1.16-alpine as base-image

WORKDIR /workingDir

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o

EXPOSE 8080

CMD ["/dnd-scheduler"]