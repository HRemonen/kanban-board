import apiClient from '../util/apiClient'
import { LoginUser, LoginUserSuccess } from '../types'

const setToken = (newToken: string) => sessionStorage.setItem('token', newToken)

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
