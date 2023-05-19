import { createContext, useContext, useEffect, useMemo, useState } from 'react'
import { PrivateUser } from '../types'

interface Config {
  headers: {
    Authorization: string
  }
}

interface AuthContextType {
  user: PrivateUser | null
  config: Config
  isAuthenticated: boolean
  login: (token: string, user: PrivateUser) => void
  logout: () => void
}

export const AuthContext = createContext<AuthContextType>({
  user: null,
  config: {
    headers: { Authorization: '' },
  },
  isAuthenticated: false,
  login: () => {},
  logout: () => {},
})

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [user, setUser] = useState<PrivateUser | null>(null)
  const [config, setConfig] = useState<Config>({
    headers: { Authorization: '' },
  })

  useEffect(() => {
    const userToken = sessionStorage.getItem('token')
    const loggedUser = sessionStorage.getItem('user')

    if (loggedUser && userToken) {
      const foundUser = JSON.parse(loggedUser)
      setUser(foundUser)
      setConfig({
        headers: { Authorization: userToken },
      })
    }
  }, [])

  const login = (userToken: string, loggedUser: PrivateUser) => {
    setIsAuthenticated(true)
    setUser(loggedUser)
    setConfig({
      headers: { Authorization: `bearer ${userToken}` },
    })
    sessionStorage.setItem('user', JSON.stringify(loggedUser))
  }

  const logout = () => {
    setIsAuthenticated(false)
    setUser(null)
    setConfig({
      headers: { Authorization: '' },
    })
    sessionStorage.clear()
  }

  const contextValues = useMemo(
    () => ({ isAuthenticated, user, config, login, logout }),
    [login, logout]
  )

  return (
    <AuthContext.Provider value={contextValues}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuthenticatedUser = () => {
  const { user, config } = useContext(AuthContext)
  return { user, config }
}
