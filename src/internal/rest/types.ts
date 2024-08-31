import { Express } from 'express'

export interface IRest {
  server: () => Express
}
