import { IRealtime } from '@/internal/realtime/types'
import { Server, Socket } from 'socket.io'
import { Server as HttpServer } from 'http'

class Realtime implements IRealtime {
  private readonly _server: Server | null = null

  constructor(httpServer: HttpServer) {
    this._server = new Server(httpServer, {})
    this._events()
  }

  private _events() {
    if (!this._server) {
      return
    }

    this._server.on('connection', (socket: Socket) => {
      console.log('[realtime] new socket connection', socket.id)

      socket.on('disconnect', () => {
        console.log('[realtime] socket disconnect', socket.id)
      })
    })
  }

  broadcastAll(event: string, data: object) {
    if (!this._server) {
      return
    }

    this._server.emit(event, data)
  }
}

export default Realtime
