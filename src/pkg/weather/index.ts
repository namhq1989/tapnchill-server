import { IApp, IModule } from '@/app'
import WeatherApplication from '@/pkg/weather/application/application'
import WeatherRest from '@/pkg/weather/rest/rest'

class WeatherModule implements IModule {
  name = () => 'Weather'
  start = async (app: IApp) => {
    // application
    const weatherApplication = new WeatherApplication(
      app.getCaching(),
      app.getLocation(),
      app.getWeather(),
    )

    // rest
    const weatherRest = new WeatherRest(app.getRest(), weatherApplication)
    weatherRest.start()

    console.log('🚀 [pkg weather] started')
    return null
  }
}

export default WeatherModule
