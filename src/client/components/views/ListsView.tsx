import React from 'react'
import { Droppable, Draggable } from 'react-beautiful-dnd'

import { Card, List } from '../../types'
import CardView from './CardView'

function sortCardsByPosition(list: List): void {
  list.Cards.sort((a: Card, b: Card) => a.Position - b.Position)
}

const ListView = ({ list }: { list: List }) => {
  sortCardsByPosition(list)

  return (
    <div data-cy={`list-${list.ID}`} className="w-[280px] shrink-0">
      <h3 className="uppercase mb-4">
        {list.Name} ({list.Cards.length})
      </h3>
      <Droppable droppableId={list.ID}>
        {(provided) => (
          <ul
            data-cy={`list-${list.ID}-cards`}
            className="flex flex-col whitespace-nowrap overflow-y-scroll scrollbar-hide h-full pb-4 gap-2"
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
