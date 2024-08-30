import { Context } from '@/internal/context'
import {
  IFetchQuoteRequestDto,
  IFetchQuoteResponseDto,
} from '@/pkg/quote/dto/fetch-quote'

export interface IQuoteApplication {
  fetchQuote: (
    ctx: Context,
    performerId: string,
    req: IFetchQuoteRequestDto,
  ) => Promise<{ response: IFetchQuoteResponseDto | null; error: Error | null }>
}
