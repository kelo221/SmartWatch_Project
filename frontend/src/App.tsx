import React from 'react'
import { Grommet, Page, PageContent } from 'grommet'

import { hpe } from 'grommet-theme-hpe'
import MainHeader from './Components/Header'
import RegisterForm from './Components/RegisterForm'
import { Route, Routes } from 'react-router-dom'
import LoginForm from './Components/LoginForm'
import Redirector from './Components/Redirector'

function App() {
    return (
        <Grommet className="App" theme={hpe}>
            <Page>
                <MainHeader></MainHeader>
                <PageContent>
                    <Routes>
                        <Route path="/" element={<Redirector />} />
                        <Route path="/login" element={<LoginForm />} />
                        <Route path="/register" element={<RegisterForm />} />
                    </Routes>
                </PageContent>
            </Page>
        </Grommet>
    )
}

export default App
