FROM golang:alpine

ENV APP_HOME /go/src/web-server
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"
COPY . .

RUN go build -o web-server

EXPOSE 8000

CMD [ "./web-server" ]
