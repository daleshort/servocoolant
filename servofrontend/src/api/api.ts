import { AxiosError, AxiosResponse } from "axios";
import { axiosPublic } from "./axios";

export type ServoResponse = {
  servo1angle: number;
  servo2angle: number;
};

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
