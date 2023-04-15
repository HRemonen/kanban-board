import React from 'react'
import { Draggable } from 'react-beautiful-dnd'

import { Card } from '../../types'

const CardView = ({ card }: { card: Card }) => (
  <Draggable key={card.ID} draggableId={card.ID} index={card.Position}>
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
          <p className="mb-3 font-normal text-gray-500">{card.Description}</p>
        </div>
      </li>
    )}
  </Draggable>
)

export default CardView
