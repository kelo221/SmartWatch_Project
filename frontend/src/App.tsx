import React, { useState } from "react";
import { Grommet, Page, PageContent } from "grommet";

import { hpe } from "grommet-theme-hpe";
import MainHeader from "./Components/SubComponents/Header";
import RegisterForm from "./Components/RegisterForm";
import { Route, Routes } from "react-router-dom";
import LoginForm from "./Components/LoginForm";
import Redirector from "./Components/SubComponents/Redirector";

function App() {

    const [showEvents, setEventPageVis] = useState(false);

    return (
      <Grommet className="App" theme={hpe}>
          <Page>
              <MainHeader setState={setEventPageVis}></MainHeader>
              <PageContent>
                  <Routes>
                      <Route path="/"
                             element={<Redirector eventVisState={showEvents} setEventVisState={setEventPageVis} />} />
                      <Route path="/login" element={<LoginForm />} />
                      <Route path="/register" element={<RegisterForm />} />
                  </Routes>
              </PageContent>
            </Page>
        </Grommet>
    )
}

export default App
