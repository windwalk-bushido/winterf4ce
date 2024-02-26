FROM golang:1.22-alpine
WORKDIR /code

COPY . /code
RUN go mod download
RUN go build -o main .

EXPOSE 55555
CMD ["nohup", "./main", "&"]
