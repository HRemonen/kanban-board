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
    <div className="px-6 text-black h-screen overflow-auto">
      <div className="w-full px-4 mx-auto">
        <div className="relative flex flex-col min-w-0 break-words bg-white w-full p-8 mb-6 mt-16">
          <div className="flex justify-between">
            <h1 className="text-3xl font-bold">Board section</h1>
            <Link
              to="./new"
              className="relative inline-flex items-center justify-start px-4 py-2 overflow-hidden text-sm font-light transition-all bg-white rounded hover:bg-white group"
            >
              <span className="w-48 h-48 rounded rotate-[-45deg] bg-purple-600 absolute bottom-0 left-0 -translate-x-full ease-out duration-500 transition-all translate-y-full mb-9 ml-9 group-hover:ml-0 group-hover:mb-32 group-hover:translate-x-0" />
              <span className="relative w-full text-left text-black transition-colors duration-300 ease-in-out group-hover:text-white">
                Create New Board
              </span>
            </Link>
          </div>
          <div className="my-10">
            <div className="grid grid-cols-2 gap-x-20">
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
