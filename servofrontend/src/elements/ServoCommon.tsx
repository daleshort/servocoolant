import Card from "react-bootstrap/Card";

import { WiggleAmplitude } from "./servo/WiggleAmplitude";
import { WiggleFrequency } from "./servo/WiggleFrequency";

export const ServoCommon = (): React.JSX.Element => {
  return (
    <Card>
      <Card.Header as="h5">Wiggle Settings</Card.Header>
      <Card.Body>
        <WiggleAmplitude />
        <WiggleFrequency />
      </Card.Body>
    </Card>
  );
};

//https://pavelkukov.github.io/react-dial-knob/?path=/story/knob-skins--donut&knob-Diameter=120&knob-Min=0&knob-Max=270&knob-Step=1&knob-Jump Limit=1&knob-Value=137&knob-SpaceMaxFromZero=&knob-Thickness=20&knob-Color=rgba(255,255,255,1)&knob-Background=#e1e1e1&knob-Background (Max reached)=rgba(255,180,42,1)&knob-Center Color=rgba(190,112,112,1)&knob-Focused Center Color=rgba(237,236,233,1)
