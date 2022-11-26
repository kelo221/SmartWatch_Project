#pragma once

#include"alarms/pitches.hh"
#include <Arduino.h>

int melody[] = {
        NOTE_G4,
        NOTE_E4, NOTE_F4, NOTE_G4, NOTE_G4, NOTE_G4,
        NOTE_A4, NOTE_B4, NOTE_C5, NOTE_C5, NOTE_C5,
        NOTE_E4, NOTE_F4, NOTE_G4, NOTE_G4, NOTE_G4,
        NOTE_A4, NOTE_G4, NOTE_F4, NOTE_F4,
        NOTE_E4, NOTE_G4, NOTE_C4, NOTE_E4,
        NOTE_D4, NOTE_F4, NOTE_B3,
        NOTE_C4, END

};


int noteDurations[] = {
        8,
        8, 8, 4, 4, 4,
        8, 8, 4, 4, 4,
        8, 8, 4, 4, 4,
        8, 8, 4, 2,
        4, 4, 4, 4,
        4, 2, 4,
        1
};

void beep() {

    int size = sizeof(melody) / sizeof(int);
    for (int thisNote = 0; thisNote < size; thisNote++) {
        const int noteDuration = 900 / noteDurations[thisNote];

        tone(23, melody[thisNote], noteDuration);

        const int pauseBetweenNotes = noteDuration * 1.30;
        delay(pauseBetweenNotes);

        noTone(23);

    }
}