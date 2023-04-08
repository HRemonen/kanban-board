import React from 'react'
import { useAuthenticatedUser } from '../../contexts/AuthContext'
import BoardView from '../views/BoardView'

const BoardSection = () => {
  const { user } = useAuthenticatedUser()

  if (!user) return null

  return (
    <div className="px-6 text-black h-screen overflow-auto">
      <BoardView id={user.Boards[0].ID} />
    </div>
  )
}

export default BoardSection
