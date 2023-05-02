import React, { ReactNode } from 'react'
import { Link } from 'react-router-dom'
import { HiHome, HiTemplate } from 'react-icons/hi'
import LogoutButton from '../common/LogoutButton'

interface SidebarItemProps {
  to: string
  name: string
  icon: ReactNode
}

const SidebarItem = ({ to, name, icon }: SidebarItemProps) => (
  <li>
    <Link
      to={to}
      className="flex items-center bg-white rounded-r-3xl font-bold text-sm text-gray-700 py-3 px-4"
    >
      {icon}
      <span className="pl-4">{name}</span>
    </Link>
  </li>
)

const Sidebar = () => (
  <section className="fixed inset-y-0 left-0 border-r bg-[#A8A4FF] border-[#635FC7] shadow-md max-h-screen w-60">
    <div className="flex flex-col justify-between h-full">
      <div className="flex-grow">
        <div className="px-4 py-6 text-center">
          <h1 className="text-3xl font-light leading-none">Kanri</h1>
        </div>
        <div className="pr-4">
          <ul className="space-y-4">
            <SidebarItem to="/" name="Home" icon={<HiHome size={20} />} />
            <SidebarItem
              to="/boards"
              name="Boards"
              icon={<HiTemplate size={20} />}
            />
          </ul>
        </div>
      </div>
      <LogoutButton />
    </div>
  </section>
)

export default Sidebar
