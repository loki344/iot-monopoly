cd ./frontend
docker build -t loki334/iot-monopoly:frontend-latest .
cd ../backend
docker build -t loki334/iot-monopoly:backend-latest .
cd ../

docker push loki334/iot-monopoly:frontend-latest
docker push loki334/iot-monopoly:backend-latest