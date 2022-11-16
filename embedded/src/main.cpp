#include <Arduino.h>
#include <HardwareSerial.h>

void hello_task(void *pvParameter)
{
    printf("Hello world!\n");
    for (int i = 10; i >= 0; i--) {
        printf("Restarting in %d seconds...\n", i);
        vTaskDelay(1000 / portTICK_RATE_MS);
    }
    printf("Restarting now.\n");
    fflush(stdout);
    esp_restart();
}

// the setup function runs once when you press reset or power the board
void setup() {

    printf("Hello world!\n");
    // Now set up two tasks to run independently.
    xTaskCreate(&hello_task, "hello_task", 2048, NULL, 5, NULL);
}


void loop()
{
    // Empty. Things are done in Tasks.
}

