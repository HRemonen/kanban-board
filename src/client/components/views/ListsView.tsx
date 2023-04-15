import React from 'react'
import { Droppable, Draggable } from 'react-beautiful-dnd'

import { Card, List } from '../../types'

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
              <Draggable
                key={card.ID}
                draggableId={card.ID}
                index={card.Position}
              >
                {(provided) => (
                  <li
                    ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                  >
                    <div className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow">
                      <h5 className="mb-2 text-2xl font-semibold tracking-tight text-gray-900">
                        {card.Title}
                      </h5>
                      <p className="mb-3 font-normal text-gray-500">
                        {card.Description}
                      </p>
                    </div>
                  </li>
                )}
              </Draggable>
            ))}
            {provided.placeholder}
          </ul>
        )}
      </Droppable>
    </div>
  )
}

export default ListView
