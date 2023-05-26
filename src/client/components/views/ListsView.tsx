import React, { useRef, useState } from 'react'
import { Droppable } from 'react-beautiful-dnd'

import { useDeleteList } from '../../services/listService'

import CardView from './CardView'
import NewCardView from './NewCardView'
import DropdownMenu from '../common/DropdownMenu'

import { Card, List } from '../../types'

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
  const newCardRef = useRef(null)
  const [showCardBones, setShowCardBones] = useState(false)

  sortCardsByPosition(list)

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
      <NewCardView list={list} />

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
          </ul>
        )}
      </Droppable>
    </div>
  )
}

export default ListView
