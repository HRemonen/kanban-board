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
        <Route index element={<HomeSection />} />
        <Route path="boards">
          <Route index element={<BoardSection />} />
          <Route path=":boardID" element={<BoardView />} />
        </Route>
      </Routes>
    </section>
  </>
)

export default Kanban
