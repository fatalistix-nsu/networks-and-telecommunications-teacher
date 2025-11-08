import { defineMutation, useQueryCache } from "@pinia/colada";
import { useCreateCopiesDetectorService } from "../services/copiesDetector";
import { copiesDetectorsQueryKey } from "../queries/copiesDetector";

export interface CreateCopiesDetectorParams {
  host: string;
  port: number;
  name: string;
}

export const useCreateCopiesDetector = (httpPort: number) => {
  const createCopiesDetector = useCreateCopiesDetectorService(httpPort);

  return defineMutation({
    mutation: (params: CreateCopiesDetectorParams) =>
      createCopiesDetector.run(params.host, params.port, params.name),
    onSettled: async () => {
      const queryCache = useQueryCache();
      await queryCache.invalidateQueries({
        key: [copiesDetectorsQueryKey, httpPort], exact: true,
      });
    },
  })()
};
