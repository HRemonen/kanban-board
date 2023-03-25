import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Login from './components/authentication/Login'

const App = () => (
  <section className="bg-slate-900 text-white">
    <Routes>
      <Route path="/" element={<Login />} />
    </Routes>
  </section>
)

export default App
