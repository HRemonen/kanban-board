import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Sidebar from './Sidebar'
import HomeSection from './HomeSection'
import BoardSection from './BoardSection'
import BoardView from '../views/BoardView'
import CardInfoView from '../views/CardInfoView'
import NewBoardView from '../views/NewBoardView'

const Kanban = () => (
  <>
    <Sidebar />
    <section className='ml-60'>
      <Routes>
        <Route index element={<HomeSection />} />
        <Route path='boards'>
          <Route index element={<BoardSection />} />
          <Route path=':boardID' element={<BoardView />}>
            <Route path='view/:cardID' element={<CardInfoView />} />
          </Route>
          <Route path='new' element={<NewBoardView />} />
        </Route>
      </Routes>
    </section>
  </>
)

export default Kanban
