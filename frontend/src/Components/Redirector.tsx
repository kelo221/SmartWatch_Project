import React from "react";
import { Navigate } from "react-router-dom";
import MainPage from "./MainPage";

function Redirector() {
    return (
      <div>
          {localStorage.getItem("token") ? (
            <MainPage />
          ) : (
            <Navigate to={{ pathname: "/login" }} />
          )}
      </div>
    );
}

export default Redirector;
