import { useStatus } from "../../hooks/useStatus";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import { postServoWiggle } from "../../api/api";
import { useState } from "react";
import { Button } from "react-bootstrap";
import "./servoInput.css"

export const WiggleAmplitude = () => {
  const { status } = useStatus();
  const [amplitude, setAmplitude] = useState("");
  const [hasInteracted, setHasInteracted] = useState(false);

  //currently only going to support setting wiggle
  // amplidude to the same value for both servos
  // so going to use servo1 status as the status
  const getServoStatus = () => {
    return status?.servostatus[1];
  };

  const getAmplitude = () => {
    const s = getServoStatus();
    if (s) {
      return s.amplitude;
    }
    return -1;
  };


  const isPositiveInteger = (s: string) => {
    const num = parseInt(s);

    return parseInt(num.toString()) == num && num >= 0;
  };

  const isInputValid = (input: string) => {
    if (!isPositiveInteger(input)) {
      return false;
    }
    if (parseInt(input) > 40) {
      return false;
    }
    return true;
  };

  const handleAmplitudeInput = (e: React.FormEvent<HTMLInputElement>) => {
    setHasInteracted(true);
    setAmplitude(e.currentTarget.value);
  };

  const getAmplitudeDisplayValue = () => {
    if (hasInteracted) {
      return amplitude;
    } else {
      return getAmplitude();
    }
  };

  const handleSubmit = () => {
    if (isInputValid(amplitude)) {
      postServoWiggle({
        servos: [1, 2],
        amplitude: parseInt(amplitude),
      });
    } else {
      setAmplitude(getAmplitude().toString());
    }
  };

  const isSubmitNeeded = () => {
    return amplitude != getAmplitude().toString();
  };

  const getButtonVariant = () => {
    if(!hasInteracted){
        return "outline-secondary"
    }
    
    return isSubmitNeeded() ? "primary" : "outline-secondary";
  };

  return (
    <>
      <InputGroup className="mb-3">
        <InputGroup.Text id="basic-addon1">A</InputGroup.Text>
        <Form.Control
        className="servo-input"
          value={getAmplitudeDisplayValue()}
          onInput={handleAmplitudeInput}
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
