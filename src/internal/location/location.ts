import { IIpInfoResponse, ILocation } from '@/internal/location/types'
import axios from 'axios'

class Location implements ILocation {
  private readonly _ipInfoToken: string = ''
  private readonly _ipInfoApi = 'https://ipinfo.io/$ip?token=$token'

  constructor(token: string) {
    if (!token) {
      throw new Error('no ip info token provided')
    }

    this._ipInfoToken = token
  }

  async getCityByIp(ip: string): Promise<string> {
    const { info, error } = await this.getLocationByIp(ip)
    if (error) {
      console.log('error getting city by ip:', error.message)
      return ''
    } else if (!info) {
      console.log('no location information found for ip:', ip)
      return ''
    }

    return info.city
  }

  private async getLocationByIp(
    ip: string,
  ): Promise<{ info: IIpInfoResponse | null; error: Error | null }> {
    try {
      const url = this._ipInfoApi
        .replace('$ip', ip)
        .replace('$token', this._ipInfoToken)
      const response = await axios.get<IIpInfoResponse>(url)
      return { info: response.data, error: null }
    } catch (error) {
      return { info: null, error: error as Error }
    }
  }
}

export default Location
