import React from 'react'
import { useParams } from 'react-router-dom'
import { DragDropContext, DragUpdate } from 'react-beautiful-dnd'

import ListView from './ListsView'

import { useUpdateCardPosition } from '../../services/cardService'
import { useCreateNewList } from '../../services/listService'
import { useUserBoard } from '../../services/boardService'

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
    if (!destination) return

    mutateCardPosition.mutateAsync({
      listID: source.droppableId,
      cardID: draggableId,
      position: destination.index,
    })
  }

  if (board.Lists.length === 0) {
    const handleCreateList = () => {
      mutateList.mutateAsync({
        boardID: board.ID,
        list: { name: 'Todo list' },
      })
    }
    return (
      <div
        data-cy={`empty-board-${board.ID}`}
        className='h-screen overflow-y-hidden px-6 text-black'
      >
        <div className='flex h-full w-full flex-col items-center justify-center'>
          <h2 className='text-center text-lg'>
            This board is empty. Create a new list to get started.
          </h2>
          <button type='button' onClick={handleCreateList} className='mt-6'>
            + Add New List
          </button>
        </div>
      </div>
    )
  }

  return (
    <div
      data-cy={`board-${board.ID}`}
      className='h-screen overflow-y-hidden px-6 text-black'
    >
      <div className='flex flex-1 space-x-4 overflow-y-hidden p-4 scrollbar-hide'>
        <DragDropContext onDragEnd={handleOnDragEnd}>
          {board.Lists.map((list) => (
            <ListView key={list.ID} listID={list.ID} />
          ))}
          <NewListView board={board} />
        </DragDropContext>
      </div>
    </div>
  )
}

export default BoardView
