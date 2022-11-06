import serial
import time
ser=serial.Serial('/dev/ttyACM0',9600)
while True:
    readedText = ser.readline()
    print(readedText)
    time.sleep(0.5)
ser.close()
