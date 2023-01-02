
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

MFRC522 mfrc522[NR_OF_READERS];

void setup() {

  Serial.begin(9600);          
  while (!Serial);              

  SPI.begin();

  for(uint8_t i = 0; i < NR_OF_READERS; i++) {
    pinMode(ledPins[i], OUTPUT);
    digitalWrite(ledPins[i], LOW);
  }

  for (uint8_t reader = 0; reader < NR_OF_READERS; reader++) {
    mfrc522[reader].PCD_Init(ssPins[reader], RST_PIN);
    Serial.print(F("Reader "));
    Serial.print(reader);
    Serial.print(F(": "));
    mfrc522[reader].PCD_DumpVersionToSerial();
    delay(300);
  }
}


void loop() {

  for (uint8_t reader = 0; reader < NR_OF_READERS; reader++) {

    delay(100);
    if (mfrc522[reader].PICC_IsNewCardPresent() && mfrc522[reader].PICC_ReadCardSerial()) {

        delay(1500);
        mfrc522[reader].PICC_ReadCardSerial();
        mfrc522[reader].PICC_IsNewCardPresent();

        if (mfrc522[reader].PICC_ReadCardSerial()) {

            digitalWrite(ledPins[reader], HIGH);
            printHex(mfrc522[reader].uid.uidByte, mfrc522[reader].uid.size, reader);

            mfrc522[reader].PICC_HaltA();
            mfrc522[reader].PCD_StopCrypto1();

            flashLED(reader);
        }
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

void flashLED(uint8_t pin) {
    delay(500);
    digitalWrite(ledPins[pin], LOW);
    delay(100);
    digitalWrite(ledPins[pin], HIGH);
    delay(100);
    digitalWrite(ledPins[pin], LOW);
    delay(100);
    digitalWrite(ledPins[pin], HIGH);
    delay(100);
    digitalWrite(ledPins[pin], LOW);
}
