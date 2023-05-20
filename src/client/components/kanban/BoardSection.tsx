import React from 'react'
import { Link } from 'react-router-dom'

import { useUserBoards } from '../../services/boardService'
import { useAuthenticatedUser } from '../../contexts/AuthContext'

import BoardCard from './BoardCard'

import { Board } from '../../types'

const BoardSection = () => {
  const { user } = useAuthenticatedUser()
  const { userBoardsData, isSuccess } = useUserBoards(user?.id)

  if (!isSuccess || !userBoardsData || !Array.isArray(userBoardsData.data))
    return null

  const boards: Board[] = userBoardsData.data

  return (
    <div className='h-screen overflow-auto px-6 text-black'>
      <div className='mx-auto w-full px-4'>
        <div className='relative mb-6 mt-16 flex w-full min-w-0 flex-col break-words bg-white p-8'>
          <div className='flex justify-between'>
            <h1 className='text-3xl font-bold'>Board section</h1>
            <Link
              to='./new'
              className='group relative inline-flex items-center justify-start overflow-hidden rounded bg-white px-4 py-2 text-sm font-light transition-all hover:bg-white'
            >
              <span className='absolute bottom-0 left-0 mb-9 ml-9 h-48 w-48 -translate-x-full translate-y-full rotate-[-45deg] rounded bg-purple-600 transition-all duration-500 ease-out group-hover:mb-32 group-hover:ml-0 group-hover:translate-x-0' />
              <span className='relative w-full text-left text-black transition-colors duration-300 ease-in-out group-hover:text-white'>
                Create New Board
              </span>
            </Link>
          </div>
          <div className='my-10'>
            <div className='grid grid-cols-2 gap-x-20'>
              {boards.map((board: Board) => (
                <BoardCard key={board.ID} board={board} />
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BoardSection
