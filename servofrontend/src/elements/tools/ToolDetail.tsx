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

  const getBarVariant = () => {
    return getToolLength() < 0 ? "danger" : "primary";
  };
  return (
    <div className="tool-detail-container">
      <Badge className="tool-badge " bg="secondary">
        <span className="tool-badge-text">{toolId}</span>
      </Badge>

      <span className="tool-length  tool-badge-vertical-center">
        {getToolLength()}{" "}
      </span>
      <div className="progress-bar">
        <ProgressBar
          now={Math.abs(getToolLength())}
          variant={getBarVariant()}
          min={0}
          max={10}
        />
      </div>
    </div>
  );
};
