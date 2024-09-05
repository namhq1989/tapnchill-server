import { IMongo } from '@/internal/mongo/types'
import { Collection } from 'mongodb'
import { IContext } from '@/internal/context/types'
import { IFeedbackRepository } from '@/pkg/feedback/repository/types'
import Feedback from '@/pkg/feedback/domain/feedback'

class FeedbackRepository implements IFeedbackRepository {
  private readonly _mongo: IMongo
  private readonly _collectionName = 'feedbacks'

  constructor(mongo: IMongo) {
    this._mongo = mongo
    ;(async () => {
      await this._ensureIndexes()
    })()
  }

  private async _ensureIndexes(): Promise<void> {
    try {
      await this._getCollection().createIndexes([
        { key: { createdAt: -1 } },
      ])
    } catch (error) {
      console.error('error creating feedbacks indexes:', error)
    }
  }

  private _getCollection(): Collection {
    return this._mongo.getDb().collection(this._collectionName)
  }

  async create(_: IContext, feedback: Feedback): Promise<Error | null> {
    const collection = this._getCollection()
    await collection.insertOne(feedback)
    return null
  }
}

export default FeedbackRepository
