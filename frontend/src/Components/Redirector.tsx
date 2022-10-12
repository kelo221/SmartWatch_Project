import React from "react";
import { Navigate } from "react-router-dom";

function Redirector() {
    return (
      <div>
          {localStorage.getItem("token") ? (
            <div>Hello</div>
          ) : (
            <Navigate to={{ pathname: "/login" }} />
          )}
      </div>
    );
}

export default Redirector;
