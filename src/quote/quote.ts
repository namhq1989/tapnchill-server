import { IQuotableDocument, IQuote, IQuoteDocument } from '@/quote/types'
import ky from 'ky'

class Quote implements IQuote {
  private _quoteApi = 'https://api.quotable.io/random'

  async fetchOne(): Promise<IQuoteDocument | null> {
    const doc = (await ky.get(this._quoteApi).json()) as IQuotableDocument
    return this.convertDtoToDomain(doc)
  }

  private convertDtoToDomain(dto: IQuotableDocument): IQuoteDocument {
    return {
      id: dto._id,
      rawId: dto._id,
      content: dto.content,
      author: dto.author,
      createdAt: new Date(),
    }
  }
}

export default Quote
