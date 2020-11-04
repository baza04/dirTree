# docker build -t mailgo_hw1 .
FROM golang:1.9.2
COPY . .
RUN go test -v

# FROM golang:latest

# RUN mkdir /app
# ADD . /app
# WORKDIR /app

# RUN go build -o main
# CMD ["/app/main"]