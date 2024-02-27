import { useStatus } from "../hooks/useStatus";
import { useEffect, useState } from "react";
import { StatusResponse } from "../api/api";
import { getStatus } from "../api/api";

export const StatusLoader = () => {
  const { setStatus } = useStatus();
  const [pollCount, setPollCount] = useState(0);

  const getStatusFromServer = async () => {
    const statusData: StatusResponse | Error = await getStatus();

    if (statusData instanceof Error) {
      console.log(Error);
      return;
    }

    if (pollCount < 2) {
      console.log(statusData);
    }

    setStatus(statusData);
    setPollCount((prev) => prev + 1);
  };

  let interval: number;

  const startInterval = () => {
    interval = setInterval(getStatusFromServer, 1000);
  };

  useEffect(() => {
    startInterval();
    return () => {
      clearInterval(interval);
    };
  }, [pollCount]);

  return null;
};
