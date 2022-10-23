import pyMultiSerial as p
import json
import requests


# Create object of class pyMultiSerial
ms = p.MultiSerial()

ms.baudrate = 9600
ms.timeout = 2

def port_connection_found_callback(portno, serial):
    print ("Port Found: "+portno)


#register callback function
ms.port_connection_found_callback = port_connection_found_callback


# Callback on receiving port data
# Parameters: Port Number, Serial Port Object, Text read from port
def port_read_callback(portno, serial, text):
    #print ("Received '"+text+"' from port "+portno)
    obj = json.loads(text)
    print(obj)
    requests.patch('http://localhost:3000/players/'+obj['playerId'], {'position':obj['fieldId']})
    pass


#register callback function
ms.port_read_callback = port_read_callback


# Callback on port disconnection. Triggered when a device is disconnected from port.
# Parameters: Port No
def port_disconnection_callback(portno):
    print("Port "+portno+" disconnected")


#register callback function
ms.port_disconnection_callback = port_disconnection_callback


# Start Monitoring ports
ms.Start()


## To stop monitoring, press Ctrl+C in the console or command line.


# Caution: Any code written below ms.Start() will be executed only after monitoring is stopped.
# Make use of callback functions to execute your code. 
