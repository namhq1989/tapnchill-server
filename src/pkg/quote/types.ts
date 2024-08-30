export interface IQuotableDocument {
  _id: string
  content: string
  author: string
}

export interface IQuoteDocument {
  id: string
  rawId: string
  content: string
  author: string
  createdAt: Date
}

export interface IQuote {
  initialize: () => Promise<void>
  fetchOne: () => Promise<IQuoteDocument | null>
}
