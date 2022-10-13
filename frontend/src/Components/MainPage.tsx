import React, { useEffect, useState } from "react";
import { getData } from "./FetchRequest";
import { TimedEvent } from "../DataFormats/EventType";
import { Box, Table, TableBody, TableCell, TableHeader, TableRow } from "grommet";


function MainPage() {

  const [events, setEvents] = useState<TimedEvent[]>({} as TimedEvent[]);

  useEffect(() => {
    getData("http://localhost:8000/api/user/events").then((data) => {
      if (!data.status) {
        let jsonObject = JSON.parse(data) as TimedEvent[];
        console.log(jsonObject);
        setEvents(jsonObject);
      } else {
        console.log("failed!!!", data.status);
      }
    });
  }, []);


  const eventsLoop = Object.entries(events).map(([index, value]) =>
    <TableRow key={index}>
      <TableCell scope="row">
        {value.eventname}
      </TableCell>
      <TableCell>{new Date(Date.parse(value.eventtime.toLocaleString())).toLocaleString()}</TableCell>
      <TableCell>{value.issilent ? "true" : "false"}</TableCell>
      <TableCell>{(Date.parse(value.eventtime.toLocaleString())) < (new Date().valueOf()) ? "true" : "false"}</TableCell>
    </TableRow>
  );

  return (

    <Box
      margin="small"
      gap="medium"
      width="medium"
      align="center"
      justify="center"
      fill
    >

      <Table>
        <TableHeader>
          <TableRow>
            <TableCell scope="col" border="bottom">
              Event Name
            </TableCell>
            <TableCell scope="col" border="bottom">
              Event Date
            </TableCell>
            <TableCell scope="col" border="bottom">
              Silent?
            </TableCell>
            <TableCell scope="col" border="bottom">
              Passed?
            </TableCell>
          </TableRow>
        </TableHeader>
        <TableBody>
          {eventsLoop}
        </TableBody>
      </Table>
    </Box>
  );
}

export default MainPage;