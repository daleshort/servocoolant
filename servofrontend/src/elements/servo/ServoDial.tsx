import { Donut } from "react-dial-knob";
import { ServoControllerProps } from "../ServoController";
import { useStatus } from "../../hooks/useStatus";
import { postServo } from "../../api/api";
import { useState } from "react";
import "./servoDial.css";

export const ServoDial = ({
  servoId,
}: ServoControllerProps): React.JSX.Element => {
  const [hasInteracted, setHasInteracted] = useState(false);
  const [desiredAngle, setDesiredAngle] = useState(0);

  const { status } = useStatus();
  const getServoStatus = () => {
    return status?.servostatus[servoId];
  };
  const getMinRange = () => {
    const s = getServoStatus();
    if (s) {
      return -1 * s.offset;
    }

    return 0;
  };

  const getMaxRange = () => {
    const s = getServoStatus();
    if (s) {
      return s.travelrange - s.offset;
    }

    return 100;
  };

  const getAngleFromServer = () => {
    const s = getServoStatus();
    if (s) {
      return s.angle;
    }
    return 0;
  };

  const getAngle = () => {
    if (!hasInteracted) {
      return getAngleFromServer();
    } else {
      return desiredAngle;
    }
  };

  const setAngle = (angle: number) => {
    setHasInteracted(true);
    setDesiredAngle(angle);
    postServo({ servos: [servoId], angle });
  };

  return (
    <div className="servo-dial">
      <Donut
        diameter={100}
        min={getMinRange()}
        max={getMaxRange()}
        step={1}
        value={getAngle()}
        theme={{
          centerColor: "#212529",
          donutThickness: 10,
          maxedBgrColor: "#ea868f",
          donutColor: "#fd7e14",
          centerFocusedColor: "#343a40",
        }}
        onValueChange={setAngle}
        ariaLabelledBy={"my-label"}
      >
        <label id={"my-label"}>Angle</label>
      </Donut>
    </div>
  );
};

//https://pavelkukov.github.io/react-dial-knob/?path=/story/knob-skins--donut&knob-Diameter=120&knob-Min=0&knob-Max=270&knob-Step=1&knob-Jump Limit=1&knob-Value=137&knob-SpaceMaxFromZero=&knob-Thickness=20&knob-Color=rgba(255,255,255,1)&knob-Background=#e1e1e1&knob-Background (Max reached)=rgba(255,180,42,1)&knob-Center Color=rgba(190,112,112,1)&knob-Focused Center Color=rgba(237,236,233,1)
