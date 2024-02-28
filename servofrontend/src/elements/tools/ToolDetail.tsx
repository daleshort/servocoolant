import { useStatus } from "../../hooks/useStatus";
import Button from "react-bootstrap/Button";
import "./toolDetail.css";
import ProgressBar from "react-bootstrap/ProgressBar";
import { postForceTool } from "../../api/api";

export type ToolDetailProps = {
  toolId: number;
};
export const ToolDetail = ({ toolId }: ToolDetailProps): React.JSX.Element => {
  const { status } = useStatus();

  const getToolLength = () => {
    if (status?.tools[toolId]) {
      return status.tools[toolId].length;
    }
    return 0;
  };

  const isToolActive = () => {
    if (status) {
      if (
        status.currenttoolqueueposition < 0 ||
        status.currenttoolqueueposition >= status.toolqueue.length
      ) {
        return false;
      }
      return status.toolqueue[status.currenttoolqueueposition] == toolId;
    }
    return false;
  };

  const getBadgeVariant = () => {
    if (isToolActive()) {
      return "primary";
    }

    return "secondary";
  };

  const getBarVariant = () => {
    return getToolLength() < 0 ? "danger" : "primary";
  };

  const getStripedAndAnimated = () => {
    if (isToolActive()) {
      return true;
    }
    return false;
  };

  const handleButtonClick = () => {
    postForceTool({ toolid: toolId });
  };

  return (
    <div className="tool-detail-container">
      <Button
        className="tool-badge "
        variant={getBadgeVariant()}
        size="sm"
        onClick={handleButtonClick}
      >
        {toolId}
      </Button>

      <span className="tool-length  tool-badge-vertical-center">
        {getToolLength()}{" "}
      </span>
      <div className="progress-bar">
        <ProgressBar
          now={Math.abs(getToolLength())}
          variant={getBarVariant()}
          min={0}
          max={10}
          striped={getStripedAndAnimated()}
          animated={getStripedAndAnimated()}
        />
      </div>
    </div>
  );
};
