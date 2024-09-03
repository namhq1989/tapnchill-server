import { IWeatherDto } from '@/pkg/weather/dto/weather'

export interface IFetchWeatherRequestDto {}

export interface IFetchWeatherResponseDto {
  city: string
  weather: IWeatherDto
}
