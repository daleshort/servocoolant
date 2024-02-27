import "./App.css";

import { StatusProvider } from "./context/StatusProvider";
import { ServoCommon } from "./elements/ServoCommon";
import { ServoController } from "./elements/ServoController";
import { StatusLoader } from "./elements/StatusLoader";
import { ToolList } from "./elements/tools/ToolList";

function App() {
  return (
    <StatusProvider>
      <div className="grid-container">
        <div className="box box1">
          <ServoController servoId={1} />
        </div>
        <div className="box box2">
          {" "}
          <ServoController servoId={2} />
        </div>
        <div className="box box3">
          <ToolList />
        </div>
        <div className="box box4">
          <ServoCommon />
        </div>
      </div>
      <StatusLoader />
    </StatusProvider>
  );
}

export default App;
