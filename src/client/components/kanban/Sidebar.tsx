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
      className='group relative inline-flex w-full items-center justify-start overflow-hidden rounded bg-white px-4 py-2 font-normal transition-all hover:bg-white'
    >
      {({ isActive }) => (
        <>
          <span
            className={
              isActive
                ? `absolute bottom-0 left-0 mb-32 ml-0 h-48 w-48 translate-x-0 translate-y-full rotate-[-45deg] rounded bg-purple-600 transition-all duration-500 ease-out`
                : `absolute bottom-0 left-0 mb-9 ml-9 h-48 w-48 -translate-x-full translate-y-full rotate-[-45deg] rounded bg-purple-600 transition-all duration-500 ease-out group-hover:mb-32 group-hover:ml-0 group-hover:translate-x-0`
            }
          />
          <span
            className={`relative flex w-full text-left transition-colors duration-300 ease-in-out ${
              isActive ? 'text-white' : 'group-hover:text-white'
            }`}
          >
            <i className='mx-2'>{icon}</i> {name}
          </span>
        </>
      )}
    </NavLink>
  </li>
)

const Sidebar = () => (
  <section className='fixed inset-y-0 left-0 max-h-screen w-60 border-r border-[#635FC7] text-gray-700 shadow-md'>
    <div className='flex h-full flex-col justify-between'>
      <div className='flex-grow'>
        <div className='px-4 py-6 text-center'>
          <h1 className='text-3xl font-light leading-none'>Kanri</h1>
        </div>
        <div className='pr-4'>
          <ul className='space-y-4'>
            <SidebarItem to='/' name='Home' icon={<HiHome size={20} />} />
            <SidebarItem
              to='/boards'
              name='Boards'
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
