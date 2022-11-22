#ifndef UNTITLED3_API_HANDLER_HH
#define UNTITLED3_API_HANDLER_HH
#pragma once
#include <HTTPClient.h>
#include "events.hh"

const char* POST_LoginPath = "http://192.168.1.185:8000/api/login";
const char* GET_EventsPath = "http://192.168.1.185:8000/api/user/events";

std::string getBearerToken(HTTPClient &http, const std::string& username, const std::string& password){

    nlohmann::json loginPayLoad;
    loginPayLoad["username"] = username;
    loginPayLoad["password"] = password;
    String tokenJSON;

    std::string serializedPayload = loginPayLoad.dump();
    Serial.println(serializedPayload.c_str());
    http.begin(POST_LoginPath);
    http.addHeader("Content-Type", "application/json");
    int httpResponseCode = http.POST((serializedPayload +"\n").c_str());

    if (httpResponseCode>0) {
        Serial.print("HTTP Response code: ");
        Serial.println(httpResponseCode);
        tokenJSON = http.getString();
        Serial.println(tokenJSON);
    }
    else {
        Serial.print("Error code: ");
        Serial.println(httpResponseCode);
    }
    http.end();

    nlohmann::json bearerJSON = nlohmann::json::parse(tokenJSON);
    std::string bearerRaw = bearerJSON["token"].get<std::string>();
    return bearerRaw;
}

eventSpace::events getEvents(HTTPClient &http,const std::string& bearerToken){

    http.begin(GET_EventsPath);
    http.addHeader("Authorization",("Bearer "+ bearerToken).c_str());
    String rawJSON;

    int httpResponseCode = http.GET();

    if (httpResponseCode>0) {
        Serial.print("HTTP Response code: ");
        Serial.println(httpResponseCode);
        rawJSON = http.getString();
        Serial.println(rawJSON);
    }
    else {
        Serial.print("Error code: ");
        Serial.println(httpResponseCode);
    }
    http.end();

    return nlohmann::json::parse(rawJSON);
}


#endif //UNTITLED3_API_HANDLER_HH
