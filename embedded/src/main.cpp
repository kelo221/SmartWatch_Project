#include <Arduino.h>


#include "arduinoFFT.h"
 
#define SAMPLES 128             
#define SAMPLING_FREQUENCY 2048 
 
arduinoFFT FFT = arduinoFFT();
 
unsigned int samplingPeriod;
unsigned long microSeconds;
 
double vReal[SAMPLES]; 
double vImag[SAMPLES]; 
const int buttonPin = digitalRead(4); 
int count =0;


void setup() 
{
    Serial.begin(115200); 
    samplingPeriod = round(1000000*(1.0/SAMPLING_FREQUENCY)); 
      pinMode(buttonPin, INPUT_PULLUP);
  }

 
void loop() 
{  
    
    for(int i=0; i<SAMPLES; i++)
    {
        microSeconds = micros();    
     
        vReal[i] = analogRead(5); 
        vImag[i] = 0; 

       
        while(micros() < (microSeconds + samplingPeriod))
        {
          
        }
    }
     enum standardTuning {
        E2 = 82,
        A2  = 110,
        D3  = 147,
        G3 = 196,
        B3 = 250,
        E4 =  330,
    };
    static int button1_last = 0;
  int button1_state = digitalRead(buttonPin);
 
  while (button1_state != button1_last)
  {
    count++;
    button1_last = button1_state;
    if (count >= 4) {
       count = 1;
    }
  }

  
     if(count==1)
        Serial.println(E2);
        
     else if(count==2)
        Serial.println(A2);
        
     else if(count==3)
        Serial.println(D3);
        
  


  
 
   
    FFT.Windowing(vReal, SAMPLES, FFT_WIN_TYP_HAMMING, FFT_FORWARD);
    FFT.Compute(vReal, vImag, SAMPLES, FFT_FORWARD);
    FFT.ComplexToMagnitude(vReal, vImag, SAMPLES);

    
    double peak = FFT.MajorPeak(vReal, SAMPLES, SAMPLING_FREQUENCY);
    Serial.println(peak/2);     
 
    
    delay(2000); 
}
