import "./App.css";
import { useEffect } from "react";
import {  getStatus, StatusResponse } from "./api/api";

function App() {
  useEffect(() => {
    const getData = async () => {
      const statusData: StatusResponse | Error = await getStatus();

      if (statusData instanceof Error) {
        console.log(Error);
        return;
      }

      console.log(statusData);

    //  postServo({servos:[1,2], angle: 90})
    };

    getData();
  });

  return <>Hi</>;
}

export default App;
