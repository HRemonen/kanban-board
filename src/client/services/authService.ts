import apiClient from '../util/apiClient'
import { LoginUser, APIResponse } from '../types'

const login = async ({ email, password }: LoginUser) => {
  const { data }: { data: APIResponse } = await apiClient.post('/auth/login', {
    email,
    password,
  })

  return data
}

export default login
