import "./App.css";

import { StatusProvider } from "./context/StatusProvider";
import { SettingsCommon } from "./elements/SettingsCommon";
import { ServoController } from "./elements/ServoController";
import { StatusLoader } from "./elements/StatusLoader";
import { ToolList } from "./elements/tools/ToolList";
import { ToolQueueCard } from "./elements/ToolQueueCard";

function App() {
  return (
    <StatusProvider>
      <div className="grid-container">
        <div className="box servo1">
          <ServoController servoId={1} />
        </div>
        <div className="box servo2">
          <ServoController servoId={2} />
        </div>
        <div className="box status">
          <ToolList />
        </div>
        <div className="box servo-common">
          <SettingsCommon />
         
        </div>
        <div className="box toolqueue">
         
          <ToolQueueCard />
        </div>
      </div>
      <StatusLoader />
    </StatusProvider>
  );
}

export default App;
