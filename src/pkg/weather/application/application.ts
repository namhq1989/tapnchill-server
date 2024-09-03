import { IContext } from '@/internal/context/types'
import { IFetchQuoteRequestDto } from '@/pkg/quote/dto/fetch-quote'
import {
  IWeatherApplication,
  IWeatherQueryFetchWeather,
} from '@/pkg/weather/application/types'
import { ICaching } from '@/internal/caching/types'
import { IWeather } from '@/internal/weather/types'
import WeatherQueryFetchWeather from '@/pkg/weather/application/fetch-weather'
import { ILocation } from '@/internal/location/types'

class WeatherApplication implements IWeatherApplication {
  private readonly _fetchWeatherHandler: IWeatherQueryFetchWeather

  constructor(caching: ICaching, location: ILocation, weather: IWeather) {
    this._fetchWeatherHandler = new WeatherQueryFetchWeather(
      caching,
      location,
      weather,
    )
  }

  fetchWeather(
    ctx: IContext,
    performerId: string,
    ip: string,
    req: IFetchQuoteRequestDto,
  ) {
    return this._fetchWeatherHandler.fetchWeather(ctx, performerId, ip, req)
  }
}

export default WeatherApplication
