import React from 'react'
import { Card } from '../../types'

type ModalProps = {
  card: Card
  setShowModal: React.Dispatch<React.SetStateAction<boolean>>
}

const ModalView = ({ card, setShowModal }: ModalProps) => {
  console.log(card)
  return (
    <div className="fixed top-0 right-0 z-40 h-screen p-4 bg-white w-[60%] border-l-2">
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
        {card.Title}
      </h5>
      <button
        type="button"
        className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 absolute top-2.5 right-2.5 inline-flex items-center"
        onClick={() => setShowModal(false)}
      >
        <svg
          aria-hidden="true"
          className="w-5 h-5"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fillRule="evenodd"
            d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
            clipRule="evenodd"
          />
        </svg>
        <span className="sr-only">Close menu</span>
      </button>
      <div className="grid grid-cols-12 gap-2">
        <div className="col-span-12">
          <h3 className="mb-2 text-xl font-bold tracking-tight text-gray-900">
            {card.Title}
          </h3>
        </div>
        <div className="col-span-7">
          <p className="font-light text-gray-700">
            {card.Description || 'No description provided'}
          </p>
        </div>
        <div className="col-span-5 row-span-6">
          <div className="grid grid-cols-2 gap-4">
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
          </div>
        </div>
        <div className="col-span-7">Additional section (TBA)</div>
      </div>
    </div>
  )
}

export default ModalView
