import { IContext } from '@/internal/context/types'
import Quote from '@/pkg/quote/domain/quote'

export interface IQuoteRepository {
  fetchLatest: (
    ctx: IContext,
  ) => Promise<{ quote: Quote | null; error: Error | null }>
  create: (ctx: IContext, quote: Quote) => Promise<Error | null>
  isDuplicate(ctx: IContext, originalId: string): Promise<boolean>
  quotableRandom(
    ctx: IContext,
  ): Promise<{ quote: Quote | null; error: Error | null }>
}
