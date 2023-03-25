import React, { useEffect, useState } from 'react'
import { getUsers } from './services/userService'
import Login from './components/authentication/Login'

const App = () => {
  const [users, setUsers] = useState<any>([])

  useEffect(() => {
    getUsers()
      .then((response) => setUsers(response.data))
      .catch((error) => console.log(error.toJSON()))
  }, [])

  console.log(users)

  return (
    <section className="bg-slate-900 text-white">
      <Login />
    </section>
  )
}

export default App
