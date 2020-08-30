#!/bin/bash

go get -u github.com/gin-gonic/gin
go get github.com/eclipse/paho.mqtt.golang
go run /go/src/mqtt/pub_server.go
