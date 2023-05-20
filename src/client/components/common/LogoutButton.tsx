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
    <div className='p-4'>
      <button
        type='button'
        className='inline-flex h-10 items-center justify-center pl-14 text-sm font-semibold text-black transition hover:text-[#EA5555]'
        onClick={onLogout}
      >
        <HiLogout size={20} />
        <span className='ml-2 text-sm font-bold'>Logout</span>
      </button>
    </div>
  )
}

export default LogoutButton
