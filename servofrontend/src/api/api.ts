import { AxiosError, AxiosResponse } from "axios";
import { axiosPublic } from "./axios";

export type StatusResponse = {
  servostatus: ServoStatus;
  istoolsensehigh: boolean;
  tools: ToolStatus;
};

export type ServoStatus = { [key: number]: ServoDetailStatus };

export type ServoDetailStatus = {
  angle: number;
  isauto: boolean;
  iswiggle: boolean;
  amplitude: number;
  frequency: number;
  travelrange: number;
  offset: number;
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
