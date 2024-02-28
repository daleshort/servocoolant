import { Button } from "react-bootstrap";
import { useStatus } from "../../hooks/useStatus";
import { postToolQueueToPosision } from "../../api/api";
import "./toolqueue.css";

export const ToolQueue = (): React.JSX.Element => {
  const { status } = useStatus();

  const getToolQueueList = () => {
    return status?.toolqueue ? status.toolqueue : [];
  };

  const getBadgeColor =(tool:number)=>{
    
    if(status){

        return status.currenttoolqueueposition === tool? "primary": "secondary"
    }
    return "secondary"
  }

  const handleButtonClick = (position:number)=>{

    postToolQueueToPosision({
        toolid: position
    })
  }

  return (
    <div>
      {getToolQueueList().map((t,index) => {
        return (
          
            <Button variant={getBadgeColor(index)} className="toolqueue-badge" onClick={()=>handleButtonClick(index)}>
              <span className="toolqueue-item-text">{t}</span>
            </Button>
          
        );
      })}
    </div>
  );
};
