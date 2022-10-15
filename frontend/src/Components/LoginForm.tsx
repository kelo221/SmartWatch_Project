import React, { useState } from "react";
import { Box, Button, Form, FormField, Header, Heading, MaskedInput, Notification, TextInput } from "grommet";
import { Link } from "react-router-dom";
import { postData } from "./FetchRequest";
import { StatusType } from "grommet/components/Notification/index";

interface loginInterface {
  username: string;
  password: string;
}

const LoginForm = () => {
  const [formValues, setFormValues] = React.useState({
    username: "",
    password: ""
  });


  const [loginSuccess, setLoginSuccess] = React.useState<StatusType>("critical");
  const [alertMessage, setAlertMessage] = React.useState("");


  const [visible, setVisible] = useState(false);
  const onOpen = () => setVisible(true);
  const onClose = () => setVisible(false);

  const onChange = (
    values: React.SetStateAction<{
      username: string
      password: string
    }>
  ) => {
    setFormValues(values);
  };

    const onSubmit = (value: loginInterface) => {
        postData("http://localhost:8000/api/login", value).then((data) => {
            if (!data.status) {
              onOpen();
              setLoginSuccess("normal");
              setAlertMessage("OK! Redirecting...");
              localStorage.setItem("token", data["token"]);
              window.location.replace(window.location.origin);
            } else {
              onOpen();
              setAlertMessage(data.errorText + " (" + data.status + "). Incorrect credentials.");
              setFormValues({ password: "", username: "" });
            }
        });
    }

    return (
      <>

        <Box align="center" gap="small">
          {visible && (
            <Notification
              toast={{ position: "top" }}
              status={loginSuccess}
              message={alertMessage}
              onClose={onClose}
            />
          )}
        </Box>

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
            pad={{ horizontal: "xxsmall" }}
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
                required: "This is a required field."
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

              <Box direction="row">
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
      </>
    )
}

export default LoginForm
