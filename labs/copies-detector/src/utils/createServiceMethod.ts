export function createServiceMethod<
  TConst,
  TArgs extends any[],
  TResult,
>(
  fn: (c: TConst, ...args: TArgs) => Promise<TResult>
) {
  return (c: TConst) => ({
    run: (...args: TArgs) => fn(c, ...args)
  })
}