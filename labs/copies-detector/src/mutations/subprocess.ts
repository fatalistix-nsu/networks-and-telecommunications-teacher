import { defineMutation, useQueryCache } from "@pinia/colada";
import { end, kill, run } from "../services/subprocess";
import { subprocessesQueryKey } from "../queries/subprocess";

export const useRun = defineMutation({
  mutation: () => run(),
  onSettled: async () => {
    const queryCache = useQueryCache();
    await queryCache.invalidateQueries({
      key: [subprocessesQueryKey],
      exact: true,
    });
  },
});

export const useEnd = defineMutation({
  mutation: (pid: number) => end(pid),
  onSettled: async () => {
    const queryCache = useQueryCache();
    await queryCache.invalidateQueries({
      key: [subprocessesQueryKey],
      exact: true,
    });
  },
});

export const useKill = defineMutation({
  mutation: (pid: number) => kill(pid),
  onSettled: async () => {
    const queryCache = useQueryCache();
    await queryCache.invalidateQueries({
      key: [subprocessesQueryKey],
      exact: true,
    });
  },
});
