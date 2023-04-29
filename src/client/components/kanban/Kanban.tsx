import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Sidebar from './Sidebar'
import HomeSection from './HomeSection'
import BoardSection from './BoardSection'
import BoardView from '../views/BoardView'

const Kanban = () => (
  <>
    <Sidebar />
    <section className="ml-60">
      <Routes>
        <Route path="/" element={<HomeSection />} />
        <Route path="/boards/:boardID" element={<BoardView />} />
        <Route path="/boards" element={<BoardSection />} />
      </Routes>
    </section>
  </>
)

export default Kanban
