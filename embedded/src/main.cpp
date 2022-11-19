#include <Arduino.h>
#include <WiFi.h>
#include "EEPROM.h"
#define CLK_FREQ 80000000
#define PRESCALER 80
#define TICKS_PER_SECOND (CLK_FREQ / PRESCALER)
/* TODO:
    De/Serialize timer UPDATE: Difficult to set up
    Integrate wifi connection UPDATE: It was easy
    Alarm funcitonality
    Set up proper project i.e, clone
*/

// FreeRTOS parameters
static uint8_t ucParameterToPass;
TaskHandle_t Task1;
TaskHandle_t Task2;

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

// Timer
hw_timer_t *timer = NULL;

#pragma region Functions
void printLocalTime()
{
    static char strtime_buf[64];
    setenv("FI", "GMT-2", 1);
    tzset();

    // time_t now;
    time(&data.time);
    localtime_r(&data.time, &timeinfo);
    strftime(strtime_buf, sizeof(strtime_buf), "%c", &timeinfo);

    Serial.println(strtime_buf);
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
    serialize(data);
}

#pragma endregion

#pragma region Tasks

void vTaskMain(void *pvParameters)
{
    Serial.println(xPortGetCoreID());

    for (;;)
    {
        printLocalTime();
        delay(1000);
    }
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
#pragma endregion

void setup()
{
    // put your setup code here, to run once:
    Serial.begin(115200);

    // Init timer
    timer = timerBegin(0, PRESCALER, true);
    timerAttachInterrupt(timer, &onTimer, true);
    timerAlarmWrite(timer, TICKS_PER_SECOND * 10, true);
    timerAlarmEnable(timer);

    deserialize(data);
    Serial.print("Time:" + data.time);
    setTMStructToTime(&data.time);

#if 0
  // wifi. Write own function
  WiFi.mode(WIFI_STA);
  WiFi.begin("Simon’s iPhone", "wizard12");
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println(" CONNECTED");
  // Update time
  configTime(gmtOffset_sec, daylightOffset_sec, ntpServer);
#endif
    xTaskCreatePinnedToCore(vTaskMain, "Task1", configMINIMAL_STACK_SIZE * 3, &ucParameterToPass, tskIDLE_PRIORITY + 1, &Task1, 1);
    xTaskCreatePinnedToCore(vTaskDebug, "Task2", configMINIMAL_STACK_SIZE * 3, &ucParameterToPass, tskIDLE_PRIORITY, &Task2, 0);
}

void loop()
{
    // put your main code here, to run repeatedly:
}