import { useStatus } from "../../hooks/useStatus";
import Badge from "react-bootstrap/Badge";
import "./toolDetail.css"

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
    <div>
      <Badge bg="secondary" > {toolId}</Badge>
      <span className="tool-length">{getToolLength()}</span>
    </div>
  );
};
