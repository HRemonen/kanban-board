import React from 'react'
import { Link } from 'react-router-dom'

import { Board } from '../../types'

const BoardCard = ({ board }: { board: Board }) => (
  <div className="relative bg-white p-4 rounded-xl w-[100%] my-2 md:my-4 border">
    <div className="mt-2">
      <Link
        className="text-xl font-semibold my-2"
        key={board.ID}
        to={`/boards/${board.ID}`}
      >
        {board.Name}
      </Link>
      <div className="flex space-x-2 text-gray-400 text-sm">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
          />
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
          />
        </svg>
        <p>{board.Description}</p>
      </div>
      <div className="flex space-x-2 text-gray-400 text-sm my-3">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-5 w-5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
          />
        </svg>
        <p>1 Weeks Left</p>
      </div>
      <div className="flex justify-between border-t-2">
        <div className="my-2">
          <p className="font-semibold text-base mb-2">Team Members</p>
          <div className="flex space-x-2">
            <img
              alt="user profile"
              src="https://images.pexels.com/photos/614810/pexels-photo-614810.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500"
              className="w-6 h-6 rounded-full"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
)

export default BoardCard
