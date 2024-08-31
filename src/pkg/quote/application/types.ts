import { IContext } from '@/internal/context/types'
import {
  IFetchQuoteRequestDto,
  IFetchQuoteResponseDto,
} from '@/pkg/quote/dto/fetch-quote'

export interface IQuoteApplication extends IQuoteQueryFetchQuote {}

export interface IQuoteQueryFetchQuote {
  fetchQuote: (
    ctx: IContext,
    performerId: string,
    req: IFetchQuoteRequestDto,
  ) => Promise<{ response: IFetchQuoteResponseDto | null; error: Error | null }>
}
