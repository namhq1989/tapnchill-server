import { Express } from 'express'
import { Server } from 'http'

export interface IRest {
  server: () => Express
  http: () => Server
}
