import React from "react";
import { Button, Clock, Header, Nav, Text } from "grommet";
import { BrowserView, MobileView } from "react-device-detect";
import { popUpStore } from "../../Stores/popUpStore";

type Props = {
  setState: React.Dispatch<React.SetStateAction<boolean>>;
};

const MainHeader = (props: Props) => {

  const popUpStoreObject = popUpStore();

  return (
    <>
      <BrowserView>
        <Header background="brand" pad="small">
          <Nav direction="row">
            <Text
              weight={"bold"} size={"large"}>Smartwatch Project</Text>
          </Nav>
          {localStorage.getItem("token") ? (
            <Nav direction="row">
              <Button primary label="Create A New Event" onClick={() => props.setState(true)} />
            </Nav>) : <></>
          }
          <Nav direction="row">
            <Clock type={"digital"} />
          </Nav>
        </Header>
      </BrowserView>

      <MobileView>
        <Header background="brand" pad="small">
          <Nav direction="row">
            <Text
              weight={"bold"} size={"large"}>Smartwatch Project</Text>
          </Nav>
          {localStorage.getItem("token") ? (
            <Nav direction="row">
              <Button primary label="Create A New Event" onClick={() => {
                popUpStoreObject.setAsNewEvent(true);
                props.setState(true);
              }} />
            </Nav>) : <></>
          }
        </Header>
      </MobileView>
    </>
  );
};


export default MainHeader;
