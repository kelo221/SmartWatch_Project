import React, { useContext } from 'react'
import {
    Box,
    Button,
    Form,
    FormField,
    Header,
    Heading,
    MaskedInput,
    ResponsiveContext,
    TextInput,
} from 'grommet'
import { Link } from 'react-router-dom'

const LoginForm = () => {
    const [formValues, setFormValues] = React.useState({
        username: '',
        password: '',
    })

    const size = useContext(ResponsiveContext)

    const onChange = (
        values: React.SetStateAction<{
            username: string
            password: string
        }>
    ) => {
        setFormValues(values)
        console.log(values)
    }

    const onSubmit = (value: any, touched: any) => {
        console.log(value, touched)
    }

    return (
        <Box
            margin="small"
            gap="medium"
            width="medium"
            align="center"
            justify="center"
            fill
        >
            <Header
                direction="column"
                align="center"
                gap="xxsmall"
                pad={{ horizontal: 'xxsmall' }}
            >
                <Heading level={2} margin="none">
                    Log in
                </Heading>
            </Header>
            <Box pad={{ horizontal: 'xxsmall' }}>
                <Form
                    validate="blur"
                    value={formValues}
                    messages={{
                        required: 'This is a required field.',
                    }}
                    onChange={(nextValue) => onChange(nextValue)}
                    onSubmit={({ value, touched }) => onSubmit(value, touched)}
                    method="post"
                >
                    <FormField
                        label="Username"
                        name="username"
                        required={{ indicator: false }}
                    >
                        <MaskedInput
                            id="username"
                            name="username"
                            type="username"
                            placeholder="Enter your username"
                        />
                    </FormField>

                    <FormField
                        label="Password"
                        htmlFor="password"
                        name="password"
                        required={{ indicator: false }}
                    >
                        <TextInput
                            id="passwordField"
                            name="password"
                            placeholder="Enter your password"
                            type="password"
                        />
                    </FormField>

                    <Box
                        direction="row"
                        margin={{ top: 'medium', bottom: 'small' }}
                    >
                        <Link to="/register">
                            <Button label="Register" margin="small" secondary />
                        </Link>

                        <Button
                            margin="small"
                            label="Log in"
                            primary
                            type="submit"
                        />
                    </Box>
                </Form>
            </Box>
        </Box>
    )
}

export default LoginForm
