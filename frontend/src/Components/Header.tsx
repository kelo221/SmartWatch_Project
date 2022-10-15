import React from "react";
import { Button, Clock, Header, Nav, Text } from "grommet";

type Props = {
  setState: React.Dispatch<React.SetStateAction<boolean>>;
};

const MainHeader = (props: Props) => {


  return (
    <Header background="brand" pad="small">
      <Nav direction="row">
        <Text
          weight={"bold"} size={"large"}>Smartwatch Project</Text>
      </Nav>
      <Nav direction="row">
        <Button primary label="Show Events" onClick={() => props.setState(true)} />
      </Nav>
      <Nav direction="row">
        <Button primary label="Create A New Event" onClick={() => props.setState(false)} />
      </Nav>
      <Nav direction="row">
        <Clock type={"digital"} />
      </Nav>
    </Header>
  );
};


export default MainHeader;
