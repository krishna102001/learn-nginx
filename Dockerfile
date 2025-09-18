FROM golang:1.25-alpine

WORKDIR /app/

COPY go.* .

RUN go mod tidy

COPY . /app/

RUN go build -o app .

EXPOSE 3000

CMD [ "./app" ]