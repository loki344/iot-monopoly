FROM golang:1.9.2 
ADD . /go/src/iot-monopoly
WORKDIR /go/src/iot-monopoly
RUN go get iot-monopoly
RUN go install
EXPOSE 3000
ENTRYPOINT ["/go/bin/iot-monopoly"]
