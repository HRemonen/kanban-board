import React from 'react'
import { DragDropContext, DragUpdate } from 'react-beautiful-dnd'

import ListView from './ListsView'
import useUserBoard from '../../hooks/useUserBoard'

import { Board } from '../../types'

const BoardView = ({ id }: { id: string }) => {
  const { boardData, isLoading } = useUserBoard(id)

  if (!boardData || isLoading || !('Lists' in boardData.data)) return null

  const board: Board = boardData.data

  const handleOnDragEnd = ({
    draggableId,
    source,
    destination,
  }: DragUpdate) => {
    console.log('INITIAL', source)
    console.log('DESTINATION', destination)

    console.log('CARD', draggableId)
  }

  return (
    <div
      data-cy={`board-${board.ID}`}
      className="px-6 text-black h-screen overflow-auto"
    >
      <div className="flex flex-1 overflow-y-hidden scrollbar-hide p-4 space-x-4">
        <DragDropContext onDragEnd={handleOnDragEnd}>
          {board.Lists.map((list) => (
            <ListView key={list.ID} list={list} />
          ))}
        </DragDropContext>
      </div>
    </div>
  )
}

export default BoardView
