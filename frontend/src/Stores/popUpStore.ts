import create from "zustand";

interface popUpInterface {
  eventName: string;
  eventDate: Date;
  isSilent: boolean;
  snoozeDisabled: boolean;
  isNewEvent: boolean;
  timeString: string;
  oldTime: string;

  setEventName: (newName: string) => void;
  setEventDate: (newDate: Date) => void;
  setSilent: (mode: boolean) => void;
  setSnooze: (mode: boolean) => void;
  setAsNewEvent: (mode: boolean) => void;
  setTimeString: (time: string) => void;
  setOldTime: (time: string) => void;

}

export const popUpStore = create<popUpInterface>()((set) => ({

  eventName: "",
  eventDate: new Date(),
  isSilent: false,
  snoozeDisabled: false,
  isNewEvent: true,
  timeString: "",
  oldTime: "",

  setSilent: (mode: boolean) => {
    set(() => ({
      isSilent: mode
    }));
  },

  setAsNewEvent: (mode: boolean) => {
    set(() => ({
      isNewEvent: mode
    }));
  },

  setEventDate: (newDate: Date) => {
    set(() => ({
      eventDate: newDate
    }));
  },


  setSnooze: (mode: boolean) => {
    set(() => ({
      snoozeDisabled: mode
    }));
  },

  setEventName: (text: string) => {
    set(() => ({
      eventName: text
    }));
  },

  setTimeString: (time: string) => {
    set(() => ({
      timeString: time
    }));
  },

  setOldTime: (text: string) => {
    set(() => ({
      oldTime: text
    }));
  }


}));