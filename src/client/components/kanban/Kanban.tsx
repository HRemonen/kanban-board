import React from 'react'
import { useAuthenticatedUser } from '../../contexts/AuthContext'
import Sidebar from './Sidebar'
import HomeSection from './HomeSection'

const Kanban = () => {
  const user = useAuthenticatedUser()

  return (
    <>
      <Sidebar />
      <HomeSection />
    </>
  )
}

export default Kanban
