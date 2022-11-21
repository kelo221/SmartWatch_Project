export interface TimedEvent {
  id: number;
  eventname: string;
  eventtime: number;
  eventid: number;
  issilent: boolean;
}

export interface NewEventUpload {
  eventName: string;
  eventTime: string;
  isSilent: string;
  SnoozeDisabled : string;
}

export interface ModifiedEventUpload {
  eventName:        string;
  eventTimeOld:     string;
  eventTimeNewUnix: string;
  isSilent:         string;
  SnoozeDisabled:   string;
}



export interface RegisterFormJSON {
  username: string;
  password: string;
  password_confirm: string;
}
