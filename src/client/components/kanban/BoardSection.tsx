import React from 'react'
import { Link } from 'react-router-dom'
import { useAuthenticatedUser } from '../../contexts/AuthContext'
import { Board } from '../../types'

const BoardSection = () => {
  const { user } = useAuthenticatedUser()

  if (!user) return null

  return (
    <div className="px-6 text-black h-screen overflow-auto">
      <div className="w-full px-4 mx-auto">
        <div className="relative flex flex-col min-w-0 break-words bg-white w-full p-8 mb-6 mt-16">
          <h1 className="text-3xl font-bold mb-10">Board section</h1>

          <div className="my-10">
            <div className="grid grid-cols-2 gap-x-20">
              {user.Boards.map((board: Board) => (
                <Link key={board.ID} to={`/boards/${board.ID}`}>
                  {board.Name}
                </Link>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BoardSection
