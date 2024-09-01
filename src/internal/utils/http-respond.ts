import { Response } from 'express'

const sendResponse = (
  res: Response,
  status: number,
  error: Error | null,
  data: object,
): void => {
  res.status(status).json({
    data: data || {},
    error: error ? error.message : '',
  })
}

const r200 = (res: Response, data: object): void => {
  sendResponse(res, 200, null, data)
}

const r400 = (res: Response, data: object, error: Error): void => {
  sendResponse(res, 400, error, data)
}

const r401 = (res: Response, data: object, error: Error): void => {
  sendResponse(res, 401, error, data)
}

const r404 = (res: Response, data: object, error: Error): void => {
  sendResponse(res, 404, error, data)
}

export default {
  r200,
  r400,
  r401,
  r404,
}
