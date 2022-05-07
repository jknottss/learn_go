FROM golang:latest
WORKDIR /crudTest
COPY . .
RUN go mod download
RUN go build -o crud .
CMD ./crud