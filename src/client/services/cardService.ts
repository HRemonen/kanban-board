import { useMutation } from 'react-query'

import apiClient from '../util/apiClient'

import { APIResponse, Card } from '../types'

export const getSingleCard = async (id: string) => {
  const { data }: { data: APIResponse } = await apiClient.get(`/card/${id}`)

  return data
}

export const createListCard = async (listID: string, card: Card) => {
  const { data }: { data: APIResponse } = await apiClient.post(
    `/list/${listID}/card`,
    card
  )

  return data
}

export const useUpdateListCardPosition = () => {
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
  const mutation = useMutation(mutationFn)

  return mutation
}

export const updateListCardPosition = async (
  listID: string,
  cardID: string,
  position: number
) => {
  const { data }: { data: APIResponse } = await apiClient.post(
    `/list/${listID}/card/${cardID}`,
    { position }
  )

  return data
}

export const deleteListCard = async (listID: string, cardID: string) => {
  const { data }: { data: APIResponse } = await apiClient.delete(
    `/list/${listID}/card/${cardID}`
  )

  return data
}
