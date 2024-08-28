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
  fetchOne: () => Promise<IQuoteDocument | null>
}
