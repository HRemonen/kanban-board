import React from 'react'
import { Route, Routes } from 'react-router-dom'
import { useAuthenticatedUser } from './contexts/AuthContext'
import Login from './components/authentication/Login'
import Kanban from './components/kanban/Kanban'

const App = () => {
  const { user } = useAuthenticatedUser()

  return (
    <section className="relative text-white bg-slate-800 overflow-hidden max-h-screen">
      <Routes>
        <Route path="/*" element={user ? <Kanban /> : <Login />} />
      </Routes>
    </section>
  )
}

export default App
