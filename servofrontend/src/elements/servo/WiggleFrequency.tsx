import { useStatus } from "../../hooks/useStatus";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import { postServoWiggle } from "../../api/api";
import { useState } from "react";
import { Button } from "react-bootstrap";
import "./servoInput.css"

export const WiggleFrequency = () => {
  const { status } = useStatus();
  const [frequency, setFrequency] = useState("");
  const [hasInteracted, setHasInteracted] = useState(false);

  //currently only going to support setting wiggle
  // amplidude to the same value for both servos
  // so going to use servo1 status as the status
  const getServoStatus = () => {
    return status?.servostatus[1];
  };

  const getFrequency = () => {
    const s = getServoStatus();
    if (s) {
      return s.frequency;
    }
    return -1;
  };

  const isPositiveFloat = (s: string) => {
    const num = parseFloat(s);

    return num >= 0;
  };

  const isInputValid = (input: string) => {
    if (!isPositiveFloat(input)) {
      return false;
    }
    if (parseFloat(input) > 3) {
      return false;
    }
    return true;
  };

  const handleFrequencyInput = (e: React.FormEvent<HTMLInputElement>) => {
    setHasInteracted(true);
    setFrequency(e.currentTarget.value);
  };

  const getFrequencyDisplayValue = () => {
    if (hasInteracted) {
      return frequency;
    } else {
      return getFrequency();
    }
  };

  const handleSubmit = () => {
    if (isInputValid(frequency)) {
      postServoWiggle({
        servos: [1, 2],
        frequency: parseFloat(frequency),
      });
    } else {
      setFrequency(getFrequency().toString());
    }
  };

  const isSubmitNeeded = () => {
    return frequency != getFrequency().toString();
  };

  const getButtonVariant = () => {
    if (!hasInteracted) {
      return "outline-secondary";
    }

    return isSubmitNeeded() ? "primary" : "outline-secondary";
  };

  return (
    <>
      <InputGroup className="mb-3">
        <InputGroup.Text id="basic-addon1">F</InputGroup.Text>
        <Form.Control
         className="servo-input"
          value={getFrequencyDisplayValue()}
          onInput={handleFrequencyInput}
          
        />
        <Button
          variant={getButtonVariant()}
          id="button-addon1"
          onClick={handleSubmit}
        >
          Set
        </Button>
      </InputGroup>
    </>
  );
};
