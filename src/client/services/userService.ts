import { APIResponse } from '../types'
import apiClient from '../util/apiClient'

export const getUsers = async () => {
  const { data }: { data: APIResponse } = await apiClient.get('/user')

  return data
}

export const getUserByID = async (id: string) => {
  const { data }: { data: APIResponse } = await apiClient.get(`/user/${id}`)

  return data
}
