import React, { useState } from 'react'
import { Droppable } from 'react-beautiful-dnd'

import { useCreateNewCard } from '../../services/cardService'
import { useDeleteList } from '../../services/listService'

import CardView from './CardView'
import NewCardView from './NewCardView'

import { Card, List, NewCard } from '../../types'
import DropdownMenu from '../common/DropdownMenu'

function sortCardsByPosition(list: List): void {
  list.Cards.sort((a: Card, b: Card) => a.Position - b.Position)
}

const ListDropdownMenu = ({ list }: { list: List }) => {
  const mutateList = useDeleteList()

  const handleListDelete = () => {
    mutateList.mutateAsync({
      boardID: list.BoardID,
      listID: list.ID,
    })
  }

  return <DropdownMenu handleDelete={handleListDelete} />
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
        <ListDropdownMenu list={list} />
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
            className='flex h-[90vh] flex-col gap-2 overflow-y-scroll pb-4 scrollbar-hide'
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
