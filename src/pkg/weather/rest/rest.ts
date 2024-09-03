import { IRest } from '@/internal/rest/types'
import { Request, Response, Router } from 'express'
import httpRespond from '@/internal/utils/http-respond'
import { IWeatherRest } from '@/pkg/weather/rest/types'
import { IWeatherApplication } from '@/pkg/weather/application/types'
import requestIp from 'request-ip'

class WeatherRest implements IWeatherRest {
  private readonly _rest: IRest
  private readonly _weatherApplication: IWeatherApplication

  constructor(rest: IRest, weatherApplication: IWeatherApplication) {
    this._rest = rest
    this._weatherApplication = weatherApplication
  }

  start(): void {
    const router = Router()
    this._rest.server().use('/api/weather', router)

    router.get('/fetch', async (req: Request, res: Response) => {
      try {
        let ip = requestIp.getClientIp(req)

        // set default value if ip is 127.0.0.1
        if (['127.0.0.1', '::ffff:127.0.0.1', '::1'].includes(ip!)) {
          ip = '171.225.184.76'
        }

        const { response, error } = await this._weatherApplication.fetchWeather(
          req.context,
          '',
          ip!,
          {},
        )
        if (error) {
          return httpRespond.r400(res, {}, error)
        }

        return httpRespond.r200(res, response!)
      } catch (error) {
        console.error('failed to fetch weather:', error)
        return httpRespond.r400(res, {}, error as Error)
      }
    })
  }
}

export default WeatherRest
