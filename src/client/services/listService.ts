import { useMutation } from 'react-query'

import apiClient from '../util/apiClient'
import queryClient from '../util/queryClient'

import { useAuthenticatedUser } from '../contexts/AuthContext'

import { NewList } from '../types'

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
