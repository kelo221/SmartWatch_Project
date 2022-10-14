import React from "react";
import { Navigate } from "react-router-dom";
import ShowEventsPage from "./ShowEventsPage";
import CreateEventsPage from "./CreateEventsPage";

type Props = {
  state: React.ComponentState;
};


function Redirector(props: Props) {

  console.log(props.state);

  function mainPageRender() {
    if (props.state) {
      return (<ShowEventsPage />);
    } else {
      return (<CreateEventsPage />);
    }

  }

  return (
    <div>
      {localStorage.getItem("token") ? (
        mainPageRender()
      ) : (
        <Navigate to={{ pathname: "/login" }} />
      )}
    </div>
  );
}

export default Redirector;
