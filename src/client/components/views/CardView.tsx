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
          <div className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow">
            <Link
              to={`./view/${card.ID}`}
              className="break-all text-left mb-2 text-md font-light tracking-tight text-gray-900 cursor-pointer hover:underline"
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
