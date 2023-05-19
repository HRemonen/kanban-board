import { useMutation, useQuery } from 'react-query'

import { config } from './authService'

import apiClient from '../util/apiClient'
import queryClient from '../util/queryClient'

import { APIResponse, BoardAPIResponse, NewBoard } from '../types'

export const useUserBoard = (boardID: string | undefined) => {
  const queryKey = ['board', boardID]

  const query = async (): Promise<BoardAPIResponse> => {
    const { data }: { data: BoardAPIResponse } = await apiClient.get(
      `/board/${boardID}`,
      config
    )
    return data
  }

  const { data: boardData, ...rest } = useQuery(queryKey, query)

  return { boardData, ...rest }
}

export const useUserBoards = (userID: string | undefined) => {
  const queryKey = ['boards', userID]

  const query = async (): Promise<BoardAPIResponse> => {
    const { data }: { data: BoardAPIResponse } = await apiClient.get(
      `/board/${userID}/boards`,
      config
    )
    return data
  }

  const { data: userBoardsData, ...rest } = useQuery(queryKey, query, {
    enabled: Boolean(userID),
  })

  return { userBoardsData, ...rest }
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
