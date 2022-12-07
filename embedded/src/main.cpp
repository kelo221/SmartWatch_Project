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
#include "arduinoFFT.h"
#define CLK_FREQ 80000000
#define PRESCALER 80
#define TICKS_PER_SECOND (CLK_FREQ / PRESCALER)
#define SSID_NAME "abcd"
#define SSID_PASSWORD "12345678"
#define BEEPERPIN 21
#define SAMPLES 128
#define SAMPLING_FREQUENCY 2048
const char *wifiSSID = SSID_NAME;
const char *wifiPassword = SSID_PASSWORD;

//  eventSpace::events data = nlohmann::json::parse(testJson);

/* TODO:
    Add GPIO interrupts for all three buttons
    Configure buttons pins
    Revise task management
    Implement pitch
    Implement metronome
*/
static HTTPClient *httpclient;
static std::string token;
// LCD
static LiquidCrystal *lcd;
// Core variables
static const std::vector<std::string> menuItmes = {"Metronome", "Tuner", "Update"};
static volatile bool button1Pressed = false;
static volatile bool button2Pressed = false;
static volatile bool button3Pressed = false;
static volatile bool setupFlag = true;
// FreeRTOS parameters
static uint8_t ucParameterToPass;
static TaskHandle_t Task1;
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
    if (!events.empty())
    {
        auto it = events.begin();
        while (it != events.end())
        {
            if (it->get_eventtime() > currentTime - 2 && it->get_eventtime() < currentTime + 2)
            {
                it = events.erase(it);
                if (!deleteEvent(*httpclient, token, it->get_eventtime()))
                {
                    return false;
                };
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
        if (i == 9)
        {
            printLocalTime();
            i = 0;
        }

#pragma region Buttons

        // Stop task when button 1&3 are pressed when a task is running.
        if (button1Pressed && button3Pressed)
        {
            button1Pressed = false;
            button3Pressed = false;
            if (CurrentTask != nullptr)
            {
                vTaskSuspend(*CurrentTask);
                CurrentTask = nullptr;
                setupFlag = true;
                Serial.println("Suspend Task");
                clearLineWithMutex(1);
            }
        }

        // Only select and execute tasks when no task is running
        if (CurrentTask == nullptr)
        {
            if (button1Pressed)
            {
                Serial.print("Button 1");
                button1Pressed = false;
                menuIndex--;
                if (menuIndex < 0)
                {
                    menuIndex = menuItmes.size() - 1;
                }
                clearLineWithMutex(1);
                printToLCDWithMutex(menuItmes[menuIndex].c_str(), 0, 1);
            }

            if (button3Pressed)
            {
                Serial.print("Button3");
                button3Pressed = false;
                menuIndex++;
                if (menuIndex > menuItmes.size() - 1)
                {
                    menuIndex = 0;
                }
                clearLineWithMutex(1);
                printToLCDWithMutex(menuItmes[menuIndex].c_str(), 0, 1);
            }

            if (button2Pressed)
            {
                Serial.print("Button 2");
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
                    events = getEvents(*httpclient, token);
                    break;
                default:
                    break;
                }
                menuIndex = 0;
            }
        }
#pragma endregion

        if (checkAlarms(data.time))
        {
            vTaskResume(TaskAlarm);
        }
        vTaskDelay(100 / portTICK_PERIOD_MS);
        i++;
    }
}
// handles the alarm
void vTaskAlarm(void *pvParameters)
{
    // start beeping
    Serial.println("ALARM!");
    int size = sizeof(melody) / sizeof(int);
    for (int thisNote = 0; thisNote < size; thisNote++)
    {
        const int noteDuration = 900 / noteDurations[thisNote];

        tone(BEEPERPIN, melody[thisNote], noteDuration);

        const int pauseBetweenNotes = noteDuration * 1.30;
        vTaskDelay(pauseBetweenNotes / portMAX_DELAY);
        noTone(BEEPERPIN);
    }

    // xSemaphoreGive(alarmSemaphore);
    CurrentTask = nullptr;
    vTaskSuspend(NULL); // suspend oneself
}

