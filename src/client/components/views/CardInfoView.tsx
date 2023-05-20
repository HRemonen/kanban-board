import React from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useCard, useDeleteListCard } from '../../services/cardService'

import CloseMenu from '../common/CloseMenu'
import { Card } from '../../types'

const CardInfoView = () => {
  const { cardID } = useParams()
  const navigate = useNavigate()
  const { cardData, isLoading } = useCard(cardID)
  const mutateCard = useDeleteListCard()

  if (!cardData || isLoading || !('Status' in cardData.data)) return null

  const card: Card = cardData.data

  const onCardDelete = () => {
    mutateCard.mutateAsync({
      listID: card.ListID,
      cardID: card.ID,
    })
    navigate(-1)
  }

  return (
    <div className='pin fixed right-0 top-0 z-40 h-screen w-[60%] border-l-2 bg-white p-4'>
      <h5
        id='drawer-label'
        className='mb-6 inline-flex items-center text-base font-semibold uppercase text-gray-500'
      >
        <svg
          className='mr-2 h-5 w-5'
          aria-hidden='true'
          fill='currentColor'
          viewBox='0 0 20 20'
          xmlns='http://www.w3.org/2000/svg'
        >
          <path
            fillRule='evenodd'
            d='M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z'
            clipRule='evenodd'
          />
        </svg>
        Card information
      </h5>
      <CloseMenu onClick={() => navigate(-1)} />
      <div className='grid grid-cols-12 gap-2'>
        <div className='col-span-12 border-b-2'>
          <h3 className='mb-2 text-xl font-bold tracking-tight text-gray-900'>
            {card.Title}
          </h3>
        </div>
        <div className='col-span-7'>
          <p className='my-4 text-sm font-thin italic text-gray-500'>
            {card.Description || 'No description provided'}
          </p>
        </div>
        <div className='col-span-5 row-span-6 border-l-2'>
          <div className='mx-2 my-4 grid grid-cols-2 gap-4'>
            <div className='text-sm font-normal text-gray-700'>Assignees</div>
            <div className='text-sm font-light text-gray-400'>
              Add assignees
            </div>
            <div className='text-sm font-normal text-gray-700'>Labels</div>
            <div className='text-sm font-light text-gray-400'>
              {card.Label || 'Add labels'}
            </div>
            <div className='text-sm font-normal text-gray-700'>Status</div>
            <div className='text-sm font-light text-gray-400'>
              <button
                type='button'
                className='rounded-full bg-gray-100 px-4 text-sm font-medium text-gray-800 hover:bg-gray-200'
              >
                {card.Status || 'Add status'}
              </button>
            </div>
            <div className='col-span-2 my-4 border-t-2'>
              <div className='mt-2 px-2 hover:rounded-lg hover:bg-red-300'>
                <button
                  type='button'
                  className='inline-flex items-center py-2 text-sm font-medium text-red-600'
                  onClick={onCardDelete}
                >
                  <svg
                    xmlns='http://www.w3.org/2000/svg'
                    className='mr-2 h-5 w-5'
                    fill='none'
                    viewBox='0 0 24 24'
                    stroke='currentColor'
                  >
                    <path
                      strokeLinecap='round'
                      strokeLinejoin='round'
                      strokeWidth='2'
                      d='M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16'
                    />
                  </svg>
                  Delete from board
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default CardInfoView
