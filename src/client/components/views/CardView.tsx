import React, { useState } from 'react'
import { Draggable } from 'react-beautiful-dnd'
import CardInfoView from './CardInfoView'

import { Card } from '../../types'

const CardView = ({ card }: { card: Card }) => {
  const [showModal, setShowModal] = useState(false)

  return (
    <>
      <Draggable key={card.ID} draggableId={card.ID} index={card.Position}>
        {(provided) => (
          <li
            ref={provided.innerRef}
            {...provided.draggableProps}
            {...provided.dragHandleProps}
          >
            <div className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow">
              <button
                type="button"
                className="break-all text-left mb-2 text-md font-light tracking-tight text-gray-900 cursor-pointer hover:underline"
                data-cy={`view-card-info-button-${card.ID}`}
                onClick={() => setShowModal(true)}
              >
                {card.Title}
              </button>
            </div>
          </li>
        )}
      </Draggable>
      {showModal && <CardInfoView card={card} setShowModal={setShowModal} />}
    </>
  )
}

export default CardView
