FROM  --platform=linux/amd64 golang:latest

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN export GO111MODULE=on
RUN go mod download
COPY . .
RUN go build -o /bass-neck-map

EXPOSE 443

CMD ["/bass-neck-map"]