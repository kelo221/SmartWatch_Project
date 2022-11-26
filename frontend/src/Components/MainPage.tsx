import React, { useEffect, useState } from "react";
import { fetchRequest } from "./Services/FetchRequest";
import { TimedEvent } from "../DataFormats/DataFormats";
import { Box, Button, Card, DataTable, Notification, Text } from "grommet";
import { Checkmark, Close, LinkUp, Trash, Unlink } from "grommet-icons";
import NewEventPopUp from "./NewEventPopUp";
import { eventStore } from "../Stores/eventStore";
import { notificationStore } from "../Stores/notificationStore";
import { popUpStore } from "../Stores/popUpStore";

type Props = {
  eventVisState: React.ComponentState;
  setEventVisState: React.Dispatch<React.SetStateAction<boolean>>;
};





function MainPage(props: Props) {

  const [queryCheck, setQueryCheck] = useState(false);
  const { removeEvent, events, initEvents } = eventStore();
  const { setAsNewEvent, setEventDate, setEventName, setSilent, setTimeString, setOldTime, oldTime } = popUpStore();

  const {
    setNotificationVis, setNotificationAlertLevel, notificationAlertLevel,
    notificationText, notificationVisibility, setNotificationText
  } = notificationStore();

  function deleteButtonHandler(id: number, unixTime: number): void {

    console.log(unixTime);

    fetchRequest("http://localhost:8000/api/user/eventDate", { "eventDate": unixTime.toString() }, "DELETE").then((data) => {

      setNotificationVis(true);

      if (!data.status) {
        setNotificationAlertLevel("normal");
        setNotificationText("OK! Event deleted.");
        removeEvent(id);
      } else {
        setNotificationAlertLevel("critical");
        setNotificationText(data.errorText + " (" + data.status + "). Failed to delete.");
      }
    });
  }


  useEffect(() => {
    fetchRequest("http://localhost:8000/api/user/events", {}, "GET").then((data) => {
      if (data !== "null" && data !== null && !data.status) {
        console.log(data);
        let jsonObject = data as TimedEvent[];
        initEvents(jsonObject);
      } else if (data === "null" && data.status) {
        console.log("failed!!!", data.status, data);
        setNotificationText("No events found!");
        setNotificationVis(true);
        setNotificationAlertLevel("critical");
        setQueryCheck(true);
      }
    });
  }, []);

  if (queryCheck && events.length === 0) {
    return (
      <Card
        pad={"small"}
        margin="small"
        gap="medium"
        width="medium"
        align="center"
        justify="center"
        alignContent="row"
        fill
      >
        <Unlink></Unlink>
        <Text>Check your connection.</Text>
        {localStorage.getItem("token") ? (
          <><Text>Log in again to get a fresh bearer token.</Text><Button primary label="Log out" onClick={() => {
            (localStorage.removeItem("token"));
            window.location.replace(window.location.origin);
          }} /></>
        ) : <></>
        }
      </Card>
    );
  }


  if (events.length === 0) {
    return (
      <Card
        pad={"small"}
        margin="small"
        gap="medium"
        width="medium"
        align="center"
        justify="center"
        alignContent="row"
        fill
      >
        {props.eventVisState ?
          <NewEventPopUp eventVisState={props.eventVisState}
                         setEventVisState={props.setEventVisState}></NewEventPopUp> :
          <></>
        }
        <LinkUp></LinkUp>
        <Text>Create some events!</Text>
      </Card>
    );
  }

  return (
    <>
      {notificationVisibility && (
        <>
          <Box align="center" gap="small">
            <Notification
              toast={{ position: "top" }}
              status={notificationAlertLevel}
              message={notificationText}
              onClose={() => setNotificationVis(true)} />
          </Box>
        </>
      )}
      <Card
        margin="small"
        gap="medium"
        width="medium"
        align="center"
        justify="center"
        fill
      >
        {props.eventVisState ?
          <NewEventPopUp eventVisState={props.eventVisState}
                         setEventVisState={props.setEventVisState}></NewEventPopUp> :
          <></>
        }
        <DataTable
          style={
            {
              filter: props.eventVisState ? `blur(5px)` : `blur(0px)`
            }
          }
          columns={[
            {
              property: "eventname",
              header: <Text>Name</Text>,
              primary: true
            },
            {
              property: "eventtime",
              header: <Text>Event Date</Text>,
              render: datum => {
                return <Text>{new Date(datum.eventtime * 1000).toLocaleString()}</Text>;
              }
            },
            {
              property: "passed",
              header: <Text>Passed?</Text>,
              render: datum => {
                if (new Date(new Date(datum.eventtime * 1000)).valueOf() < (new Date().valueOf())) {
                  return <Checkmark></Checkmark>;
                } else {
                  return <Close></Close>;
                }
              }
            },
            {
              property: "issilent",
              header: <Text>Silent?</Text>,
              render: datum => {
                if (datum.issilent) {
                  return <Checkmark></Checkmark>;
                } else {
                  return <Close></Close>;
                }
              }
            },
            {
              property: "deleteButton",
              header: <Text>Delete?</Text>,
              render: datum => {
                return (
                  <span onClick={() => deleteButtonHandler(datum.id, datum.eventtime)}
                        id={datum.id.toString()}
                  >
                <Trash style={{ cursor: "pointer" }}></Trash>
                </span>);
              }
            }
          ]}
          data={events}
          fill
          pin
          sortable
        />
      </Card>
    </>
  );
}

export default MainPage;

