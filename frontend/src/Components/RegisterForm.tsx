import React from "react";
import { Box, Button, Form, FormField, Header, Heading, List, MaskedInput, Text, TextInput } from "grommet";
import { Link } from "react-router-dom";
import { fetchRequestNoBearer } from "./Services/FetchRequest";
import { RegisterFormJSON } from "../DataFormats/DataFormats";

const RegisterForm = () => {
  const [formValues, setFormValues] = React.useState({
    username: "",
    password: "",
    passwordConfirm: ""
  });

  const onChange = (
    values: React.SetStateAction<{
            username: string
            password: string
            passwordConfirm: string
        }>
    ) => {
        setFormValues(values)
        console.log(values)
    }

    const onSubmit = (value: any) => {

      const registerFormPost: RegisterFormJSON = {
        username: value.username,
        password: value.password,
        password_confirm: value.passwordConfirm
      };

      fetchRequestNoBearer("http://localhost:8000/api/register", registerFormPost).then((data) => {
        if (!data.status) {
          console.log("ok!");
          window.location.replace(window.location.origin + "/login");
        } else {
          console.log("failed!!!", data.status);
        }
      });
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
                    Register
                </Heading>
                <Text>An account is needed to set up alarms</Text>
            </Header>
            <Box pad={{ horizontal: 'xxsmall' }}>
                <Form
                  validate="blur"
                  value={formValues}
                  messages={{
                        required: 'This is a required field.',
                    }}
                  onChange={(nextValue) => onChange(nextValue)}
                  onSubmit={({ value }) => onSubmit(value)}
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

                    <FormField
                        required={{ indicator: false }}
                        label="Password Confirm"
                        htmlFor="passwordConfirm"
                        name="passwordConfirm"
                        info={
                          <List border={{ color: "none" }} pad="none"></List>
                        }
                    >
                      <TextInput
                        id="passwordConfirmField"
                        name="passwordConfirm"
                        placeholder="Enter your password"
                        type="password"
                      />
                    </FormField>
                  <Box direction="row">
                    <Link to="/login">
                      <Button label="Log in" margin="small" secondary />
                    </Link>

                    <Button
                      label="Register"
                      primary
                      margin="small"
                      type="submit"
                    />
                    </Box>
                </Form>
            </Box>
        </Box>
    )
}

export default RegisterForm
