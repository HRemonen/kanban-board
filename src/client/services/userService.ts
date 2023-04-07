import { APIResponse } from '../types'
import apiClient from '../util/apiClient'

import { config } from './authService'

export const getUsers = async () => {
  const { data }: { data: APIResponse } = await apiClient.get('/user')

  return data
}

export const getUserByID = async (id: string) => {
  const { data }: { data: APIResponse } = await apiClient.get(
    `/user/${id}`,
    config
  )

  return data
}
