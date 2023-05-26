import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { DragDropContext, DragUpdate } from 'react-beautiful-dnd'

import ListView from './ListsView'
import NewListView from './NewListView'

import { useUpdateCardPosition } from '../../services/cardService'
import { useCreateNewList } from '../../services/listService'
import { useUserBoard } from '../../services/boardService'

import { Board, Card, List } from '../../types'

const BoardView = () => {
  const { boardID: id } = useParams()
  const mutateList = useCreateNewList()
  const mutateCardPosition = useUpdateCardPosition()
  const { boardData, isLoading } = useUserBoard(id)
  const [lists, setLists] = useState<List[]>()

  useEffect(() => {
    if (!boardData || isLoading || !('Lists' in boardData.data)) return
    setLists(boardData.data.Lists)
  }, [boardData, isLoading])

  if (!boardData || isLoading || !('Lists' in boardData.data) || !lists)
    return null

  const board: Board = boardData.data

  const reorder = (list: Card[], card: Card, endPosition: number) => {
    const result = Array.from(list) // no side effects man
    const currentPosition = card.Position

    // Position not updated so don't do anything
    if (endPosition === currentPosition) return result

    const cardToUpdate = result.find((aCard) => aCard.ID === card.ID)
    if (!cardToUpdate) return result

    /* eslint no-param-reassign: "error" */

    if (endPosition < currentPosition) {
      // shift items between new and old position up by 1
      result.forEach((aCard) => {
        if (aCard.Position >= endPosition && aCard.Position < currentPosition)
          aCard.Position += 1
      })
    } else {
      // shift items between new and old position down by 1
      result.forEach((aCard) => {
        if (aCard.Position > currentPosition && aCard.Position <= endPosition)
          aCard.Position -= 1
      })
    }

    // Update the cards position that user wanted to move
    cardToUpdate.Position = endPosition

    return result
  }

  const handleOnDragEnd = ({
    draggableId,
    source,
    destination,
  }: DragUpdate) => {
    if (!destination) return

    const listToUpdate = lists.find((list) => list.ID === source.droppableId)
    if (!listToUpdate) return

    const cards: Card[] = listToUpdate.Cards

    const cardToUpdate = cards.find((c) => c.ID === draggableId)
    if (!cardToUpdate) return

    const updatedList = {
      ...listToUpdate,
      Cards: reorder(cards, cardToUpdate, destination.index),
    }

    const updatedLists = lists.map((list) =>
      list.ID === updatedList.ID ? { ...updatedList } : { ...list }
    )

    setLists(updatedLists)

    mutateCardPosition.mutateAsync({
      listID: source.droppableId,
      cardID: draggableId,
      position: destination.index,
    })
  }

  if (lists.length === 0) {
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
          {lists.map((list) => (
            <ListView key={list.ID} list={list} />
          ))}
          <NewListView board={board} />
        </DragDropContext>
      </div>
    </div>
  )
}

export default BoardView
