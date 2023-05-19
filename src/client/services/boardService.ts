import { useMutation, useQuery } from 'react-query'

import { config } from './authService'

import apiClient from '../util/apiClient'
import queryClient from '../util/queryClient'

import { APIResponse, NewBoard } from '../types'

export const useUserBoard = (boardID: string | undefined) => {
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

export const useCreateBoard = () => {
  const mutationFn = async ({ board }: { board: NewBoard }) => {
    await apiClient.post(`/board`, board, config)
  }

  const mutation = useMutation(mutationFn, {
    onSuccess: () => queryClient.invalidateQueries('board'),
  })

  return mutation
}
