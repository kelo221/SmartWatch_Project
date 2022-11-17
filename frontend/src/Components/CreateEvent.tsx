import { Box, Button, Card, CheckBox, DateInput, Form, FormField, MaskedInput } from "grommet";
import React from "react";
import { NewEventUpload } from "../DataFormats/DataFormats";
import { fetchRequest } from "./Services/FetchRequest";
import { Close } from "grommet-icons";
import { eventStore } from "../Stores/eventStore";
import { notificationStore } from "../Stores/notificationStore";


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

  const [eventNameLocal, setEventNameLocal] = React.useState("");
  const [isSilentLocal, setIsSilentLocal] = React.useState(false);
  const [snoozeDisable, setSnoozeAble] = React.useState(false);
  const [date, setDate] = React.useState(new Date());
  const [timeString, setTimeString] = React.useState("");

  const { addEvent } = eventStore();

  const { setNotificationVis, setNotificationAlertLevel, setNotificationText } = notificationStore();

  const formatAndUpload = () => {


    date.setHours(Number(timeString.split(":")[0]), Number(timeString.split(":")[1]), 0);


    const newEvent: NewEventUpload = {
      eventName: eventNameLocal,
      eventTime: (Math.floor(date.getTime() / 1000)).toString(),
      isSilent: isSilentLocal.toString(),
      SnoozeDisabled: snoozeDisable.toString()
    };

    console.log(newEvent);

    fetchRequest("http://localhost:8000/api/user/event", newEvent).then((data) => {
      setNotificationVis(true);
      if (!data.status) {
        setNotificationAlertLevel("normal");
        setNotificationText("OK! Event created.");
        props.setEventVisState(false);
        addEvent({
          id: 0,
          created_at: new Date(),
          eventtime: new Date(date.getTime()),
          eventid: 0,
          eventname: eventNameLocal,
          issilent: isSilentLocal

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
            value={eventNameLocal}
            onChange={event => setEventNameLocal(event.target.value)}
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
            onChange={event => setTimeString(event.target.value)}
          />
        </FormField>

        <Box>
          <DateInput
            height={"30px"}
            format="yyyy-mm-dd"
            value={date.toISOString().substring(0, 10)}
            onChange={({ value }) => {

              if (value.length) {
                setDate(new Date(value[0].toString()));
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
            checked={isSilentLocal}
            label="Silent?"
            onChange={(event) => setIsSilentLocal(event.target.checked)}
          />
          <CheckBox
            checked={snoozeDisable}
            label="Disable Snooze"
            onChange={(event) => setSnoozeAble(event.target.checked)}
          />
        </FormField>


        <Box direction="row" gap="small" pad={"small"}>
          <Button type="submit" label="Send to Smartwatch" primary onClick={() => {
            if (eventNameLocal !== "") {
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