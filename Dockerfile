# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /sorter

# Download go modules
COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main
RUN chmod +x /main

CMD [ "/main" ]