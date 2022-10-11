import React from 'react'
import { Box, Grommet, Page, PageContent } from 'grommet'

import { hpe } from 'grommet-theme-hpe'
import MainHeader from './Components/Header'
import RegisterForm from './Components/RegisterForm'

function App() {
    return (
        <Grommet className="App" theme={hpe}>
            <Page>
                <MainHeader></MainHeader>
                <PageContent>
                    <Box>
                        <RegisterForm />
                    </Box>
                </PageContent>
            </Page>
        </Grommet>
    )
}
export default App
