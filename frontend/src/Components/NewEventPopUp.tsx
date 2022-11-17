import { Box, Button, Card, CheckBox, DateInput, Form, FormField, MaskedInput } from "grommet";
import React from "react";
import { NewEventUpload } from "../DataFormats/DataFormats";
import { fetchRequest } from "./Services/FetchRequest";
import { Close } from "grommet-icons";
import { eventStore } from "../Stores/eventStore";
import { notificationStore } from "../Stores/notificationStore";
import { popUpStore } from "../Stores/popUpStore";

const hourRegex = /^\b2[0-3]\b|\b[0-1]?[0-9]\b$/;
const minuteRegex = /^[0-5]?[0-9]$/;

type Props = {
  eventVisState: React.ComponentState;
  setEventVisState: React.Dispatch<React.SetStateAction<boolean>>;
};

const maskTime = [
  {
    length: [1, 2],
    regexp: hourRegex,
    placeholder: "xx"
  },
  { fixed: ":" },
  {
    length: [1, 2],
    regexp: minuteRegex,
    placeholder: "xx"
  }
];

const CreateEvent = (props: Props) => {


  const { addEvent } = eventStore();
  const {
    setAsNewEvent, setEventDate, setEventName, eventName, eventDate, isNewEvent, isSilent, setSilent, setTimeString,
    timeString, setSnooze, snoozeDisabled, oldTime
  } = popUpStore();

  const { setNotificationVis, setNotificationAlertLevel, setNotificationText } = notificationStore();

  const formatAndUpload = () => {

    try {
      eventDate.setHours(Number(timeString.split(":")[0]), Number(timeString.split(":")[1]), 0);
    } catch (e) {
      console.log(e);
    }

    const newEvent: NewEventUpload = {
      eventName: eventName,
      eventTime: (Math.floor(eventDate.getTime() / 1000)).toString(),
      isSilent: isSilent.toString(),
      SnoozeDisabled: snoozeDisabled.toString()
    };

    fetchRequest("http://localhost:8000/api/user/event", newEvent).then((data) => {
      setNotificationVis(true);
      if (!data.status) {
        setNotificationAlertLevel("normal");
        setNotificationText("OK! Event created.");
        props.setEventVisState(false);
        addEvent({
          id: 0,
          created_at: new Date(),
          eventtime: new Date(eventDate.getTime()),
          eventid: 0,
          eventname: eventName,
          issilent: isSilent

        });
      } else {
        setNotificationAlertLevel("critical");
        setNotificationText(data.errorText + " (" + data.status + "). Failed to create.");
      }
    });

  };

  return (
    <Card
      pad={"small"}
      align="center"
      style={{
        position: "absolute",
        zIndex: 2,
        width: "300px",
        height: "400px",
        top: "100px"
      }}
    >
      <Close
        style={{
          position: "absolute",
          right: "30px",
          cursor: "pointer"
        }}
        onClick={() => props.setEventVisState(false)}
      ></Close>

      <Form>
        <FormField
          label="Event Name"
          htmlFor="masked-event-name"
          name="maskedSizeUnits"
        >
          <MaskedInput
            name="maskedEventName"
            id="masked-event-name"
            mask={
              [
                { placeholder: "Enter the name of the event" }
              ]
            }
            value={eventName}
            onChange={event => setEventName(event.target.value)}
          />
        </FormField>

        <FormField
          label="Event Time and Date"
          htmlFor="masked-event-time"
          name="maskedSizeUnits"
        >
          <MaskedInput
            name="maskedDateTime"
            id="masked-date-time"
            value={timeString}
            mask={maskTime}
            onChange={event => {
              setTimeString(event.target.value);
            }}
          />
        </FormField>

        <Box>
          <DateInput
            height={"30px"}
            format="yyyy-mm-dd"
            value={new Date(eventDate).toISOString().substring(0, 10)}
            onChange={({ value }) => {

              if (value.length) {
                setEventDate(new Date(value[0].toString()));
              }
            }}
          />
        </Box>

        <FormField
          label="Options"
          htmlFor="masked-event-time"
          name="maskedSizeUnits"
        >
          <CheckBox
            checked={isSilent}
            label="Silent?"
            onChange={(event) => setSilent(event.target.checked)}
          />
          <CheckBox
            checked={snoozeDisabled}
            label="Disable Snooze"
            onChange={(event) => setSnooze(event.target.checked)}
          />
        </FormField>


        <Box direction="row" gap="small" pad={"small"}>
          <Button type="submit" label="Send to Smartwatch" primary onClick={() => {
            if (eventName !== "") {
              formatAndUpload();
            } else {
              /// TODO
              console.log("implement a check for the empty field");
            }
          }} />
          <Button type="reset" label="Reset" />
        </Box>
      </Form>

    </Card>
  );
};

export default CreateEvent;