export interface CopiesDetector {
  id: string,
  detectedCopies: DetectedCopy[]
}

export interface DetectedCopy {
  name: string,
  host: string,
  port: number,
  lastRefresh: Date,
}