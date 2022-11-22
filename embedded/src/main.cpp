#include <Arduino.h>
#include <iostream>
#include "API_Handler.hh"
#include <WiFi.h>
#include <HTTPClient.h>


const char* wifiSSID = "ESP_POINT";
const char* wifiPassword = "";

//  eventSpace::events data = nlohmann::json::parse(testJson);

void setup() {
    Serial.begin(115200);
    WiFi.begin(wifiSSID, wifiPassword);

    for(int loopCounter = 0; WiFiClass::status() != WL_CONNECTED; ++loopCounter) {
        Serial.println("Connecting...");
        sleep(1);

        if (loopCounter==5){
            Serial.println("Failed to connect.");
            exit(69);
        }
    }

    auto http = new HTTPClient();
    auto token = getBearerToken(*http,"test","test");
    auto events = getEvents(*http, token);

    for(eventSpace::event i : events) {
        Serial.printf("%s\n", i.get_eventname().c_str());
    }

}

void loop() {
    sleep(1);
}
