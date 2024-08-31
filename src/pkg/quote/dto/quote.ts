import Quote from '@/pkg/quote/domain/quote'

export interface IQuoteDto {
  id: string
  originalId: string
  content: string
  author: string
}

const convertQuoteFromDomainToDto = (quote: Quote): IQuoteDto => {
  return {
    id: quote.id,
    originalId: quote.originalId,
    content: quote.content,
    author: quote.author,
  }
}

export { convertQuoteFromDomainToDto }
