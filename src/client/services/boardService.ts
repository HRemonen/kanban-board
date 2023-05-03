import { useMutation } from 'react-query'

import { config } from './authService'

import apiClient from '../util/apiClient'
import queryClient from '../util/queryClient'

import { NewBoard } from '../types'

const useCreateBoard = () => {
  const mutationFn = async ({ board }: { board: NewBoard }) => {
    await apiClient.post(`/board`, board, config)
  }

  const mutation = useMutation(mutationFn, {
    onSuccess: () => queryClient.invalidateQueries('board'),
  })

  return mutation
}

export default useCreateBoard
