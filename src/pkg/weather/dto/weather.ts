import { IVisualCrossingWeatherResponse } from '@/internal/weather/types'

export interface IWeatherDto {
  current: IWeatherCurrentDto
  days: IVisualDayDto[]
}

interface IWeatherCurrentDto {
  temp: number
  feelsLike: number
  humidity: number
  windSpeed: number
  precipitationProbability: number
  icon: string
}

interface IVisualDayDto {
  date: Date
  tempMax: number
  tempMin: number
  temp: number
  humidity: number
  windSpeed: number
  precipitationProbability: number
  icon: string
  hours: IWeatherHour[]
}

interface IWeatherHour {
  date: Date
  temp: number
  humidity: number
  windSpeed: number
  precipitationProbability: number
  icon: string
}

const convertWeatherFromApiToDto = (
  data: IVisualCrossingWeatherResponse,
): IWeatherDto => {
  return {
    current: {
      temp: data.currentConditions.temp,
      feelsLike: data.currentConditions.feelslike,
      humidity: data.currentConditions.humidity,
      windSpeed: data.currentConditions.windspeed,
      precipitationProbability: data.currentConditions.precipprob,
      icon: data.currentConditions.icon,
    },
    days: data.days.map((day) => ({
      date: new Date(day.datetimeEpoch * 1000),
      tempMax: day.tempmax,
      tempMin: day.tempmin,
      temp: day.temp,
      humidity: day.humidity,
      windSpeed: day.windspeed,
      precipitationProbability: day.precipprob,
      icon: day.icon,
      hours: day.hours.map((hour) => ({
        date: new Date(hour.datetimeEpoch * 1000),
        temp: hour.temp,
        humidity: hour.humidity,
        windSpeed: hour.windspeed,
        precipitationProbability: hour.precipprob,
        icon: hour.icon,
      })),
    })),
  }
}

export { convertWeatherFromApiToDto }
