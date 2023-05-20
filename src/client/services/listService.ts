import { useMutation, useQuery } from 'react-query'

import apiClient from '../util/apiClient'
import queryClient from '../util/queryClient'

import { useAuthenticatedUser } from '../contexts/AuthContext'

import { ListAPIResponse, NewList } from '../types'

export const useList = (listID: string | undefined) => {
  const queryKey = ['list', listID]

  const query = async (): Promise<ListAPIResponse> => {
    const { data }: { data: ListAPIResponse } = await apiClient.get(
      `/list/${listID}`
    )
    return data
  }

  const { data: listData, ...rest } = useQuery(queryKey, query)

  return { listData, ...rest }
}

export const useCreateNewList = () => {
  const { config } = useAuthenticatedUser()

  const mutationFn = async ({
    boardID,
    list,
  }: {
    boardID: string
    list: NewList
  }) => {
    await apiClient.post(`/board/${boardID}/list`, list, config)
  }

  const mutation = useMutation(mutationFn, {
    onSuccess: (_, variables) =>
      queryClient.invalidateQueries(['board', variables.boardID]),
  })

  return mutation
}

export const useDeleteList = () => {
  const { config } = useAuthenticatedUser()

  const mutationFn = async ({
    boardID,
    listID,
  }: {
    boardID: string
    listID: string
  }) => {
    await apiClient.delete(`/board/${boardID}/list/${listID}`, config)
  }
  const mutation = useMutation(mutationFn, {
    onSuccess: (_, variables) =>
      queryClient.invalidateQueries(['board', variables.boardID]),
  })

  return mutation
}
