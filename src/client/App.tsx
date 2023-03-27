import React, { useContext } from 'react'
import { Route, Routes } from 'react-router-dom'
import { AuthProvider, useAuthenticatedUser } from './contexts/AuthContext'
import Login from './components/authentication/Login'

const App = () => {
  const user = useAuthenticatedUser()

  console.log(user)
  return (
    <section className="bg-slate-900 text-white">
      <Routes>
        <Route path="/" element={<Login />} />
      </Routes>
    </section>
  )
}

export default App
