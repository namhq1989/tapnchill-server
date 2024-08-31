import { IQuoteApplication } from '@/pkg/quote/application/types'
import { IQuoteRepository } from '@/pkg/quote/domain/quote'
import {
  IFetchQuoteRequestDto,
  IFetchQuoteResponseDto,
} from '@/pkg/quote/dto/fetch-quote'
import { convertQuoteFromDomainToDto } from '@/pkg/quote/dto/quote'
import { IContext } from '@/internal/context/types'

class QuoteApplication implements IQuoteApplication {
  private readonly _quoteRepository: IQuoteRepository

  constructor(quoteRepository: IQuoteRepository) {
    this._quoteRepository = quoteRepository
  }

  async fetchQuote(
    ctx: IContext,
    performerId: string,
    _: IFetchQuoteRequestDto,
  ): Promise<{ response: IFetchQuoteResponseDto | null; error: Error | null }> {
    ctx.logger().info('[api] new fetch quote request', { performerId })

    ctx.logger().info('find quote in db')
    const { quote, error } = await this._quoteRepository.fetchLatest(ctx)
    if (error) {
      ctx.logger().error('failed to fetch quote in db', error)
      return { response: null, error }
    }
    if (!quote) {
      ctx.logger().info('no quote found in db')
      return { response: null, error: Error('no quote found') }
    }

    ctx.logger().info('convert quote and respond')
    return {
      response: {
        quote: convertQuoteFromDomainToDto(quote),
      },
      error: null,
    }
  }
}

export default QuoteApplication
