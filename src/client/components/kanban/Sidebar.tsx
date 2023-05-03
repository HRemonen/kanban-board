import React, { ReactNode } from 'react'
import { NavLink } from 'react-router-dom'
import { HiHome, HiTemplate } from 'react-icons/hi'
import LogoutButton from '../common/LogoutButton'

interface SidebarItemProps {
  to: string
  name: string
  icon: ReactNode
}

const SidebarItem = ({ to, name, icon }: SidebarItemProps) => (
  <li>
    <NavLink
      to={to}
      className="w-full relative inline-flex items-center justify-start px-4 py-2 overflow-hidden font-normal transition-all bg-white rounded hover:bg-white group"
    >
      {({ isActive }) => (
        <>
          <span
            className={
              isActive
                ? `w-48 h-48 rounded rotate-[-45deg] bg-purple-600 absolute bottom-0 left-0 ease-out duration-500 transition-all translate-y-full ml-0 mb-32 translate-x-0`
                : `w-48 h-48 rounded rotate-[-45deg] bg-purple-600 absolute bottom-0 left-0 -translate-x-full ease-out duration-500 transition-all translate-y-full mb-9 ml-9 group-hover:ml-0 group-hover:mb-32 group-hover:translate-x-0`
            }
          />
          <span
            className={`flex relative w-full text-left transition-colors duration-300 ease-in-out ${
              isActive ? 'text-white' : 'group-hover:text-white'
            }`}
          >
            <i className="mx-2">{icon}</i> {name}
          </span>
        </>
      )}
    </NavLink>
  </li>
)

const Sidebar = () => (
  <section className="text-gray-700 fixed inset-y-0 left-0 border-r border-[#635FC7] shadow-md max-h-screen w-60">
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
