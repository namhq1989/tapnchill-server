import { IContext } from '@/internal/context/types'
import { IWeatherQueryFetchWeather } from '@/pkg/weather/application/types'
import { ICaching } from '@/internal/caching/types'
import { IWeather } from '@/internal/weather/types'
import {
  IFetchWeatherRequestDto,
  IFetchWeatherResponseDto,
} from '@/pkg/weather/dto/fetch-weather'
import { ILocation } from '@/internal/location/types'
import { toSlug } from '@/internal/utils/string'
import {
  convertWeatherFromApiToDto,
  IWeatherDto,
} from '@/pkg/weather/dto/weather'

class WeatherQueryFetchWeather implements IWeatherQueryFetchWeather {
  private readonly _caching: ICaching
  private readonly _location: ILocation
  private readonly _weather: IWeather

  constructor(caching: ICaching, location: ILocation, weather: IWeather) {
    this._caching = caching
    this._location = location
    this._weather = weather
  }

  async fetchWeather(
    ctx: IContext,
    performerId: string,
    ip: string,
    _: IFetchWeatherRequestDto,
  ): Promise<{
    response: IFetchWeatherResponseDto | null
    error: Error | null
  }> {
    ctx.logger().info('[api] new fetch weather request', { performerId, ip })

    ctx.logger().info('find city by ip')
    const city = await this._location.getCityByIp(ip)
    if (!city) {
      ctx.logger().info('no city found')
      return { response: null, error: Error('no city found') }
    }

    ctx.logger().info('city found from ip', { city })
    ctx.logger().info('find weather in caching')
    const cachingKey = this._caching.generateKey(`weather:${toSlug(city)}`)
    const weatherCachedData = await this._caching.get(cachingKey)
    if (weatherCachedData) {
      ctx.logger().info('found weather in caching, respond')
      const weatherData = JSON.parse(weatherCachedData) as IWeatherDto
      return {
        response: {
          city,
          weather: weatherData,
        },
        error: null,
      }
    }

    ctx.logger().info('found city, find weather data', { city })
    const { data: weatherData, error: weatherError } =
      await this._weather.getByCity(city)
    if (weatherError) {
      ctx.logger().error('failed to fetch weather data', weatherError)
      return { response: null, error: weatherError }
    }
    if (!weatherData) {
      ctx.logger().info('no weather data found')
      return { response: null, error: Error('no weather data found') }
    }

    ctx.logger().info('convert weather to response data')
    const responseData = convertWeatherFromApiToDto(weatherData)

    ctx.logger().info('store weather in caching')
    await this._caching.set(cachingKey, JSON.stringify(responseData), 60 * 60) // 1h

    ctx.logger().info('respond')
    return {
      response: {
        city,
        weather: convertWeatherFromApiToDto(weatherData),
      },
      error: null,
    }
  }
}

export default WeatherQueryFetchWeather
