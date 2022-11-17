import create from "zustand";
import { StatusType } from "grommet/components/Notification/index";

interface notificationInterface {
  notificationVisibility: boolean;
  notificationText: string;
  notificationAlertLevel: StatusType;

  setNotificationVis: (vis: boolean) => void;
  setNotificationText: (text: string) => void;
  setNotificationAlertLevel: (text: StatusType) => void;
}

export const notificationStore = create<notificationInterface>()((set) => ({

  notificationVisibility: false,
  notificationText: "",
  notificationAlertLevel: "unknown",

  setNotificationVis: (vis: boolean) => {
    set(() => ({
      notificationVisibility: vis
    }));
  },
  setNotificationText: (text: string) => {
    set(() => ({
      notificationText: text
    }));
  },

  setNotificationAlertLevel: (text: StatusType) => {
    set(() => ({
      notificationAlertLevel: text
    }));
  }

}));