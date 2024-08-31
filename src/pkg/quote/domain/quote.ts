import { Mongo } from '@/internal/mongo'
import { IContext } from '@/internal/context/types'

export interface IQuoteRepository {
  fetchLatest: (
    ctx: IContext,
  ) => Promise<{ quote: Quote | null; error: Error | null }>
  create: (ctx: IContext, quote: Quote) => Promise<Error | null>
  isDuplicate(ctx: IContext, originalId: string): Promise<boolean>
}

class Quote {
  readonly id: string
  originalId: string
  content: string
  author: string
  createdAt: Date

  constructor(originalId: string, content: string, author: string) {
    this.id = Mongo.generateId()
    this.originalId = originalId
    this.content = content
    this.author = author
    this.createdAt = new Date()
  }
}

export default Quote
