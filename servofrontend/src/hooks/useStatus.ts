import { useContext } from "react";
import StatusContext, { StatusContextType } from "../context/StatusProvider";

export const useStatus = () =>{
    return useContext(StatusContext) as StatusContextType
}