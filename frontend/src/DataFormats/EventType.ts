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
}
