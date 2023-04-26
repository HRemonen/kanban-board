import React from 'react'
import { Card } from '../../types'
import { useDeleteListCard } from '../../services/cardService'
import CloseMenu from '../common/CloseMenu'

type ModalProps = {
  card: Card
  setShowModal: React.Dispatch<React.SetStateAction<boolean>>
}

const CardInfoView = ({ card, setShowModal }: ModalProps) => {
  const mutateCard = useDeleteListCard()

  const onCardDelete = () => {
    mutateCard.mutateAsync({
      listID: card.ListID,
      cardID: card.ID,
    })

    setShowModal(false)
  }

  return (
    <div className="fixed pin top-0 right-0 z-40 h-screen p-4 bg-white w-[60%] border-l-2">
      <h5
        id="drawer-label"
        className="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase"
      >
        <svg
          className="w-5 h-5 mr-2"
          aria-hidden="true"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fillRule="evenodd"
            d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
            clipRule="evenodd"
          />
        </svg>
        Card information
      </h5>
      <CloseMenu onClick={() => setShowModal(false)} />
      <div className="grid grid-cols-12 gap-2">
        <div className="col-span-12 border-b-2">
          <h3 className="mb-2 text-xl font-bold tracking-tight text-gray-900">
            {card.Title}
          </h3>
        </div>
        <div className="col-span-7">
          <p className="my-4 text-sm font-thin italic text-gray-500">
            {card.Description || 'No description provided'}
          </p>
        </div>
        <div className="col-span-5 row-span-6 border-l-2">
          <div className="my-4 mx-2 grid grid-cols-2 gap-4">
            <div className="font-normal text-sm text-gray-700">Assignees</div>
            <div className="font-light text-sm text-gray-400">
              Add assignees
            </div>
            <div className="font-normal text-sm text-gray-700">Labels</div>
            <div className="font-light text-sm text-gray-400">
              {card.Label || 'Add labels'}
            </div>
            <div className="font-normal text-sm text-gray-700">Status</div>
            <div className="font-light text-sm text-gray-400">
              <button
                type="button"
                className="px-4 bg-gray-100 hover:bg-gray-200 text-gray-800 text-sm font-medium rounded-full"
              >
                {card.Status || 'Add status'}
              </button>
            </div>
            <div className="my-4 border-t-2 col-span-2">
              <div className="mt-2 px-2 hover:bg-red-300 hover:rounded-lg">
                <button
                  type="button"
                  className="inline-flex items-center py-2 text-red-600 text-sm font-medium"
                  onClick={onCardDelete}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
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
