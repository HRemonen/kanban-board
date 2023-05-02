import React from 'react'
import { useParams } from 'react-router-dom'
import { DragDropContext, DragUpdate } from 'react-beautiful-dnd'

import ListView from './ListsView'
import useUserBoard from '../../hooks/useUserBoard'

import { useUpdateCardPosition } from '../../services/cardService'
import { useCreateNewList } from '../../services/listService'

import { Board } from '../../types'
import NewListView from './NewListView'

const BoardView = () => {
  const { boardID: id } = useParams()
  const mutateList = useCreateNewList()
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

  if (board.Lists.length === 0) {
    const onClick = () => {
      mutateList.mutateAsync({
        boardID: board.ID,
        list: { name: 'Todo list' },
      })
    }
    return (
      <div
        data-cy={`empty-board-${board.ID}`}
        className="px-6 text-black h-screen overflow-y-hidden"
      >
        <div className="flex flex-col justify-center items-center w-full h-full">
          <h2 className="text-lg text-center">
            This board is empty. Create a new list to get started.
          </h2>
          <button type="button" onClick={onClick} className="mt-6">
            + Add New List
          </button>
        </div>
      </div>
    )
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
