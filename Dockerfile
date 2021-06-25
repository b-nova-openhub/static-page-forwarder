FROM golang:1.16.5-alpine

WORKDIR $GOPATH/src/github.com/b-nova-openhub/stapafor

COPY . .

RUN go get -d -v ./... \
    && go build -o bin/stapafor cmd/stapafor/main.go \
    && go install -v ./... \
    && chmod +x stapafor.sh

EXPOSE 8080

CMD ["sh", "stapafor.sh"]