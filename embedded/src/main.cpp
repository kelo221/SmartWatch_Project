#include <Arduino.h>
#include <iostream>
#include "events.hh"

void setup() {

#pragma region String
    Serial.begin(115200);
    std::string testJson = "[\n"
                       "    {\n"
                       "        \"id\": 64,\n"
                       "        \"eventname\": \"testingNow\",\n"
                       "        \"eventtime\": \"2022-11-18T11:13:00Z\",\n"
                       "        \"created_at\": \"2022-11-17T23:47:07.603253Z\",\n"
                       "        \"userid\": 1,\n"
                       "        \"issilent\": true,\n"
                       "        \"snoozedisabled\": true\n"
                       "    },\n"
                       "    {\n"
                       "        \"id\": 47,\n"
                       "        \"eventname\": \"newUnixMod\",\n"
                       "        \"eventtime\": \"2001-04-13T00:00:00Z\",\n"
                       "        \"created_at\": \"2022-11-17T19:01:33.524045Z\",\n"
                       "        \"userid\": 1,\n"
                       "        \"issilent\": true,\n"
                       "        \"snoozedisabled\": false\n"
                       "    },\n"
                       "    {\n"
                       "        \"id\": 49,\n"
                       "        \"eventname\": \"12121212\",\n"
                       "        \"eventtime\": \"2022-11-17T21:15:02Z\",\n"
                       "        \"created_at\": \"2022-11-17T21:15:08.987922Z\",\n"
                       "        \"userid\": 1,\n"
                       "        \"issilent\": true,\n"
                       "        \"snoozedisabled\": true\n"
                       "    }\n"
                       "]";
#pragma endregion
    eventSpace::events data = nlohmann::json::parse(testJson);

    const std::string test = "testing";

    for(eventSpace::event i : data) {
        Serial.printf("%s\n", i.get_eventname().c_str());
    }
    
}

void loop() {}
