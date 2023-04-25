import React from 'react'

import { DragDropContext, DragUpdate } from 'react-beautiful-dnd'

import ListView from './ListsView'
import useUserBoard from '../../hooks/useUserBoard'

import { useUpdateCardPosition } from '../../services/cardService'

import { Board } from '../../types'

const BoardView = ({ id }: { id: string }) => {
  const mutateCardPosition = useUpdateCardPosition()
  const { boardData, isLoading } = useUserBoard(id)

  if (!boardData || isLoading || !('Lists' in boardData.data)) return null

  const board: Board = boardData.data

  const handleOnDragEnd = ({
    draggableId,
    source,
    destination,
  }: DragUpdate) => {
    if (!destination) return null

    mutateCardPosition.mutateAsync({
      listID: source.droppableId,
      cardID: draggableId,
      position: destination.index,
    })

    return null
  }

  return (
    <div
      data-cy={`board-${board.ID}`}
      className="px-6 text-black h-screen overflow-y-hidden"
    >
      <div className="flex flex-1 overflow-y-hidden scrollbar-hide p-4 space-x-4">
        <DragDropContext onDragEnd={handleOnDragEnd}>
          {board.Lists.map((list) => (
            <ListView key={list.ID} list={list} />
          ))}
        </DragDropContext>

        <div className="mt-20">
          <button
            type="button"
            aria-label="Add a new list to the board"
            data-cy="add-new-list-button"
            className="p-2 inline-block align-text-bottom select-none overflow-visible border border-gray-200 rounded-lg shadow "
          >
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
          </button>
        </div>
      </div>
    </div>
  )
}

export default BoardView
