import { IContext } from '@/context/types'

declare global {
  namespace Express {
    interface Request {
      context: IContext
    }
  }
}
