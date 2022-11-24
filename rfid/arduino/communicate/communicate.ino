
/**
   --------------------------------------------------------------------------------------------------------------------
   Example sketch/program showing how to read data from more than one PICC to serial.
   --------------------------------------------------------------------------------------------------------------------
   This is a MFRC522 library example; for further details and other examples see: https://github.com/miguelbalboa/rfid

   Example sketch/program showing how to read data from more than one PICC (that is: a RFID Tag or Card) using a
   MFRC522 based RFID Reader on the Arduino SPI interface.

   Warning: This may not work! Multiple devices at one SPI are difficult and cause many trouble!! Engineering skill
            and knowledge are required!

   @license Released into the public domain.

   Typical pin layout used:
   -----------------------------------------------------------------------------------------
               MFRC522      Arduino       Arduino   Arduino    Arduino          Arduino
               Reader/PCD   Uno/101       Mega      Nano v3    Leonardo/Micro   Pro Micro
   Signal      Pin          Pin           Pin       Pin        Pin              Pin
   -----------------------------------------------------------------------------------------
   RST/Reset   RST          9             5         D9         RESET/ICSP-5     RST
   SPI SS 1    SDA(SS)      ** custom, take a unused pin, only HIGH/LOW required *
   SPI SS 2    SDA(SS)      ** custom, take a unused pin, only HIGH/LOW required *
   SPI MOSI    MOSI         11 / ICSP-4   51        D11        ICSP-4           16
   SPI MISO    MISO         12 / ICSP-1   50        D12        ICSP-1           14
   SPI SCK     SCK          13 / ICSP-3   52        D13        ICSP-3           15

*/

#include <SPI.h>
#include <MFRC522.h>

// PIN Numbers : RESET + SDAs
#define RST_PIN         8
#define SS_1_PIN        2
#define SS_2_PIN        3
#define SS_3_PIN        4
#define SS_4_PIN        5

// Led and Relay PINS
#define LED_1        6
#define LED_2        7
#define LED_3        9
#define LED_4        10


#define NR_OF_READERS   4

byte ssPins[] = {SS_1_PIN, SS_2_PIN, SS_3_PIN, SS_4_PIN};
byte ledPins[] = {LED_1, LED_2, LED_3, LED_4};
uint8_t tagTracking[] = {0, 0, 0, 0};

// Create an MFRC522 instance :
MFRC522 mfrc522[NR_OF_READERS];

/**
   Initialize.
*/
void setup() {

  Serial.begin(9600);           // Initialize serial communications with the PC
  while (!Serial);              // Do nothing if no serial port is opened (added for Arduinos based on ATMEGA32U4)

  SPI.begin();                  // Init SPI bus

  for(uint8_t i = 0; i < NR_OF_READERS; i++) {
    pinMode(ledPins[i], OUTPUT);
    digitalWrite(ledPins[i], LOW);
  }


  /* looking for MFRC522 readers */
  for (uint8_t reader = 0; reader < NR_OF_READERS; reader++) {
    mfrc522[reader].PCD_Init(ssPins[reader], RST_PIN);
    Serial.print(F("Reader "));
    Serial.print(reader);
    Serial.print(F(": "));
    mfrc522[reader].PCD_DumpVersionToSerial();
    //mfrc522[reader].PCD_SetAntennaGain(mfrc522[reader].RxGain_max);
    delay(1000);
  }
}

/*
   Main loop.
*/

void loop() {

  for (uint8_t reader = 0; reader < NR_OF_READERS; reader++) {

    delay(100);
    // Looking for new cards
    if (mfrc522[reader].PICC_IsNewCardPresent() && mfrc522[reader].PICC_ReadCardSerial()) {
      
        digitalWrite(ledPins[reader], HIGH);

        printHex(mfrc522[reader].uid.uidByte, mfrc522[reader].uid.size, reader);

        // Halt PICC
        mfrc522[reader].PICC_HaltA();
        // Stop encryption on PCD
        mfrc522[reader].PCD_StopCrypto1();
  
        delay(500);
        digitalWrite(ledPins[reader], LOW);
        delay(100);
        digitalWrite(ledPins[reader], HIGH);
        delay(100);
        digitalWrite(ledPins[reader], LOW);
        delay(100);
        digitalWrite(ledPins[reader], HIGH);
        delay(100);
        digitalWrite(ledPins[reader], LOW);
    
    }
  }
}
void printHex(byte *buffer, byte bufferSize, uint8_t reader) {

  Serial.print("{\"deviceId\":2,\"fieldId\":");
  Serial.print(reader);
  Serial.print(", \"playerId\":\"");
  Serial.print(buffer[0] < 0x10 ? "0" : "");
  Serial.print(buffer[0], HEX);

  for (byte i = 1; i < bufferSize; i++) {
    Serial.print(buffer[i] < 0x10 ? "0" : "-");
    Serial.print(buffer[i], HEX);
  }
  Serial.print("\"}");
}
