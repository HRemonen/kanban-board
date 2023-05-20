import React from 'react'
import { Link } from 'react-router-dom'

import { useAuthenticatedUser } from '../../contexts/AuthContext'
import { useUserBoards } from '../../services/boardService'

import { Board } from '../../types'

const StatsSection = () => {
  const { user } = useAuthenticatedUser()
  const { userBoardsData, isSuccess } = useUserBoards(user?.id)

  if (
    !isSuccess ||
    !userBoardsData ||
    !Array.isArray(userBoardsData.data) ||
    !user
  )
    return null

  const boards: Board[] = userBoardsData.data

  return (
    <div>
      <h2 className='mb-4 text-2xl font-bold'>Stats</h2>

      <div className='grid grid-cols-2 gap-4'>
        <div className='col-span-2'>
          <div className='rounded-xl bg-green-100 p-4'>
            <div className='text-xl font-bold leading-none text-gray-800'>
              Hello {user.username}!
            </div>
            <div className='mt-5'>
              <Link
                to='/boards/new'
                className='inline-flex items-center justify-center rounded-xl bg-white px-3 py-2 text-sm font-semibold text-gray-800 transition hover:text-green-400'
              >
                Create a new board
              </Link>
            </div>
          </div>
        </div>
        <div className='rounded-xl bg-yellow-100 p-4 text-gray-800'>
          <div className='text-2xl font-bold leading-none'>
            {boards.length ?? 0}
          </div>
          <div className='mt-2'>Active boards</div>
        </div>
        <div className='rounded-xl bg-yellow-100 p-4 text-gray-800'>
          <div className='text-2xl font-bold leading-none'>0</div>
          <div className='mt-2'>Incomplete cards</div>
        </div>
        <div className='col-span-2'>
          <div className='rounded-xl bg-purple-100 p-4 text-gray-800'>
            <div className='text-xl font-bold leading-none'>
              Your daily plan
            </div>
            <div className='mt-2'>5 of 8 completed</div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default StatsSection
