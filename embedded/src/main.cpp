#include <Arduino.h>
#include <iostream>
#include "API_Handler.hh"
#include "beeper.hh"
#include <WiFi.h>
#include <HTTPClient.h>
#include <WiFi.h>
#include "EEPROM.h"
#include <vector>
#include <LiquidCrystal.h>
#define CLK_FREQ 80000000
#define PRESCALER 80
#define TICKS_PER_SECOND (CLK_FREQ / PRESCALER)
#define SSID_NAME "Simon’s iPhone"
#define SSID_PASSWORD "wizard12"
#define NUMAPPS 3
const char *wifiSSID = SSID_NAME;
const char *wifiPassword = SSID_PASSWORD;

//  eventSpace::events data = nlohmann::json::parse(testJson);

/* TODO:
    Add GPIO interrupts for all three buttons

*/

// LCD
static LiquidCrystal *lcd;
// Core variables
static const std::vector<std::string> menuItmes = {"Metronome", "Tuner", "Whatever"};
static bool button1Pressed = false;
static bool button2Pressed = false;
static bool button3Pressed = false;
// FreeRTOS parameters
static uint8_t ucParameterToPass;
static TaskHandle_t Task1;
static TaskHandle_t Task2;
static TaskHandle_t TaskAlarm;
static TaskHandle_t TaskMetronome;
static TaskHandle_t TaskPitch;
static TaskHandle_t *CurrentTask = nullptr; // The main task is never the current task
static SemaphoreHandle_t alarmSemaphore;
static SemaphoreHandle_t lcdSemaphore;
// Time parameters
const char *ntpServer = "pool.ntp.org";
const long gmtOffset_sec = 7200;
const int daylightOffset_sec = 3600;
static struct tm timeinfo;
typedef struct Data
{
    time_t time;
} Data;
static Data data;

static std::vector<time_t> alarms;
static eventSpace::events events;
// Timer
hw_timer_t *timer = NULL;

#pragma region Function definitons
void vTaskDebug(void *pvParameters);
void vTaskMain(void *pvParameters);
void vTaskMetronome(void *pvParameters);
void printToLCDWithMutex(const std::string &str, uint8_t col, uint8_t line);
void clearLineWithMutex(uint8_t line);
#pragma endregion

#pragma region Functions
void printLocalTime()
{
    char strtime_buf[64];
    setenv("FI", "GMT-2", 1);
    tzset();

    // time_t now;
    time(&data.time);
    localtime_r(&data.time, &timeinfo);
    strftime(strtime_buf, sizeof(strtime_buf), "  %a %H:%M:%S  ", &timeinfo);

    std::string s(strtime_buf);
    printToLCDWithMutex(s, 0, 0);
}

void setTMStructToTime(const time_t *time)
{
    gmtime_r(time, &timeinfo);
}

template <typename T>
void serialize(const T &t)
{
    EEPROM.put(0, t);
}
template <typename T>
void deserialize(const T &t)
{
    EEPROM.get(0, data);
}

#pragma region ISRs
void IRAM_ATTR onTimer()
{
}

void IRAM_ATTR onButton1()
{
    button1Pressed = true;
}
void IRAM_ATTR onButton2()
{
    button2Pressed = true;
}
void IRAM_ATTR onButton3()
{
    button3Pressed = true;
}
#pragma endregion

bool checkAlarms(const time_t currentTime)
{
    if (!alarms.empty())
    {
        auto it = alarms.begin();
        while (it != alarms.end())
        {
            if (*it > currentTime - 2 && *it < currentTime + 2)
            {
                it = alarms.erase(it);
                return true;
            }
            ++it;
        }
    }
    return false;
}

void printToLCDWithMutex(const std::string &str, uint8_t col, uint8_t line)
{
    if (xSemaphoreTake(lcdSemaphore, portMAX_DELAY) == pdTRUE)
    {
        lcd->setCursor(col, line);
        lcd->print(str.c_str());
        xSemaphoreGive(lcdSemaphore);
    }
}

void clearLineWithMutex(uint8_t line)
{
    if (xSemaphoreTake(lcdSemaphore, portMAX_DELAY) == pdTRUE)
    {
        lcd->clearLine(1);
        xSemaphoreGive(lcdSemaphore);
    }
}

#pragma endregion

#pragma region Tasks

void vTaskMain(void *pvParameters)
{
    Serial.println(xPortGetCoreID());
    static int menuIndex = 0;

    int i = 0;
    for (;;)
    {
        printLocalTime();

#pragma region Buttons
        if (button1Pressed && button3Pressed)
        {
            button1Pressed = false;
            button3Pressed = false;
            if (CurrentTask != nullptr)
            {
                vTaskSuspend(*CurrentTask);
                CurrentTask = nullptr;
            }
        }
        // Only select and execute tasks when no task is running
        if (CurrentTask == nullptr)
        {
            if (button1Pressed)
            {
                button1Pressed = false;
                menuIndex--;
                if (menuIndex < 0)
                {
                    menuIndex == NUMAPPS - 1;
                }
                clearLineWithMutex(1);
                printToLCDWithMutex(menuItmes[menuIndex].c_str(), 0, 1);
            }

            if (button3Pressed)
            {
                button3Pressed;
                menuIndex++;
                if (menuIndex > NUMAPPS - 1)
                {
                    menuIndex = 0;
                }
                clearLineWithMutex(1);
                printToLCDWithMutex(menuItmes[menuIndex].c_str(), 0, 1);
            }

            if (button2Pressed)
            {
                button2Pressed = false;
                switch (menuIndex)
                {
                case 0: // Metronome
                    CurrentTask = &TaskMetronome;
                    vTaskResume(TaskMetronome);
                    break;
                case 1: // Pitch
                    CurrentTask = &TaskPitch;
                    vTaskResume(TaskPitch);
                    break;
                case 2: // Other task...
                default:
                    break;
                }
                menuIndex = 0;
            }
        }
#pragma endregion

        if (checkAlarms(data.time))
        {
            if (xSemaphoreTake(alarmSemaphore, portMAX_DELAY) == pdTRUE)
            {
                // vTaskSuspend(*CurrentTask);
                // vTaskResume(TaskAlarm);
                // CurrentTask = &TaskAlarm;
            }
        }
        vTaskDelay(1000 / portTICK_PERIOD_MS);
    }
}
// handles the alarm
void vTaskAlarm(void *pvParameters)
{
    // start beeping
    bool exit = false;
    int i = 0;
    while (!exit)
    {
        // beep for x miliseconds
        // stop beeping
        // check input
        // exit
        Serial.println("Beep");
        i++;
        if (i == 10)
        {
            return;
        }
    }
    xSemaphoreGive(alarmSemaphore);
    CurrentTask = nullptr;
    vTaskSuspend(NULL); // suspend oneself
}

