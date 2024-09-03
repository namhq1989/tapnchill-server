import { IContext } from '@/internal/context/types'
import {
  IFetchWeatherRequestDto,
  IFetchWeatherResponseDto,
} from '@/pkg/weather/dto/fetch-weather'

export interface IWeatherApplication extends IWeatherQueryFetchWeather {}

export interface IWeatherQueryFetchWeather {
  fetchWeather: (
    ctx: IContext,
    performerId: string,
    ip: string,
    req: IFetchWeatherRequestDto,
  ) => Promise<{
    response: IFetchWeatherResponseDto | null
    error: Error | null
  }>
}
