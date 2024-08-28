import express, { Express, NextFunction, Request, Response } from 'express'
import cors from 'cors'
import bodyParser from 'body-parser'
import compression from 'compression'
import morgan from 'morgan'
import { Context } from '@/context'
import { IContext } from '@/context/types'
import rateLimit from 'express-rate-limit'

declare global {
  namespace Express {
    interface Request {
      context: IContext
    }
  }
}

const app: Express = express()
const port = process.env.PORT || 3000

const corsOptions = {
  origin: ['http://localhost:5173'],
  methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
  credentials: true,
  optionsSuccessStatus: 204,
}
app.use(cors(corsOptions))

const rateLimitOptions = rateLimit({
  windowMs: 15 * 60 * 1000, // 15 minutes
  limit: 100, // Limit each IP to 100 requests per `window`
  standardHeaders: 'draft-7',
  legacyHeaders: false,
})
app.use(rateLimitOptions)

app.use(bodyParser.json())
app.use(compression())
if (process.env.ENV !== 'release') {
  app.use(morgan('dev'))
}

app.use((req: Request, _: Response, next: NextFunction) => {
  req.context = new Context()
  next()
})

app.get('/', (req: Request, res: Response) => {
  req.context.logger().info('Hello World!', { a: 1, b: 2 })
  res.send('Express + TypeScript Server')
})

app.listen(port, () => {
  console.log(`🚀 server is running at http://localhost:${port}`)
})
