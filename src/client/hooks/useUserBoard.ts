import { useQuery } from 'react-query'

import apiClient from '../util/apiClient'
import { config } from '../services/authService'
import { APIResponse } from '../types'

const useUserBoard = (boardID: string) => {
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
