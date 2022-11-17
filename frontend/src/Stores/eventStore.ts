import create from "zustand";
import { TimedEvent } from "../DataFormats/DataFormats";


interface timedEventInterface {
  events: TimedEvent[];
  addEvent: (newEvent: TimedEvent) => void;
  initEvents: (newEvents: TimedEvent[]) => void;
  removeEvent: (id: number) => void;
}

export const eventStore = create<timedEventInterface>()((set) => ({

  events: [],

  addEvent: (newEvent: TimedEvent) => {
    set((state) => ({
      events: [
        ...state.events,
        {
          ...newEvent
        } as TimedEvent
      ]
    }));
  },
  initEvents: (newEvents: TimedEvent[]) => {
    set(() => ({
      events: [
        ...newEvents
      ]
    }));
  },
  removeEvent: (id: number) => {
    set((state) => ({
      events: state.events.filter((todo) => todo.id !== id)
    }));
  }

}));