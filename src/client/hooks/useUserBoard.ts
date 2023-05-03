import { useQuery } from 'react-query'

import { config } from '../services/authService'

import apiClient from '../util/apiClient'

import { APIResponse } from '../types'

const useUserBoard = (boardID: string | undefined) => {
  const queryKey = ['board', boardID]

  const query = async (): Promise<APIResponse> => {
    const { data }: { data: APIResponse } = await apiClient.get(
      `/board/${boardID}`,
      config
    )
    return data
  }

  const { data: boardData, ...rest } = useQuery(queryKey, query)

  return { boardData, ...rest }
}

export default useUserBoard
