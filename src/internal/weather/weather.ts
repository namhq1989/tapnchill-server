import {
  IVisualCrossingWeatherResponse,
  IWeather,
} from '@/internal/weather/types'
import axios from 'axios'

class Weather implements IWeather {
  private readonly _visualCrossingToken: string = ''
  private readonly _visualCrossingWeatherApi =
    'https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/$city/$fromDate/$toDate?unitGroup=metric&include=events%2Ccurrent%2Cdays%2Chours&key=$token&contentType=json'
  // city: Da Nang
  // fromDate: 2022-01-01
  // toDate: 2022-01-02

  constructor(token: string) {
    if (!token) {
      throw new Error('no visual crossing token provided')
    }

    this._visualCrossingToken = token
  }

  async getByCity(city: string): Promise<{
    data: IVisualCrossingWeatherResponse | null
    error: Error | null
  }> {
    try {
      const today = new Date()
      const fromDate = today.toISOString().split('T')[0] // Format: YYYY-MM-DD

      const nextTwoDays = new Date(today)
      nextTwoDays.setDate(today.getDate() + 2)
      const toDate = nextTwoDays.toISOString().split('T')[0] // Format: YYYY-MM-DD

      const url = this._visualCrossingWeatherApi
        .replace('$city', city)
        .replace('$fromDate', fromDate)
        .replace('$toDate', toDate)
        .replace('$token', this._visualCrossingToken)
      const response = await axios.get<IVisualCrossingWeatherResponse>(url)
      return { data: response.data, error: null }
    } catch (error) {
      return { data: null, error: error as Error }
    }
  }
}

export default Weather
