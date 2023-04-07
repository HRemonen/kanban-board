import apiClient from '../util/apiClient'
import { LoginUser, LoginUserSuccess } from '../types'

interface Config {
  headers: {
    Authorization: string
  }
}

export const config: Config = { headers: { Authorization: '' } }

const setToken = (newToken: string) => {
  config.headers.Authorization = `bearer ${newToken}`
}

const loginService = async ({ email, password }: LoginUser) => {
  const { data }: { data: LoginUserSuccess } = await apiClient.post(
    '/auth/login',
    {
      email,
      password,
    }
  )

  setToken(data.data.token)

  return data
}

export default loginService
