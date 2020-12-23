FROM golang:1.15

WORKDIR /wwgt-mm-server
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .
EXPOSE 8080

COPY dev-runscript.sh ./
CMD [ "sh", "./dev-runscript.sh" ]
