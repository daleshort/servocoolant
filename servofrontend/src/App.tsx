import "./App.css";
import { useEffect } from "react";
import { getServo, ServoResponse } from "./api/api";

function App() {
  useEffect(() => {
    const getData = async () => {
      const servoData: ServoResponse | Error = await getServo();

      if (servoData instanceof Error) {
        console.log(Error);
        return;
      }

      console.log(servoData);
    };

    getData();
  });

  return <>Hi</>;
}

export default App;
