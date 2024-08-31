import Quote from '@/pkg/quote/domain/quote'

export interface IQuoteDto {
  id: string
  content: string
  author: string
}

const convertQuoteFromDomainToDto = (quote: Quote): IQuoteDto => {
  return {
    id: quote._id.toHexString(),
    content: quote.content,
    author: quote.author,
  }
}

export { convertQuoteFromDomainToDto }
