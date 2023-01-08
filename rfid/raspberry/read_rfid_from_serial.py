import serial
import json
import io
import requests


deviceToFieldIndexMap = { "0": [16, 15, 14, 13], "1": [9, 10, 11, 12], "2": [8, 7, 6, 5], "3": [4, 3, 2, 1] }
tagIdToPlayerIdMap = {"5A-B8-D4-5C": "Player_1", "BA-84-DE-5C": "Player_2", "2A-3A-B6-5C": "Player_3", "AA-19-F6-5B": "Player_4"}

def map_message_to_field_index(device, field_id):
    field_index = deviceToFieldIndexMap[str(device)][field_id]
    print("mapped message to fieldIndex: " + str(field_index))
    return field_index

def map_tag_id_to_player_id(tag_id):
    player_id = tagIdToPlayerIdMap[tag_id]
    print("mapped tagId: "+ tag_id +" to playerId: " + str(player_id))
    return player_id

deviceCount = 4

devices = []


for i in range(deviceCount):
    device = serial.Serial("/dev/ttyACM"+str(i),9600,timeout=0.5, )
    device.baudrate=9600
    textWrapper = io.TextIOWrapper(io.BufferedRWPair(device, device))
    devices.append(textWrapper)


while True: # Run forever

    for i in range(deviceCount):
        line = devices[i].readline()
        if line.startswith("{"):
            try:
                message = json.loads(line)
                print("read message: ")
                print(message)

                field_index = map_message_to_field_index(message["deviceId"], message["fieldId"])
                player_id = map_tag_id_to_player_id(message["playerId"])

                requests.patch("http://localhost:3000/games/current/players/" + player_id, data={"position": field_index})
            except Exception:
                print("Request failed")
                
