FROM golang:latest

WORKDIR WORKDIR /go/src/github.com/MicBun/62teknologi-senior-backend-test-Michael_Buntarman

COPY . .

RUN go get -d -v ./...

RUN go build -o 62teknologi-senior-backend-test-Michael_Buntarman ./bin/web/main.go

EXPOSE 8080

ENTRYPOINT ["./62teknologi-senior-backend-test-Michael_Buntarman"]