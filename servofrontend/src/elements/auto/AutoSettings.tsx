import { useStatus } from "../../hooks/useStatus";
import { Button } from "react-bootstrap";
import { postToolToQueue, getProgramEnd, getProgramStart } from "../../api/api";
import { useState } from "react";

export const AutoSettings = (): React.JSX.Element => {
  const { status } = useStatus();
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
    <div>
      <div>Program running: {status?.isprogramrunning ? "true" : "false"}</div>
      <div>Queue Position: {status?.currenttoolqueueposition}</div>
      <div> Queue: {status?.toolqueue}</div>
      <Button onClick={handleStartClick}> Force Start </Button>
      <Button onClick={handleEndClick}> Force End</Button>
      <Button onClick={handleAddClick}>Add {count} to Queue </Button>
    </div>
  );
};
