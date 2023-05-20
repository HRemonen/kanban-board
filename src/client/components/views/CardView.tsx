import React, { useState } from 'react'
import { Link, Outlet } from 'react-router-dom'
import { Draggable } from 'react-beautiful-dnd'

import { Card } from '../../types'

const CardView = ({ card }: { card: Card }) => (
  <>
    <Draggable key={card.ID} draggableId={card.ID} index={card.Position}>
      {(provided) => (
        <li
          ref={provided.innerRef}
          {...provided.draggableProps}
          {...provided.dragHandleProps}
        >
          <div className='max-w-sm rounded-lg border border-gray-200 bg-white p-6 shadow'>
            <Link
              to={`./view/${card.ID}`}
              className='text-md mb-2 cursor-pointer break-all text-left font-light tracking-tight text-gray-900 hover:underline'
              data-cy={`view-card-info-button-${card.ID}`}
            >
              {card.Title}
            </Link>
          </div>
        </li>
      )}
    </Draggable>
    <Outlet />
  </>
)

export default CardView
