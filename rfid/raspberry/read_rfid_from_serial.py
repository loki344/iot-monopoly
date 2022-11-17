import serial
import json
import io
import requests


deviceToFieldIndexMap = { "0": [16, 15, 14, 13], "1": [9, 10, 11, 12], "2": [8, 7, 6, 5], "3": [4, 3, 2, 1] }

def map_message_to_field_index(device, field_id):
    field_index = deviceToFieldIndexMap[str(device)][field_id]
    print("mapped message to fieldIndex: " + str(field_index))
    return field_index

deviceCount = 4
devices = []

for i in range(deviceCount):
    device = serial.Serial("/dev/ttyACM"+str(i),9600,timeout=1, )
    device.baudrate=9600
    sio = io.TextIOWrapper(io.BufferedRWPair(device, device))
    devices.append(sio)
    
    

while True: # Run forever

    for i in range(deviceCount):
        line = devices[i].readline()
        if (line.startswith("{")):
            message = json.loads(line)
            print("read message: ")
            print(message)
            field_index = map_message_to_field_index(message["deviceId"], message["fieldId"])
            try:
                requests.patch("http://localhost:3000/players/" + message["playerId"], data={"position": field_index})
            except Exception:
                print("Request failed")
                
