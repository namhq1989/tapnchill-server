import { IContext } from '@/internal/context/types'
import Rest from '@/rest'
import { Config } from '@/internal/config'

declare global {
  namespace Express {
    interface Request {
      context: IContext
    }
  }
}

const start = async () => {
  const config = new Config()
  const rest = new Rest()

  rest.server().listen(config.restPort(), () => {
    console.log(`🚀 [server] running at http://localhost:${config.restPort()}`)
  })
}

start().then()
