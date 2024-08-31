import { IQuoteDto } from '@/pkg/quote/dto/quote'

export interface IFetchQuoteRequestDto {}

export interface IFetchQuoteResponseDto {
  quote: IQuoteDto
}