void vTaskDebug(void *pvParameters)
{
    Serial.println(xPortGetCoreID());
    for (;;)
    {
        Serial.println("Task 2");
        delay(1000);
    }
}

void vTaskMetronome(void *pvParameters)
{
    // set up metronome
    int bpm = 60;
    std::string s = "BPM: ";
    clearLineWithMutex(1);
    printToLCDWithMutex(s + std::to_string(bpm), 0, 1);
    while (!button2Pressed)
    {
        if (button1Pressed)
        {
            button1Pressed = false;
            bpm--;
            clearLineWithMutex(1);
            printToLCDWithMutex(s + std::to_string(bpm), 0, 1);
        }
        if (button3Pressed)
        {
            button3Pressed = false;
            bpm++;
            clearLineWithMutex(1);
            printToLCDWithMutex(s + std::to_string(bpm), 0, 1);
        }
        vTaskDelay(100 / portTICK_PERIOD_MS);
    }
    for (;;)
    {
        for (int i = 0; i < 4; i++)
        {
            vTaskDelay(((6000 / bpm) / 4) / portTICK_PERIOD_MS);

            if (xSemaphoreTake(lcdSemaphore, portMAX_DELAY) == pdTRUE)
            {
                lcd->clearLine(1);
                xSemaphoreGive(lcdSemaphore);
            }
            printToLCDWithMutex("----", i, 1);
            if (i == 3)
            {
                Serial.println("Metronome");
            }
        }
    }
}

void vTaskPitch(void *pvParameters)
{
    for (;;)
    {
        vTaskDelay(1000 / portTICK_PERIOD_MS);
    }
}
#pragma endregion

void setup()
{
    // put your setup code here, to run once:
    Serial.begin(115200);
    alarmSemaphore = xSemaphoreCreateBinary();
    lcdSemaphore = xSemaphoreCreateMutex();
    lcd = new LiquidCrystal(15, 16, 17, 18, 8, 3);
    lcd->begin(16, 2);
    lcd->setCursor(0, 0); // Left: horizontal, right: vertical (16x2)
    printToLCDWithMutex("hi", 0, 0);
#if 1
    // Init timer
    timer = timerBegin(0, PRESCALER, true);
    timerAttachInterrupt(timer, &onTimer, true);
    timerAlarmWrite(timer, TICKS_PER_SECOND * 10, true);
    timerAlarmEnable(timer);
#if 0
    // Init GPIO interrupts
    pinMode(0, INPUT_PULLUP);
    attachInterrupt(0, onButton1, FALLING);

    pinMode(1, INPUT_PULLUP);
    attachInterrupt(1, onButton2, FALLING);

    pinMode(2, INPUT_PULLUP);
    attachInterrupt(2, onButton2, FALLING);
#endif
    deserialize(data);
    Serial.print("Time:" + data.time);
    setTMStructToTime(&data.time);
    // wifi. Write own function
    WiFi.mode(WIFI_STA);
    WiFi.begin(wifiSSID, wifiPassword);

    for (int loopCounter = 0; WiFiClass::status() != WL_CONNECTED; ++loopCounter)
    {
        Serial.println("Connecting...");
        printToLCDWithMutex("Connecting...", 0, 1);

        sleep(1);

        if (loopCounter == 5)
        {
            Serial.println("Failed to connect.");
            exit(69);
        }
    }
    lcd->clearLine(1);
    lcd->setCursor(0, 0);
    configTime(gmtOffset_sec, daylightOffset_sec, ntpServer);
#endif
#if 0
    // Poll events from the server
    auto http = new HTTPClient();
    auto token = getBearerToken(*http, "test", "test");
    events = getEvents(*http, token);

    for (eventSpace::event i : events)
    {
        Serial.printf("%s\n", i.get_eventname().c_str());
        // alarms.push_back((time_t)i.get_eventtime());
    }

#endif

    lcd->clear();
    xTaskCreatePinnedToCore(vTaskMain, "Task1", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 2, &Task1, 1);
    xTaskCreatePinnedToCore(vTaskDebug, "Task2", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &Task2, 0);
    vTaskSuspend(Task2);

    xTaskCreatePinnedToCore(vTaskAlarm, "TaskAlarm", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskAlarm, 0);
    vTaskSuspend(TaskAlarm);

    xTaskCreatePinnedToCore(vTaskMetronome, "TaskMetronome", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskMetronome, 0);
    // vTaskSuspend(TaskMetronome);

    xTaskCreatePinnedToCore(vTaskPitch, "TaskPitch", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskPitch, 0);
    vTaskSuspend(TaskPitch);
}

void loop()
{
}