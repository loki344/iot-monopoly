#!/usr/bin/env python
# -*- coding: utf8 -*-
#
#    Copyright 2014,2018 Mario Gomez <mario.gomez@teubi.co>
#
#    This file is part of MFRC522-Python
#    MFRC522-Python is a simple Python implementation for
#    the MFRC522 NFC Card Reader for the Raspberry Pi.
#
#    MFRC522-Python is free software: you can redistribute it and/or modify
#    it under the terms of the GNU Lesser General Public License as published by
#    the Free Software Foundation, either version 3 of the License, or
#    (at your option) any later version.
#
#    MFRC522-Python is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU Lesser General Public License for more details.
#
#    You should have received a copy of the GNU Lesser General Public License
#    along with MFRC522-Python.  If not, see <http://www.gnu.org/licenses/>.
#

import RPi.GPIO as GPIO
import mfrc522 as MFRC522
import signal
import time
import requests

continue_reading = True

tagIdToAccountIdMap = {"51-168-138-16": "Player_1", "19-4-182-26": "Player_2", "67-241-231-14": "Player_3", "163-217-53-15": "Player_4"}

def map_tag_id_to_account_id(tag_id):
    account_id = tagIdToAccountIdMap[tag_id]
    print("mapped tagId: "+ tag_id +" to accountId: " + str(account_id))
    return account_id


# Capture SIGINT for cleanup when the script is aborted
def end_read(signal,frame):
    global continue_reading
    print ("Ctrl+C captured, ending read.")
    continue_reading = False
    GPIO.cleanup()

# Hook the SIGINT
signal.signal(signal.SIGINT, end_read)

# Create an object of the class MFRC522
MIFAREReader = MFRC522.MFRC522()

# This loop keeps checking for chips. If one is near it will get the UID and authenticate
while continue_reading:

    # Scan for cards    
    (status,TagType) = MIFAREReader.MFRC522_Request(MIFAREReader.PICC_REQIDL)

    # If a card is found
    if status == MIFAREReader.MI_OK:

        # Get the UID of the card
        (status, uid) = MIFAREReader.MFRC522_Anticoll()

    # If we have the UID, continue
    if status == MIFAREReader.MI_OK:

        LED = 18
        GPIO.setup(LED, GPIO.OUT)
        GPIO.output(LED, GPIO.HIGH)
        time.sleep(3)
        GPIO.output(LED, GPIO.LOW)
        tagId = "{}-{}-{}-{}".format(uid[0], uid[1], uid[2], uid[3])
        map_tag_id_to_account_id(tagId)
        try:
            requests.patch("http://localhost:3000/transactions/latest", data={"accepted": True, "senderId": map_tag_id_to_account_id(tagId)})
        except Exception:
            print("Request failed")
    
