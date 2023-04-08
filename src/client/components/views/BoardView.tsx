import React from 'react'
import { DragDropContext } from 'react-beautiful-dnd'

import ListView from './ListsView'
import useUserBoard from '../../hooks/useUserBoard'

import { Board } from '../../types'

const BoardView = ({ id }: { id: string }) => {
  const { boardData, isLoading } = useUserBoard(id)

  if (!boardData || isLoading || !('Lists' in boardData.data)) return null

  const board: Board = boardData.data

  console.log('BOARDVIEW', board)

  const handleOnDragEnd = ({ source, destination }: any) => {
    console.log(source, destination)
  }

  return (
    <div className="px-6 text-black h-screen overflow-auto">
      <div className="overflow-y-hidden scrollbar-thin scrollbar-thumb-mainPurple scrollbar-track-transparent flex-1 p-4 space-x-4 flex">
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
