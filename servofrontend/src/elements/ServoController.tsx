import Card from "react-bootstrap/Card";
import { ListGroup } from "react-bootstrap";
import { ServoNumber } from "../api/api";
import { ServoDial } from "./servo/ServoDial";
import { ServoAutoManualCheckbox } from "./servo/ServoAutoManualCheckbox";
import { ServoWiggleCheckbox } from "./servo/ServoWiggleCheckbox";

export type ServoControllerProps = {
  servoId: ServoNumber;
};

export const ServoController = ({
  servoId,
}: ServoControllerProps): React.JSX.Element => {
  return (
    <Card>
      <Card.Header as="h5">Servo {servoId}</Card.Header>
      <Card.Body>
        <ServoDial servoId={servoId} />
      </Card.Body>
      <ListGroup variant="flush">
        <ListGroup.Item>
          <ServoAutoManualCheckbox servoId={servoId} />
        </ListGroup.Item>
        <ListGroup.Item>
          <ServoWiggleCheckbox servoId={servoId} />
        </ListGroup.Item>
    
      </ListGroup>
    </Card>
  );
};

//https://pavelkukov.github.io/react-dial-knob/?path=/story/knob-skins--donut&knob-Diameter=120&knob-Min=0&knob-Max=270&knob-Step=1&knob-Jump Limit=1&knob-Value=137&knob-SpaceMaxFromZero=&knob-Thickness=20&knob-Color=rgba(255,255,255,1)&knob-Background=#e1e1e1&knob-Background (Max reached)=rgba(255,180,42,1)&knob-Center Color=rgba(190,112,112,1)&knob-Focused Center Color=rgba(237,236,233,1)
