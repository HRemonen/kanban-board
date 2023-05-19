import apiClient from '../util/apiClient'

import { useAuthenticatedUser } from '../contexts/AuthContext'

import { APIResponse } from '../types'

export const useUsers = async () => {
  const { data }: { data: APIResponse } = await apiClient.get('/user')

  return data
}

export const useUserByID = async (id: string) => {
  const { config } = useAuthenticatedUser()
  const { data }: { data: APIResponse } = await apiClient.get(
    `/user/${id}`,
    config
  )

  return data
}
