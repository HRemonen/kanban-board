import apiClient from '../util/apiClient'
import { LoginUser, LoginUserSuccess } from '../types'

const loginService = async ({ email, password }: LoginUser) => {
  const { data }: { data: LoginUserSuccess } = await apiClient.post(
    '/auth/login',
    {
      email,
      password,
    }
  )

  return data
}

export default loginService
