import { createContext, useMemo, useState } from 'react'
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

export const AuthProvider = ({ children }: any) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [user, setUser] = useState<PrivateUser | null>(null)
  const [token, setToken] = useState('')

  const login = (token: string, user: PrivateUser) => {
    setIsAuthenticated(true)
    setUser(user)
    setToken(token)
  }

  const logout = () => {
    setIsAuthenticated(false)
    setUser(null)
    setToken('')
  }

  const contextValues = useMemo(
    () => ({ isAuthenticated, user, token, login, logout }),
    []
  )

  return (
    <AuthContext.Provider value={contextValues}>
      {children}
    </AuthContext.Provider>
  )
}
