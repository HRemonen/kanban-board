import { useMutation, useQuery } from 'react-query'

import apiClient from '../util/apiClient'
import queryClient from '../util/queryClient'

import { CardAPIResponse, NewCard } from '../types'

export const useCard = (cardID: string | undefined) => {
  const queryKey = ['card', cardID]

  const query = async (): Promise<CardAPIResponse> => {
    const { data }: { data: CardAPIResponse } = await apiClient.get(
      `/card/${cardID}`
    )
    return data
  }

  const { data: cardData, ...rest } = useQuery(queryKey, query)

  return { cardData, ...rest }
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

export const useDeleteListCard = () => {
  const mutationFn = async ({
    listID,
    cardID,
  }: {
    listID: string
    cardID: string
  }) => {
    await apiClient.delete(`/list/${listID}/card/${cardID}`)
  }
  const mutation = useMutation(mutationFn, {
    onSuccess: () => queryClient.invalidateQueries('board'),
  })

  return mutation
}
