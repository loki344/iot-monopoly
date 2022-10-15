# import the necessary packages
from picamera.array import PiRGBArray # Generates a 3D RGB array
from picamera import PiCamera # Provides a Python interface for the RPi Camera Module
import time # Provides time-related functions
import cv2 # OpenCV library
import numpy as np
from pyzbar.pyzbar import decode, ZBarSymbol
 
# Initialize the camera
camera = PiCamera()
 
# Set the camera resolution
camera.resolution = (1600, 925)
 
# Set the number of frames per second
camera.framerate = 1
 
# Generates a 3D RGB array and stores it in rawCapture
raw_capture = PiRGBArray(camera, size=(1600, 925))
camera.color_effects = (128,128) # turn camera to black and white
# turn camera to black and white
 
# Wait a certain number of seconds to allow the camera time to warmup
time.sleep(2)

def detect_codes(image, wait=False):
    # codes = decode(im, symbols=[ZBarSymbol.QRCODE])  # specify code type
    codes = decode(image, symbols=[ZBarSymbol.QRCODE])  # auto detect code type

    edge_codes = list(filter(lambda code: code.data.decode('ascii') in ['2','4'], codes))
    
    if len(edge_codes) == 2:
        
        start_point = (0,0)
        end_point = (0,0)
        for edge_code in edge_codes:
            for x in range(len(edge_code.polygon)):
                cv2.putText(image, str(x), edge_code.polygon[x], cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 255, 0), 1)
                
            if edge_code.data.decode('ascii') == '2':
                end_point = edge_code.polygon[3]
            else:
                start_point = edge_code.polygon[1]
        print(start_point, end_point)
        
        cv2.rectangle(image, start_point,end_point,(0, 255, 0), 2)
    
    #TODO only non-edge_codes should be in this for
    for code in codes:
        data = code.data.decode('ascii')
        
        #print('Data:', code.data.decode('ascii'))
        #print('Code Type:', code.type)
        #print('BBox:', code.rect)
        ll, ul, ur, lr = code.polygon[0], code.polygon[1], code.polygon[2], code.polygon[3]
        pts = np.array([ll,ul,ur,lr], np.int32)
        pts = pts.reshape((-1,1,2))
        cv2.polylines(image, [pts], True,(255, 0, 0))
        #print('Polygon:', code.polygon)

        x, y= code.rect.left, code.rect.top
        txt = '(' + code.type + ')  ' + data + ' ' + code.orientation
        cv2.putText(image, txt, (x - 10, y - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 50, 255), 2)


    text1 = 'No. Codes: %s' % len(codes)
    cv2.putText(image, text1, (5, 15), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 255, 0), 2)

    #masked = cv2.resize(masked, (960,720) , interpolation = cv2.INTER_AREA)
    cv2.imshow('bounding box', image)
    
    if wait:
        # waiting using waitKey method
        cv2.waitKey(0)
    

detect_codes(cv2.imread('Picture1.png'), True)


# Capture frames continuously from the camera
for frame in camera.capture_continuous(raw_capture, format="bgr", use_video_port="true"):
     
    image = frame.array  # if use cv2
    
    #blurred_image = cv2.GaussianBlur(gray_image, (7, 7), 0)

    #ret, threshInv = cv2.threshold(gray_image, 0, 255, cv2.THRESH_BINARY_INV | cv2.THRESH_OTSU)
    #masked = cv2.bitwise_and(image, image, mask=threshInv)
    masked = image
    
    detect_codes(image)

    # Wait for keyPress for 1 millisecond
    key = cv2.waitKey(1) & 0xFF
         # If the `q` key was pressed, break from the loop
    if key == ord("q"):
        break
    
    # Clear the stream in preparation for the next frame
    raw_capture.truncate(0)



