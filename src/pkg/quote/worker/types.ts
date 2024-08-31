import { IContext } from '@/internal/context/types'

export interface IQuoteWorker {
  start: () => Promise<void>
}

export interface IQuoteWorkerFetchQuote {
  fetchQuote: (ctx: IContext) => Promise<void>
}
