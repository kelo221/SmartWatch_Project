#pragma once
#include <HTTPClient.h>
#include "events.hh"

const std::string HOST_ADDRESS   =   "192.168.1.185";
const std::string PORT           =   ":8000";

const std::string POST_LOGIN     =   "http://" + HOST_ADDRESS + PORT + "/api/login";
const std::string GET_EVENTS     =   "http://" + HOST_ADDRESS + PORT + "/api/user/events";
const std::string DELETE_EVENT   =   "http://" + HOST_ADDRESS + PORT + "/api/user/eventDate";
const std::string PATCH_EVENT   =   "http://" + HOST_ADDRESS + PORT + "/api/user/event";

std::string getBearerToken(HTTPClient &http, const std::string& username, const std::string& password){

    nlohmann::json loginPayLoad;
    loginPayLoad["username"] = username;
    loginPayLoad["password"] = password;
    std::string tokenJSON;

    std::string serializedPayload = loginPayLoad.dump();
    http.begin(POST_LOGIN.c_str());
    http.addHeader("Content-Type", "application/json");
    int httpResponseCode = http.POST((serializedPayload +"\n").c_str());

    if (httpResponseCode>0) {
        Serial.print("HTTP Response code: ");
        Serial.println(httpResponseCode);
        tokenJSON = http.getString().c_str();
        Serial.println(tokenJSON.c_str());
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

    http.begin(GET_EVENTS.c_str());
    http.addHeader("Authorization",("Bearer "+ bearerToken).c_str());
    std::string rawJSON;

    int httpResponseCode = http.GET();

    if (httpResponseCode>0) {
        Serial.print("HTTP Response code: ");
        Serial.println(httpResponseCode);
        rawJSON = http.getString().c_str();
        Serial.println(rawJSON.c_str());
    }
    else {
        Serial.print("Error code: ");
        Serial.println(httpResponseCode);
    }
    http.end();

    return nlohmann::json::parse(rawJSON);
}

bool deleteEvent(HTTPClient &http,const std::string& bearerToken,const int unixTime){

    http.begin(DELETE_EVENT.c_str());
    http.addHeader("Content-Type", "application/json");
    http.addHeader("Authorization",("Bearer "+ bearerToken).c_str());
    std::string rawJSON;

    nlohmann::json deletePayload;
    deletePayload["eventDate"] = std::to_string(unixTime);

    std::string serializedPayload = deletePayload.dump();

    bool httpSuccess;

    int httpResponseCode = http.sendRequest("DELETE", serializedPayload.c_str());

    if (httpResponseCode==200) {
        Serial.print("HTTP Response code: ");
        Serial.println(httpResponseCode);
        rawJSON = http.getString().c_str();
        Serial.println(rawJSON.c_str());
        http.end();
        httpSuccess = true;
    }
    else {
        Serial.print("Error code: ");
        Serial.println(httpResponseCode);
        http.end();
        httpSuccess = false;
    }

    return httpSuccess;

}


bool changeEventTime(HTTPClient &http, const std::string& bearerToken, const int newUnixTime, const eventSpace::event& newEvent){

    http.begin(PATCH_EVENT.c_str());
    http.addHeader("Content-Type", "application/json");
    http.addHeader("Authorization",("Bearer "+ bearerToken).c_str());
    std::string rawJSON;

    nlohmann::json patchPayload;
    patchPayload["eventTimeNewUnix"] = std::to_string(newUnixTime);
    patchPayload["eventTimeOldUnix"] = std::to_string(newEvent.get_eventtime());
    patchPayload["eventName"] = newEvent.get_eventname();
    patchPayload["isSilent"] = std::to_string(newEvent.get_issilent());
    patchPayload["SnoozeDisabled"] = std::to_string(newEvent.get_snoozedisabled());

    std::string serializedPayload = patchPayload.dump();

    bool httpSuccess;

    int httpResponseCode = http.sendRequest("PATCH", serializedPayload.c_str());

    if (httpResponseCode==200) {
        Serial.print("HTTP Response code: ");
        Serial.println(httpResponseCode);
        rawJSON = http.getString().c_str();
        Serial.println(rawJSON.c_str());
        http.end();
        httpSuccess = true;
    }
    else {
        Serial.print("Error code: ");
        Serial.println(httpResponseCode);
        http.end();
        httpSuccess = false;
    }

    return httpSuccess;

}