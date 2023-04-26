import React from 'react'

import { DragDropContext, DragUpdate } from 'react-beautiful-dnd'

import ListView from './ListsView'
import useUserBoard from '../../hooks/useUserBoard'

import { useUpdateCardPosition } from '../../services/cardService'

import { Board } from '../../types'
import NewListView from './NewListView'

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

        <NewListView board={board} />
      </div>
    </div>
  )
}

export default BoardView
