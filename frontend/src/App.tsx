import React from 'react'
import {
    Grommet,
    Heading,
    Box,
    Text,
    Page,
    PageContent,
    Paragraph,
} from 'grommet'
function App() {
    return (
        <Grommet className="App">
            <Page>
                <Heading>Hello World</Heading>
                <PageContent>
                    <Box>
                        <Text>Hello Grommet</Text>
                    </Box>
                </PageContent>
            </Page>
        </Grommet>
    )
}
export default App
