import React from 'react'
import { Route, Routes } from 'react-router-dom'
import { useAuthenticatedUser } from '../../contexts/AuthContext'
import Sidebar from './Sidebar'
import HomeSection from './HomeSection'

const Kanban = () => {
  const user = useAuthenticatedUser()

  return (
    <>
      <Sidebar />
      <Routes>
        <Route path="/" element={<HomeSection />} />
      </Routes>
    </>
  )
}

export default Kanban
