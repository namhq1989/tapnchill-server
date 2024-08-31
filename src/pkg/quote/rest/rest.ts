import { IRest } from '@/internal/rest/types'
import { IQuoteApplication } from '@/pkg/quote/application/types'
import { Request, Response, Router } from 'express'
import { IQuoteRest } from '@/pkg/quote/rest/types'

class QuoteRest implements IQuoteRest {
  private readonly _rest: IRest
  private readonly _quoteApplication: IQuoteApplication

  constructor(rest: IRest, quoteApplication: IQuoteApplication) {
    this._rest = rest
    this._quoteApplication = quoteApplication
  }

  start(): void {
    const router = Router()
    this._rest.server().use('/api/quote', router)

    router.get('/fetch', async (req: Request, res: Response) => {
      try {
        const { response, error } = await this._quoteApplication.fetchQuote(
          req.context,
          '',
          {},
        )
        if (error) {
          return res.status(400).json({ error: (error as Error).message })
        }

        return res.status(200).json(response)
      } catch (error) {
        console.error('Failed to fetch quote:', error)
        res.status(500).json({ error: (error as Error).message })
      }
    })
  }
}

export default QuoteRest
