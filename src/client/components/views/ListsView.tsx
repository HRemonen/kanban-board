import React, { useState } from 'react'
import { Droppable } from 'react-beautiful-dnd'

import { useDeleteList } from '../../services/listService'

import CardView from './CardView'
import NewCardView from './NewCardView'

import { Card, List } from '../../types'

function sortCardsByPosition(list: List): void {
  list.Cards.sort((a: Card, b: Card) => a.Position - b.Position)
}

const DropdownMenu = ({ list }: { list: List }) => {
  const [showDropdown, setShowDropdown] = useState(false)
  const mutateList = useDeleteList()

  const onListDelete = () => {
    mutateList.mutateAsync({
      boardID: list.BoardID,
      listID: list.ID,
    })

    setShowDropdown(false)
  }

  return (
    <div className="relative">
      <button
        type="button"
        className="inline-flex items-center p-2 text-sm font-medium text-center text-gray-900 bg-white rounded-lg hover:bg-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-50"
        onClick={() => setShowDropdown(!showDropdown)}
      >
        <span className="inline-block align-text-bottom select-text overflow-visible">
          <svg
            aria-hidden="true"
            focusable="false"
            role="img"
            viewBox="0 0 16 16"
            width="16"
            height="16"
            fill="currentColor"
          >
            <path d="M8 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3ZM1.5 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Zm13 0a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Z" />
          </svg>
        </span>
      </button>
      {showDropdown && (
        <div className="absolute z-10 bg-white divide-y divide-gray-200 rounded-lg shadow w-30">
          <ul
            className="pt-2 text-sm text-gray-700"
            aria-labelledby="dropdownMenuIconButton"
          >
            <li className="rounded-lg hover:bg-gray-100">
              <button type="button" className="block px-4 py-2">
                Edit
              </button>
            </li>
          </ul>
          <div className="mt-2">
            <button
              type="button"
              className="block px-4 py-2 text-sm text-red-500 rounded-lg hover:bg-red-100"
              onClick={onListDelete}
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
  const [showModal, setShowModal] = useState(false)

  sortCardsByPosition(list)

  return (
    <div data-cy={`list-${list.ID}`} className="w-[280px] shrink-0">
      <div className="flex ">
        <h3 className="uppercase mb-4">
          {list.Name}
          <span className="mx-2 inline-flex items-center justify-center bg-blue-100 text-blue-800 text-sm font-semibold px-2 py-[2px] rounded-full">
            {list.Cards.length}
          </span>
        </h3>
        <DropdownMenu list={list} />
      </div>
      <div className="mt-2 px-2 hover:bg-gray-300 hover:rounded-lg">
        <button
          type="button"
          className="inline-flex items-center py-2 text-gray-600 text-sm font-medium"
          data-cy={`add-card-button-${list.ID}`}
          onClick={() => setShowModal(true)}
        >
          <span className="mr-2 inline-block align-text-bottom select-text overflow-visible">
            <svg
              aria-hidden="true"
              focusable="false"
              role="img"
              viewBox="0 0 16 16"
              width="16"
              height="16"
              fill="currentColor"
            >
              <path d="M7.75 2a.75.75 0 0 1 .75.75V7h4.25a.75.75 0 0 1 0 1.5H8.5v4.25a.75.75 0 0 1-1.5 0V8.5H2.75a.75.75 0 0 1 0-1.5H7V2.75A.75.75 0 0 1 7.75 2Z" />
            </svg>
          </span>
          Add card
        </button>
      </div>

      <Droppable droppableId={list.ID}>
        {(provided) => (
          <ul
            data-cy={`list-${list.ID}-cards`}
            className="flex flex-col whitespace-nowrap overflow-y-scroll scrollbar-hide h-[90vh] pb-4 gap-2"
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

      {showModal && <NewCardView list={list} setShowModal={setShowModal} />}
    </div>
  )
}

export default ListView
