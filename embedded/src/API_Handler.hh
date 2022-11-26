#ifndef UNTITLED3_API_HANDLER_HH
#define UNTITLED3_API_HANDLER_HH
#pragma once
#include <HTTPClient.h>
#include "events.hh"

const String HOST_ADDRESS   =   "192.168.1.185";
const String PORT           =   ":8000";

const String POST_LOGIN     =   "http://" + HOST_ADDRESS + PORT + "/api/login";
const String GET_EVENTS     =   "http://" + HOST_ADDRESS + PORT + "/api/user/events";
const String DELETE_EVENT   =   "http://" + HOST_ADDRESS + PORT + "/api/user/eventDate";

std::string getBearerToken(HTTPClient &http, const std::string& username, const std::string& password){

    nlohmann::json loginPayLoad;
    loginPayLoad["username"] = username;
    loginPayLoad["password"] = password;
    String tokenJSON;

    std::string serializedPayload = loginPayLoad.dump();
    http.begin(POST_LOGIN);
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

    http.begin(GET_EVENTS);
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

bool deleteEvent(HTTPClient &http,const std::string& bearerToken,const int unixTime){

    http.begin(DELETE_EVENT);
    http.addHeader("Content-Type", "application/json");
    http.addHeader("Authorization",("Bearer "+ bearerToken).c_str());
    String rawJSON;

    nlohmann::json deletePayload;
    deletePayload["eventDate"] = std::to_string(unixTime);

    std::string serializedPayload = deletePayload.dump();


    int httpResponseCode = http.sendRequest("DELETE", serializedPayload.c_str());

    if (httpResponseCode>0) {
        Serial.print("HTTP Response code: ");
        Serial.println(httpResponseCode);
        rawJSON = http.getString();
        Serial.println(rawJSON);
        http.end();
        return true;
    }
    else {
        Serial.print("Error code: ");
        Serial.println(httpResponseCode);
        http.end();
        return false;
    }


}



#endif //UNTITLED3_API_HANDLER_HH
