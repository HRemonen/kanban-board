import apiClient from '../util/apiClient'
import { LoginUser, LoginUserSuccess } from '../types'

interface Config {
  headers: {
    Authorization: string
  }
}

export const config: Config = {
  headers: { Authorization: sessionStorage.getItem('token') || '' },
}

const setToken = (newToken: string) => {
  const token = `bearer ${newToken}`

  sessionStorage.setItem('token', token)
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
