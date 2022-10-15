# import the necessary packages
from picamera.array import PiRGBArray # Generates a 3D RGB array
from picamera import PiCamera # Provides a Python interface for the RPi Camera Module
import time # Provides time-related functions
import cv2 # OpenCV library
from pyzbar.pyzbar import decode, ZBarSymbol
 
# Initialize the camera
camera = PiCamera()
 
# Set the camera resolution
camera.resolution = (1600, 925)
 
# Set the number of frames per second
camera.framerate = 5
 
# Generates a 3D RGB array and stores it in rawCapture
raw_capture = PiRGBArray(camera, size=(1600, 925))

camera.color_effects = (128,128) # turn camera to black and white
 
# Wait a certain number of seconds to allow the camera time to warmup
time.sleep(1)
det = cv2.QRCodeDetector()
 
# Capture frames continuously from the camera
for frame in camera.capture_continuous(raw_capture, format="rgb", use_video_port="true"):
     
    im = frame.array  # if use cv2
        
    
    # codes = decode(im, symbols=[ZBarSymbol.QRCODE])  # specify code type
    codes = decode(im)  # auto detect code type
    print('Decoded:', codes)

    for code in codes:
        data = code.data.decode('ascii')
        print('Data:', code.data.decode('ascii'))
        print('Code Type:', code.type)
        print('BBox:', code.rect)
        x, y, w, h = code.rect.left, code.rect.top, code.rect.width, code.rect.height
        cv2.rectangle(im, (x,y),(x+w, y+h),(255, 0, 0), 8)
        print('Polygon:', code.polygon)
        cv2.rectangle(im, code.polygon[0], code.polygon[1],(0, 255, 0), 4)

        txt = '(' + code.type + ')  ' + data
        cv2.putText(im, txt, (x - 10, y - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 50, 255), 2)

    text1 = 'No. Codes: %s' % len(codes)
    cv2.putText(im, text1, (5, 15), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 255, 0), 2)

    im = cv2.resize(im, (960,720) , interpolation = cv2.INTER_AREA)
    cv2.imshow('bounding box', im)

    # Wait for keyPress for 1 millisecond
    key = cv2.waitKey(1) & 0xFF
     
    # Clear the stream in preparation for the next frame
    raw_capture.truncate(0)

    # If the `q` key was pressed, break from the loop
    if key == ord("q"):
        break

    

