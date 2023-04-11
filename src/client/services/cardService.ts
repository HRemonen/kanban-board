import apiClient from '../util/apiClient'
import { APIResponse, Card, List } from '../types'

export const getSingleCard = async (id: string) => {
  const { data }: { data: APIResponse } = await apiClient.get(`/card/${id}`)

  return data
}

export const CreateListCard = async (list: List, card: Card) => {
  const { data }: { data: APIResponse } = await apiClient.post(
    `/list/${list.ID}/card`,
    card
  )

  return data
}

export const UpdateListCardPosition = async (
  list: List,
  card: Card,
  position: number
) => {
  const { data }: { data: APIResponse } = await apiClient.post(
    `/list/${list.ID}/card/${card.ID}`,
    { position }
  )

  return data
}

export const DeleteListCard = async (list: List, card: Card) => {
  const { data }: { data: APIResponse } = await apiClient.delete(
    `/list/${list.ID}/card/${card.ID}`
  )

  return data
}
