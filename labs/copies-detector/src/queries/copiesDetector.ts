import { defineQuery } from "@pinia/colada"
import { useGetCopiesDetectorsIdsService } from "../services/copiesDetector"
import { useGetCopiesDetectorService } from "../services/copiesDetector"

export const copiesDetectorsQueryKey = "copies_detectors"

export const useGetCopiesDetectorsIds = (httpPort: number) => {
  const getCopiesDetectorIds = useGetCopiesDetectorsIdsService(httpPort)

  return defineQuery({
    key: [copiesDetectorsQueryKey, httpPort],
    query: () => getCopiesDetectorIds.run(),
    placeholderData: prev => prev,
  })()
}

// export const useGetCopiesDetector = (httpPort: number, id: string) => {
//   const getCopiesDetector = useGetCopiesDetectorService(httpPort)

//   return defineQuery({
//     key: [copiesDetectorsQueryKey, id],
//     query: () => getCopiesDetector.run(id),
//     placeholderData: prev => prev,
//   })()
// }