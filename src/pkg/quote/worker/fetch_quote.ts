import { IQuoteRepository } from '@/pkg/quote/repository/types'
import { IQuoteWorkerFetchQuote } from '@/pkg/quote/worker/types'
import { IContext } from '@/internal/context/types'

class QuoteWorkerFetchQuote implements IQuoteWorkerFetchQuote {
  private readonly _quoteRepository: IQuoteRepository

  constructor(quoteRepository: IQuoteRepository) {
    this._quoteRepository = quoteRepository
  }

  async fetchQuote(ctx: IContext) {
    ctx.logger().info('[worker] new fetch quote job')

    let quote
    let error
    let isDuplicated
    let attempts = 0

    do {
      if (attempts > 3) {
        error = new Error('3 attempts exceeded')
        ctx.logger().error('failed to fetch quote from api', error)
        throw error
      }

      ctx.logger().info('fetching quote from api')
      ;({ quote, error } = await this._quoteRepository.quotableRandom(ctx))

      if (error) {
        ctx.logger().error('failed to fetch quote from api', error)
        throw error
      }

      if (!quote) {
        ctx.logger().warn('no quote found in api')
        throw new Error('no quote found in api')
      }

      isDuplicated = await this._quoteRepository.isDuplicate(
        ctx,
        quote.originalId,
      )
      if (isDuplicated) {
        ctx.logger().info('quote already exists in db, fetching another one...')
        attempts++
        await new Promise((resolve) => setTimeout(resolve, 1000))
      }
    } while (isDuplicated)

    ctx.logger().info('unique quote found, persisting in db')
    error = await this._quoteRepository.create(ctx, quote)
    if (error) {
      ctx.logger().error('failed to persist quote in db', error)
      throw error
    }

    ctx.logger().info('new quote persisted in db', quote)
  }
}

export default QuoteWorkerFetchQuote
