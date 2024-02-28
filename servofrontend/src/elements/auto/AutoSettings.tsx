import { Button } from "react-bootstrap";
import { postToolToQueue, getProgramEnd, getProgramStart } from "../../api/api";
import { useState } from "react";
import "./autosettings.css"

export const AutoSettings = (): React.JSX.Element => {
  
  const [count, setCount] = useState(1);

  const handleAddClick = () => {
    postToolToQueue({ toolid: count });
    setCount((old: number) => old + 1);
  };

  const handleStartClick = () => {
    getProgramStart();
  };

  const handleEndClick = () => {
    getProgramEnd();
  };

  return (
    <div >
      <Button className="button-in-group"  variant="secondary" onClick={handleStartClick}>
        Force Start
      </Button>
      <Button className="button-in-group" variant="danger" onClick={handleEndClick}>
        Force End
      </Button>
      <Button className="button-in-group" onClick={handleAddClick}>Add {count} to Queue </Button>
    </div>
  );
};
