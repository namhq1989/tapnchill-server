export interface IWeather {
  getByCity: (city: string) => Promise<{
    data: IVisualCrossingWeatherResponse | null
    error: Error | null
  }>
}

export interface IVisualCrossingWeatherResponse {
  currentConditions: IVisualCrossingWeatherCurrentConditions
  days: IVisualCrossingWeatherDay[]
}

interface IVisualCrossingWeatherCurrentConditions {
  datetimeEpoch: number
  temp: number
  humidity: number
  windspeed: number
  precipprob: number
  conditions: string
  icon: string
}

interface IVisualCrossingWeatherDay {
  datetimeEpoch: number
  tempmax: number
  tempmin: number
  temp: number
  humidity: number
  windspeed: number
  precipprob: number
  icon: string
  hours: IVisualCrossingWeatherHour[]
}

interface IVisualCrossingWeatherHour {
  datetimeEpoch: number
  temp: number
  humidity: number
  windspeed: number
  precipprob: number
  icon: string
}
