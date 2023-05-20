import React from 'react'

const SaveButton = () => (
  <div className='px-2 hover:rounded-lg hover:bg-green-300'>
    <button
      type='submit'
      className='inline-flex items-center py-2 text-sm font-medium text-green-500'
    >
      <span className='mr-2 inline-block select-text overflow-visible align-text-bottom'>
        <svg
          aria-hidden='true'
          aria-label='Save changes button'
          data-cy='save-button'
          focusable='false'
          role='img'
          viewBox='0 0 16 16'
          width='16'
          height='16'
          fill='currentColor'
        >
          <path d='M0 2.75C0 1.784.784 1 1.75 1h12.5c.966 0 1.75.784 1.75 1.75v1.5A1.75 1.75 0 0 1 14.25 6H1.75A1.75 1.75 0 0 1 0 4.25ZM1.75 7a.75.75 0 0 1 .75.75v5.5c0 .138.112.25.25.25h10.5a.25.25 0 0 0 .25-.25v-5.5a.75.75 0 0 1 1.5 0v5.5A1.75 1.75 0 0 1 13.25 15H2.75A1.75 1.75 0 0 1 1 13.25v-5.5A.75.75 0 0 1 1.75 7Zm0-4.5a.25.25 0 0 0-.25.25v1.5c0 .138.112.25.25.25h12.5a.25.25 0 0 0 .25-.25v-1.5a.25.25 0 0 0-.25-.25ZM6.25 8h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1 0-1.5Z' />
        </svg>
      </span>
      Save changes
    </button>
  </div>
)

export default SaveButton
