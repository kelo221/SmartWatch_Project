export interface TimedEvent {
  id: number;
  eventname: string;
  eventtime: Date;
  created_at: Date;
  eventid: number;
  issilent: boolean | null;
}

export interface NewEventUpload {
  eventName: string;
  eventTime: string;
  isSilent: string;
  SnoozeDisabled : string;
}


export interface RegisterFormJSON {
  username: string;
  password: string;
  password_confirm: string;
}
