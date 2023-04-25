import { useMutation } from 'react-query'

import apiClient from '../util/apiClient'

import { APIResponse, Card, NewCard } from '../types'
import queryClient from '../util/queryClient'

export const getSingleCard = async (id: string) => {
  const { data }: { data: APIResponse } = await apiClient.get(`/card/${id}`)

  return data
}

export const useCreateNewCard = () => {
  const mutationFn = async ({
    listID,
    card,
  }: {
    listID: string
    card: NewCard
  }) => {
    await apiClient.post(`/list/${listID}/card`, card)
  }

  const mutation = useMutation(mutationFn, {
    onSuccess: () => queryClient.invalidateQueries('board'),
  })

  return mutation
}

export const useUpdateCardPosition = () => {
  const mutationFn = async ({
    listID,
    cardID,
    position,
  }: {
    listID: string
    cardID: string
    position: number
  }) => {
    await apiClient.put(`/list/${listID}/card/${cardID}`, { position })
  }
  const mutation = useMutation(mutationFn, {
    onSuccess: () => queryClient.invalidateQueries('board'),
  })

  return mutation
}

export const deleteListCard = async (listID: string, cardID: string) => {
  const { data }: { data: APIResponse } = await apiClient.delete(
    `/list/${listID}/card/${cardID}`
  )

  return data
}
