import { defineQuery } from "@pinia/colada";
import { getAll } from "../services/subprocess";

export const subprocessesQueryKey = "subprocesses"

export const useSubprocesses = defineQuery({
    key: [subprocessesQueryKey],
    query: () => getAll(),
    placeholderData: prev => prev,
})