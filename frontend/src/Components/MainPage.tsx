import React, { useEffect, useState } from "react";
import { getData } from "./Services/FetchRequest";
import { TimedEvent } from "../DataFormats/DataFormats";
import { Button, Card, DataTable, Text } from "grommet";
import { Checkmark, Close, Trash, Unlink } from "grommet-icons";
import CreateEvent from "./CreateEvent";

type Props = {
  eventVisState: React.ComponentState;
  setEventVisState: React.Dispatch<React.SetStateAction<boolean>>;
};


function MainPage(props: Props) {

  function deleteButtonHandler(id: number): void {
    console.log("hello", id);
  }


  const [events, setEvents] = useState<TimedEvent[]>({} as TimedEvent[]);

  useEffect(() => {
    getData("http://localhost:8000/api/user/events").then((data) => {
      if (!data.status) {
        let jsonObject = JSON.parse(data) as TimedEvent[];
        setEvents(jsonObject);
      } else {
        console.log("failed!!!", data.status);
      }
    });
  }, []);

  if (!events.length) {
    return (
      <Card
        margin="small"
        gap="medium"
        width="medium"
        align="center"
        justify="center"
        alignContent="row"
        fill
      >
        <Unlink></Unlink>
        <Text>Check your connection, you may also have a stale bearer token. Log in again to fresh one.</Text>
        <Button primary label="Log out" onClick={() => {
          (localStorage.removeItem("token"));
          window.location.replace(window.location.origin);
        }} />
      </Card>
    );
  }

  return (
    <>
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
                return <Text>{new Date(Date.parse(datum.eventtime.toLocaleString())).toLocaleString()}</Text>;
              }
            },
            {
              property: "passed",
              header: <Text>Passed?</Text>,
              render: datum => {
                if ((Date.parse(datum.eventtime.toLocaleString())) < (new Date().valueOf())) {
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
                  <span onClick={() => deleteButtonHandler(datum.id)}
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

