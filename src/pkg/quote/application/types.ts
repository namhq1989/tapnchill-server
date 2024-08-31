import {
  IFetchQuoteRequestDto,
  IFetchQuoteResponseDto,
} from '@/pkg/quote/dto/fetch-quote'
import { IContext } from '@/internal/context/types'

export interface IQuoteApplication {
  fetchQuote: (
    ctx: IContext,
    performerId: string,
    req: IFetchQuoteRequestDto,
  ) => Promise<{ response: IFetchQuoteResponseDto | null; error: Error | null }>
}
