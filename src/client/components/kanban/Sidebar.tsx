import React, { ReactNode } from 'react'
import { Link } from 'react-router-dom'
import { HiHome, HiTemplate, HiUserCircle } from 'react-icons/hi'
import Logout from '../common/Logout'

interface SidebarItemProps {
  to: string
  name: string
  icon: ReactNode
}

const SidebarItem = ({ to, name, icon }: SidebarItemProps) => (
  <li>
    <Link
      to={to}
      className="flex items-center bg-white rounded-xl font-bold text-sm text-gray-700 py-3 px-4"
    >
      {icon}
      <span className="pl-2">{name}</span>
    </Link>
  </li>
)

const Sidebar = () => (
  <section className="fixed inset-y-0 left-0 bg-slate-900 shadow-md max-h-screen w-60">
    <div className="flex flex-col justify-between h-full">
      <div className="flex-grow">
        <div className="px-4 py-6 text-center border-b">
          <h1 className="text-xl font-bold leading-none">Kanri</h1>
        </div>
        <div className="p-4">
          <ul className="space-y-2">
            <SidebarItem to="/" name="Home" icon={<HiHome size={20} />} />
            <SidebarItem
              to="/"
              name="Profile"
              icon={<HiUserCircle size={20} />}
            />
            <SidebarItem to="/" name="Boards" icon={<HiTemplate size={20} />} />
          </ul>
        </div>
      </div>
      <Logout />
    </div>
  </section>
)

export default Sidebar
