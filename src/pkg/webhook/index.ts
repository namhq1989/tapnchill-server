import { IApp, IModule } from '@/app'
import WebhookRest from '@/pkg/webhook/rest/rest'

class WebhookModule implements IModule {
  name = () => 'Webhook'
  start = async (app: IApp) => {
    // rest
    const webhookRest = new WebhookRest(app.getRest(), app.getRealtime())
    webhookRest.start()

    console.log('🚀 [pkg webhook] started')
    return null
  }
}

export default WebhookModule
