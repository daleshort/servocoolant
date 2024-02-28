import { useStatus } from "../../hooks/useStatus";

export const AutoSettings = (): React.JSX.Element => {
  const { status } = useStatus();

  return (
    <div>
      <div>Program running: {status?.isprogramrunning?"true":"false"}</div>
      <div>Queue Position: {status?.currenttoolqueueposition}</div>
      <div> Queue: {status?.toolqueue}</div>
    </div>
  );
};
