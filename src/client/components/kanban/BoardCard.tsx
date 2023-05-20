import React from 'react'
import { Link } from 'react-router-dom'

import { useDeleteBoard } from '../../services/boardService'

import DropdownMenu from '../common/DropdownMenu'

import { Board } from '../../types'

const BoardCard = ({ board }: { board: Board }) => {
  const mutateBoards = useDeleteBoard()

  const handleBoardDelete = () => {
    mutateBoards.mutateAsync({
      boardID: board.ID,
    })
  }

  return (
    <div className='relative my-2 w-[100%] rounded-xl border bg-white p-4 md:my-4'>
      <div className='absolute right-4 top-4'>
        <DropdownMenu handleDelete={handleBoardDelete} />
      </div>
      <div>
        <Link
          className='mb-2 text-xl font-semibold'
          key={board.ID}
          to={`/boards/${board.ID}`}
        >
          {board.Name}
        </Link>
        <div className='my-2 flex space-x-2 text-sm text-gray-400 md:my-4'>
          <svg
            className='h-5 w-5'
            fill='currentColor'
            stroke='none'
            aria-hidden='true'
            height='16'
            viewBox='0 0 16 16'
            version='1.1'
            width='16'
            data-view-component='true'
          >
            <path d='M0 1.75A.75.75 0 0 1 .75 1h4.253c1.227 0 2.317.59 3 1.501A3.743 3.743 0 0 1 11.006 1h4.245a.75.75 0 0 1 .75.75v10.5a.75.75 0 0 1-.75.75h-4.507a2.25 2.25 0 0 0-1.591.659l-.622.621a.75.75 0 0 1-1.06 0l-.622-.621A2.25 2.25 0 0 0 5.258 13H.75a.75.75 0 0 1-.75-.75Zm7.251 10.324.004-5.073-.002-2.253A2.25 2.25 0 0 0 5.003 2.5H1.5v9h3.757a3.75 3.75 0 0 1 1.994.574ZM8.755 4.75l-.004 7.322a3.752 3.752 0 0 1 1.992-.572H14.5v-9h-3.495a2.25 2.25 0 0 0-2.25 2.25Z' />
          </svg>
          <i>{board.Description || 'No desctiption provided'}</i>
        </div>
        <div className='flex justify-between border-t-2'>
          <div className='my-2'>
            <p className='mb-2 text-base font-semibold'>Team Members</p>
            <div className='flex space-x-2'>
              <img
                alt='user profile'
                src='https://images.pexels.com/photos/614810/pexels-photo-614810.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500'
                className='h-6 w-6 rounded-full'
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BoardCard
