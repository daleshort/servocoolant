import Card from "react-bootstrap/Card";
import { ToolQueue } from "./auto/ToolQueue";
import { AutoSettings } from "./auto/AutoSettings";
import { ListGroup } from "react-bootstrap";
import Spinner from "react-bootstrap/Spinner";
import { useStatus } from "../hooks/useStatus";
import { SerialPort } from "./serial/SerialPort";

export const ToolQueueCard = (): React.JSX.Element => {
  const { status } = useStatus();

  const getSpinner = () => {
    if (status?.isprogramrunning) {
      return <Spinner animation="border" size="sm" />;
    }
  };



  return (
    <Card>
      <Card.Header as="h5">Tool Queue {getSpinner()}</Card.Header>
      <Card.Body>
        <ToolQueue />
      </Card.Body>
      <ListGroup variant="flush">
        <ListGroup.Item>
          <AutoSettings />
          <SerialPort />
        </ListGroup.Item>
        <ListGroup.Item>
         probe sense high?: { status?.isprobesensehigh}
        </ListGroup.Item>
      </ListGroup>
    </Card>
  );
};

//https://pavelkukov.github.io/react-dial-knob/?path=/story/knob-skins--donut&knob-Diameter=120&knob-Min=0&knob-Max=270&knob-Step=1&knob-Jump Limit=1&knob-Value=137&knob-SpaceMaxFromZero=&knob-Thickness=20&knob-Color=rgba(255,255,255,1)&knob-Background=#e1e1e1&knob-Background (Max reached)=rgba(255,180,42,1)&knob-Center Color=rgba(190,112,112,1)&knob-Focused Center Color=rgba(237,236,233,1)
