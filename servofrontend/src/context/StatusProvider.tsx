import { createContext, useState } from "react";
import { StatusResponse } from "../api/api";

export type StatusContextType = {
  status: StatusResponse | null;
  setStatus: (status: StatusResponse) => void;
};

const StatusContext = createContext<StatusContextType | null>(null);

export const StatusProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [status, setStatus] = useState<StatusResponse | null >(null);

  return (
    <StatusContext.Provider value={{ status, setStatus }}>
      {children}
    </StatusContext.Provider>
  );
};

export default StatusContext;
