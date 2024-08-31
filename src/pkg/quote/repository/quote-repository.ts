import Quote, { IQuoteRepository } from '@/pkg/quote/domain/quote'
import { IMongo } from '@/internal/mongo/types'
import { Collection } from 'mongodb'
import { IContext } from '@/internal/context/types'

class QuoteRepository implements IQuoteRepository {
  private readonly _mongo: IMongo
  private readonly _collectionName = 'quotes'

  constructor(mongo: IMongo) {
    this._mongo = mongo
    ;(async () => {
      await this._ensureIndexes()
    })()
  }

  private async _ensureIndexes(): Promise<void> {
    try {
      await this._getCollection().createIndexes([
        { key: { originalId: 1 }, unique: true },
        { key: { createdAt: -1 } },
      ])
      console.log('Index on rawId field created')
    } catch (error) {
      console.error('Error creating index on rawId field:', error)
    }
  }

  private _getCollection(): Collection {
    return this._mongo.getDb().collection(this._collectionName)
  }

  async fetchLatest(
    _: IContext,
  ): Promise<{ quote: Quote | null; error: Error | null }> {
    const collection = this._getCollection()
    const quote = await collection.findOne({}, { sort: { createdAt: -1 } })
    return {
      quote: quote
        ? new Quote(quote.originalId, quote.content, quote.author)
        : null,
      error: null,
    }
  }

  async create(_: IContext, quote: Quote): Promise<Error | null> {
    const collection = this._getCollection()
    await collection.insertOne(quote)
    return null
  }

  async isDuplicate(_: IContext, originalId: string): Promise<boolean> {
    const collection = this._getCollection()
    const count = await collection.countDocuments({ originalId })
    return count > 0
  }
}

export default QuoteRepository
