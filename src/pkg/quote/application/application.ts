import {
  IQuoteApplication,
  IQuoteQueryFetchQuote,
} from '@/pkg/quote/application/types'
import { IQuoteRepository } from '@/pkg/quote/repository/types'
import { IContext } from '@/internal/context/types'
import { IFetchQuoteRequestDto } from '@/pkg/quote/dto/fetch-quote'
import QuoteQueryFetchQuote from '@/pkg/quote/application/fetch-quote'

class QuoteApplication implements IQuoteApplication {
  private readonly _fetchQuoteHandler: IQuoteQueryFetchQuote

  constructor(quoteRepository: IQuoteRepository) {
    this._fetchQuoteHandler = new QuoteQueryFetchQuote(quoteRepository)
  }

  fetchQuote(ctx: IContext, performerId: string, req: IFetchQuoteRequestDto) {
    return this._fetchQuoteHandler.fetchQuote(ctx, performerId, req)
  }
}

export default QuoteApplication
