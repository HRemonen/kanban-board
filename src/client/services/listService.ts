import { useMutation } from 'react-query'

import apiClient from '../util/apiClient'

import { NewList } from '../types'
import queryClient from '../util/queryClient'
import { config } from './authService'

export const useCreateNewList = () => {
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
    onSuccess: () => queryClient.invalidateQueries('board'),
  })

  return mutation
}

export const useDeleteList = () => {
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
    onSuccess: () => queryClient.invalidateQueries('board'),
  })

  return mutation
}
