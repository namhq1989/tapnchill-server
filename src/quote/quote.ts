import { IQuotableDocument, IQuote, IQuoteDocument } from '@/quote/types'
import ky from 'ky'
import { Db } from 'mongodb'

class Quote implements IQuote {
  private _quoteApi = 'https://api.quotable.io/random'
  private _collectionName = 'quotes'
  private _db: Db

  constructor(db: Db) {
    this._db = db
  }

  async initialize(): Promise<void> {
    await this.ensureIndexes()
  }

  async fetchOne(): Promise<IQuoteDocument | null> {
    let doc: IQuotableDocument
    let quote: IQuoteDocument | null

    do {
      doc = (await ky.get(this._quoteApi).json()) as IQuotableDocument
      quote = this.convertDtoToDomain(doc)
    } while (await this.isDuplicate(quote.rawId))

    await this.save(quote)
    return quote
  }

  //
  // PRIVATE METHODS
  //

  private async ensureIndexes(): Promise<void> {
    try {
      await this._db.collection(this._collectionName).createIndexes([
        { key: { rawId: 1 }, unique: true },
        { key: { createdAt: -1 } },
      ])
      console.log('Index on rawId field created')
    } catch (error) {
      console.error('Error creating index on rawId field:', error)
    }
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

  private async isDuplicate(rawId: string): Promise<boolean> {
    const existingQuote = await this._db
      .collection(this._collectionName)
      .findOne({ rawId })
    return !!existingQuote
  }

  private async save(quote: IQuoteDocument): Promise<void> {
    await this._db.collection(this._collectionName).insertOne(quote)
  }
}

export default Quote
