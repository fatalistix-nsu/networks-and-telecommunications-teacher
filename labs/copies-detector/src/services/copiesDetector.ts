import axios from "axios";
import { baseUrl } from "./constant";
import { createServiceMethod } from "../utils/createServiceMethod";
import { CopiesDetector } from "../model/copiesDetector";

export const uri = "/copies_detectors";

const getCopiesDetectorsIds = async (httpPort: number): Promise<string[]> => {
  interface GetResponse {
    copies_detectors_ids: string[];
  }

  const resp = await axios.get<GetResponse>(baseUrl + httpPort + uri);
  return resp.data.copies_detectors_ids;
};

export const useGetCopiesDetectorsIdsService = (httpPort: number) => {
  return createServiceMethod(getCopiesDetectorsIds)(httpPort);
};

const createCopiesDetector = async (
  httpPort: number,
  host: string,
  port: number,
  name: string
) => {
  interface Request {
    host: string;
    port: number;
    name: string;
  }

  interface Response {
    id: string;
  }

  const req: Request = {
    host: host,
    port: port,
    name: name,
  };

  const resp = await axios.post<Response>(baseUrl + httpPort + uri, req);
  return resp.data.id;
};

export const useCreateCopiesDetectorService = (httpPort: number) => {
  return createServiceMethod(createCopiesDetector)(httpPort);
};

const getCopiesDetector = async (httpPort: number, id: string): Promise<CopiesDetector> => {
  interface Response {
    id: string,
    active_copies: {
      active_copy: string,
      last_refresh: string,
    }[],
  }

  const resp = await axios.get<Response>(baseUrl + httpPort + uri + "/" + id)
  return {
    id: resp.data.id,
    detectedCopies: resp.data.active_copies.map(v => ({
      name: v.active_copy,
      host: "",
      port: 0,
      lastRefresh: new Date(v.last_refresh),
    }))
  }
}

export const useGetCopiesDetectorService = (httpPort: number) => {
  return createServiceMethod(getCopiesDetector)(httpPort)
}