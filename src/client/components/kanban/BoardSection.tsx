import React from 'react'
import { Link } from 'react-router-dom'
import { useAuthenticatedUser } from '../../contexts/AuthContext'
import { Board } from '../../types'

const BoardSection = () => {
  const { user } = useAuthenticatedUser()

  if (!user) return null

  return (
    <div className="px-6 text-black h-screen overflow-auto">
      {user.Boards.map((board: Board) => (
        <Link key={board.ID} to={`/boards/${board.ID}`}>
          {' '}
          {board.Name}{' '}
        </Link>
      ))}
    </div>
  )
}

export default BoardSection
