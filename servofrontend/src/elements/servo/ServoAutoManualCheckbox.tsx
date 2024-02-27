import { ServoControllerProps } from "../ServoController";
import { useStatus } from "../../hooks/useStatus";
import { FormCheck } from "react-bootstrap";
import { postServoAuto } from "../../api/api";

export const ServoAutoManualCheckbox = ({
  servoId,
}: ServoControllerProps): React.JSX.Element => {
  const { status } = useStatus();
  const getServoStatus = () => {
    return status?.servostatus[servoId];
  };
  const getIsAuto= () => {
    const s = getServoStatus();
    if (s) {
      return s.isauto
    }

    return false;
  };

  const handleClick = ()=>{

    postServoAuto({servos:[servoId], isauto: !getIsAuto()})
  }


  return (
    <div>
      <FormCheck type={"checkbox"} checked={getIsAuto()} label={`Auto`} onChange={handleClick} />
    </div>
  );
};

//https://pavelkukov.github.io/react-dial-knob/?path=/story/knob-skins--donut&knob-Diameter=120&knob-Min=0&knob-Max=270&knob-Step=1&knob-Jump Limit=1&knob-Value=137&knob-SpaceMaxFromZero=&knob-Thickness=20&knob-Color=rgba(255,255,255,1)&knob-Background=#e1e1e1&knob-Background (Max reached)=rgba(255,180,42,1)&knob-Center Color=rgba(190,112,112,1)&knob-Focused Center Color=rgba(237,236,233,1)
