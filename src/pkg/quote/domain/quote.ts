import { Mongo } from '@/internal/mongo'
import { Context } from '@/internal/context'

export interface IQuoteRepository {
  fetchLatest: (
    ctx: Context,
  ) => Promise<{ quote: Quote | null; error: Error | null }>
  create: (ctx: Context, quote: Quote) => Promise<Error | null>
  isDuplicate(ctx: Context, originalId: string): Promise<boolean>
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
