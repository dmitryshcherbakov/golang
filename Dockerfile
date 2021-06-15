FROM golang:1.14.2
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gorilla/websocket
RUN go get github.com/githubnemo/CompileDaemon
#RUN go mod download
#RUN go build -o main .
#RUN go install .
#CMD ["/app/main"]
ENTRYPOINT CompileDaemon --build="go build -o main ." --command=./main