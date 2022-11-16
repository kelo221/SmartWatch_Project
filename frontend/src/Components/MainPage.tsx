import React, { useEffect, useState } from "react";
import { fetchRequest } from "./Services/FetchRequest";
import { TimedEvent } from "../DataFormats/DataFormats";
import { Box, Button, Card, DataTable, Notification, Text } from "grommet";
import { Checkmark, Close, Trash, Unlink } from "grommet-icons";
import CreateEvent from "./CreateEvent";
import { StatusType } from "grommet/components/Notification/index";
import { eventStore } from "../Store/store";

type Props = {
  eventVisState: React.ComponentState;
  setEventVisState: React.Dispatch<React.SetStateAction<boolean>>;
};


function padTo2Digits(num: number) {
  return num.toString().padStart(2, "0");
}

function formatDate(date: Date) {
  return (
    [
      date.getFullYear(),
      padTo2Digits(date.getMonth() + 1),
      padTo2Digits(date.getDate())
    ].join("-") +
    " " +
    [
      padTo2Digits(date.getHours()),
      padTo2Digits(date.getMinutes()),
      padTo2Digits(date.getSeconds())
    ].join(":")
  );
}


/// TODO
// Check if there's a way to get the id so that the delete function works, or maybe just change it to delete by a
// specific date instead of the id, as that is unique as well
// Make a separate store for popups? to clean up the codebase a bit?

function MainPage(props: Props) {

  const [loginSuccess, setLoginSuccess] = React.useState<StatusType>("normal");
  const [alertMessage, setAlertMessage] = React.useState("hello nothing here");
  const [visible, setVisible] = useState(false);

  const { removeEvent, events, initEvents } = eventStore();

  function deleteButtonHandler(id: number, date: Date): void {

    const fixedHours = new Date(date).setHours(new Date(date).getHours() - 2);

    fetchRequest("http://localhost:8000/api/user/eventDate", { "eventDate": formatDate(new Date(fixedHours)).toString() }, "DELETE").then((data) => {
      if (!data.status) {
        setLoginSuccess("normal");
        setAlertMessage("OK! Event deleted.");
        removeEvent(id);
        setVisible(true);
      } else {
        setLoginSuccess("critical");
        setAlertMessage(data.errorText + " (" + data.status + "). Failed to delete.");
        setVisible(true);
      }
    });
  }


  useEffect(() => {
    fetchRequest("http://localhost:8000/api/user/events", {}, "GET").then((data) => {
      if (!data.status) {
        let jsonObject = JSON.parse(data) as TimedEvent[];
        initEvents(jsonObject);
      } else {
        console.log("failed!!!", data.status);
      }
    });
  }, []);

  if (events === null) {
    return (

      <><Box align="center" gap="small">
        {visible && (
          <Notification
            toast={{ position: "top" }}
            status={loginSuccess}
            message={alertMessage}
            onClose={() => setVisible(false)} />
        )}
      </Box><Card
        margin="small"
        gap="medium"
        width="medium"
        align="center"
        justify="center"
        fill
      >

        {props.eventVisState ?
          <CreateEvent eventVisState={props.eventVisState} setEventVisState={props.setEventVisState}></CreateEvent> :
          <></>}
        Nothing here, create some events!</Card></>);
  }

  if (!events.length) {
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

  return (
    <>
      {visible && (
        <>
          <Box align="center" gap="small">
            <Notification
              toast={{ position: "top" }}
              status={loginSuccess}
              message={alertMessage}
              onClose={() => setVisible(false)} />
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
          <CreateEvent eventVisState={props.eventVisState} setEventVisState={props.setEventVisState}></CreateEvent> :
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
                return <Text>{new Date(new Date(datum.eventtime).toISOString().slice(0, -1)).toLocaleString()}</Text>;
              }
            },
            {
              property: "passed",
              header: <Text>Passed?</Text>,
              render: datum => {
                if (new Date(new Date(datum.eventtime).toISOString().slice(0, -1)).valueOf() < (new Date().valueOf())) {
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