void vTaskMetronome(void *pvParameters)
{
    // set up metronome
    int bpm = 60;
    std::string s = "BPM: ";
    for (;;)
    {
        if (setupFlag)
        {
            setupFlag = false;
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
            button2Pressed = false;
        }
        for (int i = 0; i < 16; i++)
        {
            vTaskDelay((60000 / bpm) / portTICK_PERIOD_MS);

            clearLineWithMutex(1);
            printToLCDWithMutex("-", i, 1);
            tone(BEEPERPIN, 440, 100);
            Serial.println("Metronome");
        }
    }
}

void vTaskPitch(void *pvParameters)
{
    arduinoFFT FFT = arduinoFFT();

    unsigned int samplingPeriod;
    unsigned long microSeconds;

    double vReal[SAMPLES];
    double vImag[SAMPLES];
    samplingPeriod = round(1000000 * (1.0 / SAMPLING_FREQUENCY));

    for (;;)
    {
        for (int i = 0; i < SAMPLES; i++)
        {
            microSeconds = micros(); // Returns the number of microseconds since the Arduino board began running the current script.

            vReal[i] = analogRead(5); // Reads the value from analog pin 0 (A0), quantize it and save it as a real term.
            vImag[i] = 0;             // Makes imaginary term 0 always

            /*remaining wait time between samples if necessary*/
            while (micros() < (microSeconds + samplingPeriod))
            {
                // do nothing
            }
        }
        /*Perform FFT on samples*/
        FFT.Windowing(vReal, SAMPLES, FFT_WIN_TYP_HAMMING, FFT_FORWARD);
        FFT.Compute(vReal, vImag, SAMPLES, FFT_FORWARD);
        FFT.ComplexToMagnitude(vReal, vImag, SAMPLES);

        /*Find peak frequency and print peak*/
        double peak = FFT.MajorPeak(vReal, SAMPLES, SAMPLING_FREQUENCY);

        Serial.println(peak);
        char output[16];
        snprintf(output, 16, "%f", peak / 2);
        printToLCDWithMutex(output, 0, 1);

        vTaskDelay(500 / portTICK_PERIOD_MS);
        Serial.println(output);
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

    // Init timer
    timer = timerBegin(0, PRESCALER, true);
    timerAttachInterrupt(timer, &onTimer, true);
    timerAlarmWrite(timer, TICKS_PER_SECOND * 10, true);
    timerAlarmEnable(timer);
    // Init GPIO interrupts
    pinMode(1, INPUT_PULLUP);
    attachInterrupt(1, onButton1, FALLING);

    pinMode(2, INPUT_PULLUP);
    attachInterrupt(2, onButton2, FALLING);

    pinMode(42, INPUT_PULLUP);
    attachInterrupt(42, onButton3, FALLING);
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
#if 1
    // Poll events from the server
    httpclient = new HTTPClient();
    token = getBearerToken(*httpclient, "admin", "1234");
    if (token == "")
    {
        exit(69);
    }
    events = getEvents(*httpclient, token);

    for (eventSpace::event i : events)
    {
        Serial.printf("%s\n", i.get_eventname().c_str());
        // alarms.push_back((time_t)i.get_eventtime());
    }

#endif

    lcd->clear();
    printToLCDWithMutex(menuItmes[0], 0, 1);
    xTaskCreatePinnedToCore(vTaskMain, "Task1", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 2, &Task1, 1);

    xTaskCreatePinnedToCore(vTaskAlarm, "TaskAlarm", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskAlarm, 0);
    vTaskSuspend(TaskAlarm);

    xTaskCreatePinnedToCore(vTaskMetronome, "TaskMetronome", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskMetronome, 0);
    vTaskSuspend(TaskMetronome);

    xTaskCreatePinnedToCore(vTaskPitch, "TaskPitch", 2048 * 2, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskPitch, 0);
    vTaskSuspend(TaskPitch);
}

void loop()
{
}