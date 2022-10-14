import { Box, Button, Card, CheckBox, DateInput, Form, FormField, MaskedInput } from "grommet";
import React from "react";


const IPv4ElementExp = /^[0-1][0-9][0-9]$|^2[0-4][0-9]$|^25[0-5]$|^[0-9][0-9]$|^[0-9]$/;

const maskIp = [
  {
    length: [1, 2],
    regexp: IPv4ElementExp,
    placeholder: "xx"
  },
  { fixed: ":" },
  {
    length: [1, 2],
    regexp: IPv4ElementExp,
    placeholder: "xx"
  }
];

const ipValidation = /^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$/;

const CreateEventsPage = () => {

  const [checked, setChecked] = React.useState(false);

  return (
    <Card
      margin="small"
      gap="medium"
      width="medium"
      align="center"
      justify="center"
      fill
    >
      <Form>
        <FormField
          label="Event Time"
          name="maskedIpAddressValidation"
          htmlFor="masked-ip-address-validation"
          validate={{
            regexp: ipValidation,
            message: "Invalid IP address"
          }}
        >
          <MaskedInput
            name="maskedIpAddressValidation"
            id="masked-ip-address-validation"
            mask={maskIp}
          />
        </FormField>

        <DateInput
          format="mm/dd/yyyy"
          value={(new Date()).toISOString()}
          onChange={({ value }) => {
          }}
        />

        <CheckBox
          checked={checked}
          label="Silent?"
          onChange={(event) => setChecked(event.target.checked)}
        />

        <Box direction="row" gap="small" pad={"small"}>
          <Button type="submit" label="Send to Smartwatch" primary />
          <Button type="reset" label="Reset" />
        </Box>
      </Form>
    </Card>
  );
};

export default CreateEventsPage;