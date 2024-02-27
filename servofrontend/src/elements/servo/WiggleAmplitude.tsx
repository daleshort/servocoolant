import { useStatus } from "../../hooks/useStatus";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";

export const WiggleAmplitude = () => {
  const { status } = useStatus();

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
    return 0;
  };

  const getFrequency = () => {
    const s = getServoStatus();
    if (s) {
      return s.frequency;
    }
    return 0;
  };

  return (
    <InputGroup className="mb-3">
      <InputGroup.Text id="basic-addon1">Wiggle Amplitude</InputGroup.Text>
      <Form.Control value={getAmplitude()} />
    </InputGroup>
  );
};
