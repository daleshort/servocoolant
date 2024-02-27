import { useStatus } from "../../hooks/useStatus";
import Badge from "react-bootstrap/Badge";
import "./toolDetail.css";
import ProgressBar from "react-bootstrap/ProgressBar";

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
  return (
    <div className="tool-detail-container">
      <Badge bg="secondary"> {toolId}</Badge>
      <span className="tool-length">{getToolLength()} </span>
      <div className="progress-bar">
        <ProgressBar now={getToolLength()} min={0} max={10} />
      </div>
    </div>
  );
};
