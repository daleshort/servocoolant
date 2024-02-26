import { AxiosError, AxiosResponse } from "axios";
import { axiosPublic } from "./axios";

export type ServoResponse = {
  servo1angle: number;
  servo2angle: number;
};

type ServoNumber = 1 | 2

export type ServoPostRequest = {
  servos: Array<ServoNumber>;
  angle: number;
};

export type ResponseOk = "ok";

export const getServo = async (): Promise<ServoResponse | Error> => {
  const url = "servo";
  try {
    const response: AxiosResponse = await axiosPublic.get(url);
    const responseData: ServoResponse = response.data;
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
