import { IRest } from '@/internal/rest/types'
import { IQuoteApplication } from '@/pkg/quote/application/types'
import { Request, Response, Router } from 'express'
import { IQuoteRest } from '@/pkg/quote/rest/types'
import httpRespond from '@/internal/utils/http-respond'

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
          return httpRespond.r400(res, {}, error)
        }

        return httpRespond.r200(res, response!)
      } catch (error) {
        console.error('failed to fetch quote:', error)
        return httpRespond.r400(res, {}, error as Error)
      }
    })
  }
}

export default QuoteRest
