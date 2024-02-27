import Card from "react-bootstrap/Card";
import { useStatus } from "../../hooks/useStatus";
import { ToolDetail } from "./ToolDetail";
import ListGroup from "react-bootstrap/ListGroup";
import { ListGroupItem } from "react-bootstrap";

export const ToolList = (): React.JSX.Element => {
  const { status } = useStatus();

  const getToolList = () => {
    const tools: Array<number> = [];

    if (status?.tools) {
      for (const t in status.tools) {
        tools.push(parseInt(t));
      }
    }
    console.log(tools);
    return tools;
  };

  const getToolElements = () => {
    return getToolList().map((t) => {
      return (
        <ListGroupItem>
          <ToolDetail toolId={t} />
        </ListGroupItem>
      );
    });
  };

  return (
    <Card>
      <Card.Header as="h5">Tools</Card.Header>
      <Card.Body>
        <ListGroup>{getToolElements()}</ListGroup>
      </Card.Body>
    </Card>
  );
};
