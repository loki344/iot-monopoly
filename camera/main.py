# import the necessary packages
import cv2  # OpenCV library
import numpy as np
import time  # Provides time-related functions
from picamera import PiCamera  # Provides a Python interface for the RPi Camera Module
from picamera.array import PiRGBArray  # Generates a 3D RGB array
from pyzbar.pyzbar import decode, ZBarSymbol
import requests

# Initialize the camera
camera = PiCamera()

# Set the camera resolution
camera.resolution = (1600, 925)

# Set the number of frames per second
camera.framerate = 1

# Generates a 3D RGB array and stores it in rawCapture
raw_capture = PiRGBArray(camera, size=(1600, 925))
camera.color_effects = (128, 128)  # turn camera to black and white
# turn camera to black and white

# Wait a certain number of seconds to allow the camera time to warmup
time.sleep(2)

is_edge_code = lambda code: code.data.decode('ascii') in ['2', '4']


#TODO what do we want
# once we detected the edges, we have to remember the rectangle coordinates in order to crop the image
# after that, we only need to relocate the edges once at a time
# For the other QR codes, we have to map the coordinates to field numbers
# once we detect a QR code in the same field for a given amount of time, we have to send a request to the backend

def draw_outer_border(edge_codes, image):
    start_point = (0, 0)
    end_point = (0, 0)
    for edge_code in edge_codes:
        if edge_code.data.decode('ascii') == '2':
            end_point = edge_code.polygon[3]
        else:
            start_point = edge_code.polygon[1]
            
    cv2.rectangle(image, start_point, end_point, (0, 255, 0), 2)
    return image[end_point[1]:start_point[1],start_point[0]:end_point[0]]
    


def mark_codes(non_edge_codes, image):
    for code in non_edge_codes:
        data = code.data.decode('ascii')

        ll, ul, ur, lr = code.polygon[0], code.polygon[1], code.polygon[2], code.polygon[3]
        pts = np.array([ll, ul, ur, lr], np.int32)
        pts = pts.reshape((-1, 1, 2))
        cv2.polylines(image, [pts], True, (255, 0, 0))

        x, y = code.rect.left, code.rect.top
        txt = '(' + code.type + ')  ' + data + ' ' + code.orientation
        cv2.putText(image, txt, (x - 10, y - 10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 50, 255), 2)
        
        

def detect_codes(image, wait=False):
    # codes = decode(im, symbols=[ZBarSymbol.QRCODE])  # specify code type
    codes = decode(image, symbols=[ZBarSymbol.QRCODE])  # auto detect code type

    edge_codes = list(filter(is_edge_code, codes))

    if len(edge_codes) == 2:
        image = draw_outer_border(edge_codes, image)
    
    non_edge_codes = list(filter(is_edge_code, codes))
    mark_codes(non_edge_codes, image)

    text1 = 'No. Codes: %s' % len(codes)
    cv2.putText(image, text1, (5, 15), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 255, 0), 2)

    # masked = cv2.resize(masked, (960,720) , interpolation = cv2.INTER_AREA)
    cv2.imshow('bounding box', image)

    if wait:
        # waiting using waitKey method
        cv2.waitKey(0)


# get id from qr code
requests.patch('http://localhost:3000/players/21898eb5-e233-47d1-bc40-17a721bbd148', json = {'position':5})


detect_codes(cv2.imread('Picture1.png'), True)

# Capture frames continuously from the camera
for frame in camera.capture_continuous(raw_capture, format="bgr", use_video_port="true"):

    image = frame.array  # if use cv2

    # blurred_image = cv2.GaussianBlur(gray_image, (7, 7), 0)

    # ret, threshInv = cv2.threshold(gray_image, 0, 255, cv2.THRESH_BINARY_INV | cv2.THRESH_OTSU)
    # masked = cv2.bitwise_and(image, image, mask=threshInv)
    masked = image

    detect_codes(image)

    # Wait for keyPress for 1 millisecond
    key = cv2.waitKey(1) & 0xFF
    # If the `q` key was pressed, break from the loop
    if key == ord("q"):
        break

    # Clear the stream in preparation for the next frame
    raw_capture.truncate(0)
