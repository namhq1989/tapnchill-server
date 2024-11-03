import { Mongo } from '@/internal/mongo'
import { ObjectId } from 'mongodb'

class Quote {
  readonly _id: ObjectId
  originalId: string
  content: string
  author: string
  createdAt: Date

  constructor(originalId: number, content: string, author: string) {
    this._id = Mongo.generateId()
    this.originalId = originalId.toString()
    this.content = content
    this.author = author
    this.createdAt = new Date()
  }
}

export default Quote
