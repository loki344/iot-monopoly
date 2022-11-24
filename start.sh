#!/bin/bash
PORT=8080 node frontend/build/index.js &
cd ./backend & go run iot-monopoly
