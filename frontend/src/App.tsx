import React from 'react'
import { Grommet, Page, PageContent } from 'grommet'

import { hpe } from 'grommet-theme-hpe'
import MainHeader from './Components/Header'
import RegisterForm from './Components/RegisterForm'
import { Route, Routes } from 'react-router-dom'
import LoginForm from './Components/LoginForm'

function App() {
    return (
        <Grommet className="App" theme={hpe}>
            <Page>
                <MainHeader></MainHeader>
                <PageContent>
                    <Routes>
                        <Route path="/login" element={<LoginForm />} />
                        <Route path="/register" element={<RegisterForm />} />
                    </Routes>
                </PageContent>
            </Page>
        </Grommet>
    )
}

export default App

function Home() {
    return <h2>Home</h2>
}

function About() {
    return <h2>About</h2>
}

function Users() {
    return <h2>Users</h2>
}
