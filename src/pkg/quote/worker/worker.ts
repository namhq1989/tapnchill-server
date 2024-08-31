import { IQuoteWorker } from '@/pkg/quote/worker/types'
import { IQuoteRepository } from '@/pkg/quote/repository/types'
import QuoteWorkerFetchQuote from '@/pkg/quote/worker/fetch_quote'
import { Context } from '@/internal/context'
import { IQueue } from '@/internal/queue/types'

class QuoteWorker implements IQuoteWorker {
  private readonly _queue: IQueue | null = null
  private readonly _queueName = 'fetchQuote'
  private readonly _quoteRepository: IQuoteRepository

  private readonly _jobNames = {
    fetchQuote: 'fetchQuote',
  }

  constructor(queue: IQueue, quoteRepository: IQuoteRepository) {
    this._queue = queue
    this._quoteRepository = quoteRepository
  }

  async start(): Promise<void> {
    await this._cronjob()
    await this._workers()
  }

  private async _cronjob(): Promise<void> {
    const ctx = new Context()

    await this._queue!.scheduleJob(
      ctx,
      this._queueName,
      this._jobNames.fetchQuote,
      '0 0 */3 * * *',
    )
  }

  private async _workers(): Promise<void> {
    const ctx = new Context()
    this._queue!.processJob(
      ctx,
      this._queueName,
      this._jobNames.fetchQuote,
      {},
      async () => {
        const w = new QuoteWorkerFetchQuote(this._quoteRepository)
        await w.fetchQuote(ctx)
      },
    )
  }
}

export default QuoteWorker
