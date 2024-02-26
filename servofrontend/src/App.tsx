import "./App.css";
import { useEffect } from "react";
import { getServo, ServoResponse, postServo } from "./api/api";

function App() {
  useEffect(() => {
    const getData = async () => {
      const servoData: ServoResponse | Error = await getServo();

      if (servoData instanceof Error) {
        console.log(Error);
        return;
      }

      console.log(servoData);

      postServo({servos:[1,2], angle: 90})
    };

    getData();
  });

  return <>Hi</>;
}

export default App;
