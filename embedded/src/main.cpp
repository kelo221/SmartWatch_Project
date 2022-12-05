#include <Arduino.h>
#include <iostream>
#include "API_Handler.hh"
#include "beeper.hh"
#include <WiFi.h>
#include <HTTPClient.h>


///TODO CLEAR PASSWORD BEFORE PUSHING
const char* wifiSSID = "ESP_POINT";
const char* wifiPassword = "";
///TODO CLEAR PASSWORD BEFORE PUSHING

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

    eventSpace::event test;

    test.set_eventname("hello");
    test.set_eventtime(123);
    test.set_issilent(false);
    test.set_snoozedisabled(true);


    Serial.println(changeEventTime(*http,token,1000,test) ? "OK!!" : "FAILED!!");


/*    for(const eventSpace::event& i : events) {
        Serial.printf("%s\n", i.get_eventname().c_str());
    }*/

}

void loop() {
    sleep(1);
}
