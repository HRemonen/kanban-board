import React from 'react'
import { Route, Routes } from 'react-router-dom'
import { useAuthenticatedUser } from './contexts/AuthContext'
import Login from './components/authentication/Login'
import Kanban from './components/kanban/Kanban'

const App = () => {
  const { user, config } = useAuthenticatedUser()

  return (
    <section className="relative text-white bg-white overflow-hidden max-h-screen">
      <Routes>
        <Route path="/*" element={user ? <Kanban /> : <Login />} />
      </Routes>
    </section>
  )
}

export default App
