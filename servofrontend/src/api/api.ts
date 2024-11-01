import { AxiosError, AxiosResponse } from "axios";
import { axiosPublic } from "./axios";

export type StatusResponse = {
  servostatus: ServoStatus;
  istoolsensehigh: boolean;
  isprobesensehigh: boolean;
  tools: ToolStatus;
  toolqueue: Array<number>;
  isprogramrunning: boolean;
  currenttoolqueueposition: number;
};

export type ToolQueueRequest = {
  toolid: number;
};

export type ToolLengthRequest = {
  toolid:number,
  toollength:number
}

export type ServoStatus = { [key: number]: ServoDetailStatus };

export type ServoDetailStatus = {
  angle: number;
  isauto: boolean;
  iswiggle: boolean;
  amplitude: number;
  frequency: number;
  travelrange: number;
  offset: number;
  softlimitmin: number;
  softlimitmax: number;
};

export type ServoNumber = 1 | 2;

export type ServoPostRequest = {
  servos: Array<ServoNumber>;
  angle: number;
};
export type ServoAutoPostRequest = {
  servos: Array<ServoNumber>;
  isauto: boolean;
};

export type ServoWigglePostRequest = {
  servos: Array<ServoNumber>;
  iswiggle?: boolean;
  amplitude?: number;
  frequency?: number;
};

export type ToolData = {
  length: number;
};
export type ToolStatus = { [key: number]: ToolData };

export type ResponseOk = "ok";

export const getStatus = async (): Promise<StatusResponse | Error> => {
  const url = "status";
  try {
    const response: AxiosResponse = await axiosPublic.get(url);
    const responseData: StatusResponse = response.data;
    return responseData;
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};

export const getProgramStart = async (): Promise<number | Error> => {
  const url = "/auto/programstart";
  try {
    const response: AxiosResponse = await axiosPublic.get(url);
    return response.status;
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};

export const getProgramEnd = async (): Promise<number | Error> => {
  const url = "/auto/programend";
  try {
    const response: AxiosResponse = await axiosPublic.get(url);
    return response.status;
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};

export const postToolLength = async (
  request: ToolLengthRequest
): Promise<Error | ResponseOk> => {
  const url = "/toollength";

  try {
    await axiosPublic.post(url, request);
    return "ok";
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};


export const postToolToQueue = async (
  request: ToolQueueRequest
): Promise<Error | ResponseOk> => {
  const url = "auto/toolqueueadd";

  try {
    await axiosPublic.post(url, request);
    return "ok";
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};
export const postToolQueueToPosision = async (
  request: ToolQueueRequest
): Promise<Error | ResponseOk> => {
  const url = "auto/queueposition";

  try {
    await axiosPublic.post(url, request);
    return "ok";
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};

export const postForceTool = async (
  request: ToolQueueRequest
): Promise<Error | ResponseOk> => {
  const url = "auto/forcetool";

  try {
    await axiosPublic.post(url, request);
    return "ok";
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};

export const getServo = async (): Promise<ServoStatus | Error> => {
  const url = "servo";
  try {
    const response: AxiosResponse = await axiosPublic.get(url);
    const responseData: ServoStatus = response.data;
    return responseData;
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};

export const postServo = async (
  request: ServoPostRequest
): Promise<Error | ResponseOk> => {
  const url = "servo";

  try {
    await axiosPublic.post(url, request);
    return "ok";
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};

export const postServoAuto = async (
  request: ServoAutoPostRequest
): Promise<Error | ResponseOk> => {
  const url = "servoauto";

  try {
    await axiosPublic.post(url, request);
    return "ok";
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};
export const postServoWiggle = async (
  request: ServoWigglePostRequest
): Promise<Error | ResponseOk> => {
  const url = "servowiggle";

  try {
    await axiosPublic.post(url, request);
    return "ok";
  } catch (error: unknown | AxiosError) {
    if (error instanceof AxiosError && !error?.response) {
      console.error("no server response");
    } else {
      console.error("failed request", url);
    }
  }
  return Error("response error");
};
