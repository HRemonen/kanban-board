import { useEffect, useState } from 'react'
import { useAuthenticatedUser } from '../contexts/AuthContext'

import { getUserByID } from '../services/userService'

import { PrivateUser } from '../types'

const useFetchUser = () => {
  const { user } = useAuthenticatedUser()
  const [userData, setUserData] = useState<PrivateUser>()

  useEffect(() => {
    if (!user) return
    const fetchUser = async (user: PrivateUser) => {
      const response = await getUserByID(user.id)
      if ('role' in response.data) setUserData(response.data)
    }

    fetchUser(user)
  }, [user])

  return userData
}

export default useFetchUser
