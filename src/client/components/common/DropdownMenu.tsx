import React, { useState } from 'react'

const DropdownMenu = ({ handleDelete }: { handleDelete: () => void }) => {
  const [showDropdown, setShowDropdown] = useState(false)

  const onDelete = () => {
    setShowDropdown(false)
    handleDelete()
  }

  return (
    <div className='relative'>
      <button
        type='button'
        className='inline-flex items-center rounded-lg bg-white p-2 text-center text-sm font-medium text-gray-900 hover:bg-gray-200 focus:outline-none focus:ring-4 focus:ring-gray-50'
        onClick={() => setShowDropdown(!showDropdown)}
      >
        <span className='inline-block select-text overflow-visible align-text-bottom'>
          <svg
            aria-hidden='true'
            focusable='false'
            role='img'
            viewBox='0 0 16 16'
            width='16'
            height='16'
            fill='currentColor'
          >
            <path d='M8 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3ZM1.5 9a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Zm13 0a1.5 1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3Z' />
          </svg>
        </span>
      </button>
      {showDropdown && (
        <div className='w-30 absolute z-10 divide-y divide-gray-200 rounded-lg bg-white shadow'>
          <ul
            className='pt-2 text-sm text-gray-700'
            aria-labelledby='dropdownMenuIconButton'
          >
            <li className='rounded-lg hover:bg-gray-100'>
              <button type='button' className='block px-4 py-2'>
                Edit
              </button>
            </li>
          </ul>
          <div className='mt-2'>
            <button
              type='button'
              className='block rounded-lg px-4 py-2 text-sm text-red-500 hover:bg-red-100'
              onClick={onDelete}
            >
              Delete
            </button>
          </div>
        </div>
      )}
    </div>
  )
}

export default DropdownMenu
