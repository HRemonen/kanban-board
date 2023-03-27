import React from 'react'
import { HiLogout } from 'react-icons/hi'

const Logout = () => (
  <div className="p-4">
    <button
      type="button"
      className="inline-flex items-center justify-center h-10 pl-14 text-white hover:text-red-400 text-sm font-semibold transition"
    >
      <HiLogout size={20} />
      <span className="font-bold text-sm ml-2">Logout</span>
    </button>
  </div>
)

export default Logout
