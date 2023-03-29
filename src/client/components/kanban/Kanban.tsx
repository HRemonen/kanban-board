import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Sidebar from './Sidebar'
import HomeSection from './HomeSection'
import ProfileSection from './ProfileSection'

const Kanban = () => (
  <>
    <Sidebar />
    <section className="ml-60">
      <Routes>
        <Route path="/" element={<HomeSection />} />
        <Route path="profile" element={<ProfileSection />} />
      </Routes>
    </section>
  </>
)

export default Kanban
