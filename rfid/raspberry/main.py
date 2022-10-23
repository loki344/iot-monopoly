import serial
import json
import requests

ser = serial.Serial('/dev/ttyACM0',9600)
ser2 = serial.Serial('/dev/ttyACM1',9600)
while True:
    read_serial=ser.readline().decode("utf-8")
    obj = json.loads(read_serial)
    
    if(obj is not None):
        requests.patch('http://localhost:3000/players/'+obj['playerId'], json = {'position':obj['fieldId']})
        print (obj)
        print (obj['fieldId'])
        print (obj['playerId'])
    
    read_serial2=ser2.readline().decode("utf-8")
    obj = json.loads(read_serial2)
    
    if(obj is not None):
        requests.patch('http://localhost:3000/players/'+obj['playerId'], json = {'position':obj['fieldId']})
        print (obj)
        print (obj['fieldId'])
        print (obj['playerId'])
