import React, { useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { HiLogout } from 'react-icons/hi'
import { AuthContext } from '../../contexts/AuthContext'

const LogoutButton = () => {
  const { logout } = useContext(AuthContext)
  const navigate = useNavigate()

  const onLogout = () => {
    logout()
    navigate('/')
  }

  return (
    <div className="p-4">
      <button
        type="button"
        className="inline-flex items-center justify-center h-10 pl-14 text-black hover:text-[#EA5555] text-sm font-semibold transition"
        onClick={onLogout}
      >
        <HiLogout size={20} />
        <span className="font-bold text-sm ml-2">Logout</span>
      </button>
    </div>
  )
}

export default LogoutButton
