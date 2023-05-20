import React, { useState } from 'react'
import { Droppable } from 'react-beautiful-dnd'

import { useCreateNewCard } from '../../services/cardService'
import { useDeleteList } from '../../services/listService'

import CardView from './CardView'
import NewCardView from './NewCardView'

import { Card, List, NewCard } from '../../types'

function sortCardsByPosition(list: List): void {
  list.Cards.sort((a: Card, b: Card) => a.Position - b.Position)
}

const DropdownMenu = ({ list }: { list: List }) => {
  const [showDropdown, setShowDropdown] = useState(false)
  const mutateList = useDeleteList()

  const handleListDelete = () => {
    mutateList.mutateAsync({
      boardID: list.BoardID,
      listID: list.ID,
    })
    setShowDropdown(false)
  }

  return (
    <div className='relative'>
      <button
        type='button'
        className='inline-flex items-center rounded-lg bg-white p-2 text-center text-sm font-medium text-gray-900 hover:bg-gray-200 focus:outline-none focus:ring-4 focus:ring-gray-50'
        onClick={() => setShowDropdown(!showDropdown)}
      >
        <span className='inline-block select-text overflow-visible align-text-bottom'>
          <svg
            aria-hidden='true'
            focusable='false'
            role='img'
            viewBox='0 0 16 16'
            width='16'
            height='16'
            fill='currentColor'
          >
            <path d='M8 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3ZM1.5 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Zm13 0a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Z' />
          </svg>
        </span>
      </button>
      {showDropdown && (
        <div className='w-30 absolute z-10 divide-y divide-gray-200 rounded-lg bg-white shadow'>
          <ul
            className='pt-2 text-sm text-gray-700'
            aria-labelledby='dropdownMenuIconButton'
          >
            <li className='rounded-lg hover:bg-gray-100'>
              <button type='button' className='block px-4 py-2'>
                Edit
              </button>
            </li>
          </ul>
          <div className='mt-2'>
            <button
              type='button'
              className='block rounded-lg px-4 py-2 text-sm text-red-500 hover:bg-red-100'
              onClick={handleListDelete}
            >
              Delete
            </button>
          </div>
        </div>
      )}
    </div>
  )
}

const ListView = ({ list }: { list: List }) => {
  const [showCardBones, setShowCardBones] = useState(false)
  const mutateCard = useCreateNewCard()

  sortCardsByPosition(list)

  const handleCreateCard = (data: NewCard) => {
    mutateCard.mutateAsync({
      listID: list.ID,
      card: data,
    })

    setShowCardBones(false)
  }

  return (
    <div data-cy={`list-${list.ID}`} className='w-[280px] shrink-0'>
      <div className='flex '>
        <h3 className='mb-4 uppercase'>
          {list.Name}
          <span className='mx-2 inline-flex items-center justify-center rounded-full bg-blue-100 px-2 py-[2px] text-sm font-semibold text-blue-800'>
            {list.Cards.length}
          </span>
        </h3>
        <DropdownMenu list={list} />
      </div>
      <div className='mt-2 px-2 hover:rounded-lg hover:bg-gray-300'>
        <button
          type='button'
          className='inline-flex items-center py-2 text-sm font-medium text-gray-600'
          data-cy={`add-card-button-${list.ID}`}
          onClick={() => setShowCardBones(true)}
        >
          <span className='mr-2 inline-block select-text overflow-visible align-text-bottom'>
            <svg
              aria-hidden='true'
              focusable='false'
              role='img'
              viewBox='0 0 16 16'
              width='16'
              height='16'
              fill='currentColor'
            >
              <path d='M7.75 2a.75.75 0 0 1 .75.75V7h4.25a.75.75 0 0 1 0 1.5H8.5v4.25a.75.75 0 0 1-1.5 0V8.5H2.75a.75.75 0 0 1 0-1.5H7V2.75A.75.75 0 0 1 7.75 2Z' />
            </svg>
          </span>
          Add card
        </button>
      </div>

      <Droppable droppableId={list.ID}>
        {(provided) => (
          <ul
            data-cy={`list-${list.ID}-cards`}
            className='scrollbar-hide flex h-[90vh] flex-col gap-2 overflow-y-scroll pb-4'
            {...provided.droppableProps}
            ref={provided.innerRef}
          >
            {list.Cards.map((card) => (
              <CardView key={card.ID} card={card} />
            ))}
            {provided.placeholder}
            {showCardBones && <NewCardView onSubmit={handleCreateCard} />}
          </ul>
        )}
      </Droppable>
    </div>
  )
}

export default ListView
