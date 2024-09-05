import express, { Express, NextFunction, Request, Response } from 'express'
import cors from 'cors'
import rateLimit from 'express-rate-limit'
import bodyParser from 'body-parser'
import compression from 'compression'
import morgan from 'morgan'
import { Context } from '@/internal/context'
import { IRest } from '@/internal/rest/types'
import { createServer, Server } from 'http'

class Rest implements IRest {
  private readonly _server: Express | null = null
  private readonly _http: Server | null = null

  constructor() {
    this._server = express()

    this._server.use(cors(corsOptions))
    this._server.use(rateLimitOptions)

    this._server.use(bodyParser.json())
    this._server.use(compression())
    if (process.env.ENV !== 'release') {
      this._server.use(morgan('dev'))
    }

    this._server.use((req: Request, _: Response, next: NextFunction) => {
      req.context = new Context()
      next()
    })

    this._http = createServer(this._server)
  }

  server(): Express {
    return this._server!
  }

  http(): Server {
    return this._http!
  }
}

const corsOptions = {
  origin: ['http://localhost:5173'],
  methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
  credentials: true,
  optionsSuccessStatus: 204,
}

const rateLimitOptions = rateLimit({
  windowMs: 5 * 60 * 1000, // 5 minutes
  limit: 100, // Limit each IP to 100 requests per `window`
  standardHeaders: 'draft-7',
  legacyHeaders: false,
})

export default Rest
