export interface ILocation {
  getCityByIp: (ip: string) => Promise<string>
}

export interface IIpInfoResponse {
  ip: string
  city: string // "Da Nang"
  region: string // "Da Nang"
  country: string // VN
  loc: string // "16.0678,108.2208"
  timezone: string // "Asia/Ho_Chi_Minh"
}
