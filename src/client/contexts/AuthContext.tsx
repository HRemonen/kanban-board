import { createContext, useContext, useEffect, useMemo, useState } from 'react'
import { PrivateUser } from '../types'

interface AuthContextType {
  user: PrivateUser | null
  token: string
  isAuthenticated: boolean
  login: (token: string, user: PrivateUser) => void
  logout: () => void
}

export const AuthContext = createContext<AuthContextType>({
  user: null,
  token: '',
  isAuthenticated: false,
  login: () => {},
  logout: () => {},
})

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [user, setUser] = useState<PrivateUser | null>(null)
  const [token, setToken] = useState('')

  useEffect(() => {
    const loggedInUser = localStorage.getItem('user')
    if (loggedInUser) {
      const foundUser = JSON.parse(loggedInUser)
      setUser(foundUser)
    }
  }, [])

  const login = (token: string, user: PrivateUser) => {
    setIsAuthenticated(true)
    setUser(user)
    setToken(token)
    localStorage.setItem('user', JSON.stringify(user))
  }

  const logout = () => {
    setIsAuthenticated(false)
    setUser(null)
    setToken('')
    localStorage.clear()
  }

  const contextValues = useMemo(
    () => ({ isAuthenticated, user, token, login, logout }),
    [login, logout]
  )

  return (
    <AuthContext.Provider value={contextValues}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuthenticatedUser = () => {
  const { user, token } = useContext(AuthContext)
  return { user, token }
}
