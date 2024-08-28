import { randomUUID as cryptoRandom } from 'crypto'

const randomUUID = (): string => {
  return cryptoRandom()
}

export { randomUUID }
