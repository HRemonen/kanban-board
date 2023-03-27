import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Sidebar from './Sidebar'
import HomeSection from './HomeSection'

const Kanban = () => (
  <>
    <Sidebar />
    <Routes>
      <Route path="/" element={<HomeSection />} />
    </Routes>
  </>
)

export default Kanban
