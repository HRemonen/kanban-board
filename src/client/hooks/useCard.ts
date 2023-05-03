import { useQuery } from 'react-query'

import apiClient from '../util/apiClient'

import { APIResponse } from '../types'

const useCard = (cardID: string | undefined) => {
  const queryKey = ['card', cardID]

  const query = async (): Promise<APIResponse> => {
    const { data }: { data: APIResponse } = await apiClient.get(
      `/card/${cardID}`
    )
    return data
  }

  const { data: cardData, ...rest } = useQuery(queryKey, query)

  return { cardData, ...rest }
}

export default useCard
