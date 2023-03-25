import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Login from './components/authentication/Login'
import { AuthProvider } from './contexts/AuthContext'

const App = () => (
  <section className="bg-slate-900 text-white">
    <AuthProvider>
      <Routes>
        <Route path="/" element={<Login />} />
      </Routes>
    </AuthProvider>
  </section>
)

export default App
