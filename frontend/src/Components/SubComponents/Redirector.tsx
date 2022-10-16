import React from "react";
import { Navigate } from "react-router-dom";
import MainPage from "../MainPage";

type Props = {
  eventVisState: React.ComponentState;
  setEventVisState: React.Dispatch<React.SetStateAction<boolean>>;
};


function Redirector(props: Props) {

  return (
    <div>
      {localStorage.getItem("token") ? (
        <MainPage eventVisState={props.eventVisState} setEventVisState={props.setEventVisState} />
      ) : (
        <Navigate to={{ pathname: "/login" }} />
      )}
    </div>
  );
}

export default Redirector;
