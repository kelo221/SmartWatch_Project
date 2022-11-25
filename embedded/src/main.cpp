#include <Arduino.h>
#include <WiFi.h>
#include "EEPROM.h"
#include <vector>
#include <LiquidCrystal.h>
#define CLK_FREQ 80000000
#define PRESCALER 80
#define TICKS_PER_SECOND (CLK_FREQ / PRESCALER)
/* TODO:
    De/Serialize timer UPDATE: Difficult to set up
    Integrate wifi connection UPDATE: It was easy
    Alarm funcitonality
    Set up proper project i.e, clone
*/

// LCD
static LiquidCrystal *lcd;
// FreeRTOS parameters
static uint8_t ucParameterToPass;
static TaskHandle_t Task1;
static TaskHandle_t Task2;
static TaskHandle_t TaskAlarm;
static TaskHandle_t TaskMetronome;
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

// Timer
hw_timer_t *timer = NULL;

#pragma region Function definitons
void vTaskDebug(void *pvParameters);
void vTaskMain(void *pvParameters);
void vTaskMetronome(void *pvParameters);
void printToLCDWithMutex(const char *str, uint8_t col, uint8_t line);
#pragma endregion

#pragma region Functions
void printLocalTime()
{
    static char strtime_buf[64];
    setenv("FI", "GMT-2", 1);
    tzset();

    // time_t now;
    time(&data.time);
    localtime_r(&data.time, &timeinfo);
    strftime(strtime_buf, sizeof(strtime_buf), "  %a %H:%M:%S \n", &timeinfo);

    printToLCDWithMutex(strtime_buf, 0, 0);
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

void IRAM_ATTR onTimer()
{
}

bool checkAlarms(const time_t currentTime)
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
    return false;
}

void printToLCDWithMutex(const char *str, uint8_t col, uint8_t line)
{
    if (xSemaphoreTake(lcdSemaphore, portMAX_DELAY) == pdTRUE)
    {
        lcd->setCursor(col, line);
        lcd->print(str);
        xSemaphoreGive(lcdSemaphore);
    }
}

#pragma endregion

#pragma region Tasks

void vTaskMain(void *pvParameters)
{
    Serial.println(xPortGetCoreID());
    bool alarmRunning = false;
    int i = 0;
    for (;;)
    {
        printLocalTime();

        if (checkAlarms(data.time))
        {
            if (xSemaphoreTake(alarmSemaphore, portMAX_DELAY) == pdTRUE)
            {
                vTaskResume(TaskAlarm);
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
    int bpm = 60;

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
    auto events = getEvents(*http, token);

    for (eventSpace::event i : events)
    {
        Serial.printf("%s\n", i.get_eventname().c_str());
        // alarms.push_back((time_t)i.get_eventtime());
    }

#endif

    lcd->clear();
    xTaskCreatePinnedToCore(vTaskMain, "Task1", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 2, &Task1, 1);
    xTaskCreatePinnedToCore(vTaskDebug, "Task2", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &Task2, 0);
    xTaskCreatePinnedToCore(vTaskAlarm, "TaskAlarm", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskAlarm, 0);
    xTaskCreatePinnedToCore(vTaskMetronome, "TaskMetronome", 2048, &ucParameterToPass, tskIDLE_PRIORITY + 1, &TaskMetronome, 0);

    vTaskSuspend(Task2);
    vTaskSuspend(TaskAlarm);

    Serial.println("Tasks set!");

    Serial.println("Suspend Debug Task");
}
<<<<<<< HEAD
#pragma endregion

void loop()
{
    // put your main code here, to run repeatedly:
    == == == =

                 void loop()
    {
>>>>>>> 501ff3e (Alarm functionality)
    }