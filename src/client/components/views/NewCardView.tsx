import React from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import BorderlessInput from '../form/BorderlessInput'

import { NewCardZod, NewCard } from '../../validators/validators'

type NewCardProps = {
  onSubmit: (data: NewCard) => void
}

const NewCardView = ({ onSubmit }: NewCardProps) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<NewCard>({
    resolver: zodResolver(NewCardZod),
    defaultValues: {
      title: '',
    },
  })

  return (
    <div className='mt-2'>
      <div>
        {errors.title && (
          <p className='mt-2 text-sm text-[#EA5555]'>{errors.title.message}</p>
        )}
      </div>
      <div
        className={`shadow' max-w-sm rounded-lg border-2 border-dashed border-gray-200 bg-white py-3 pl-6
        ${errors.title ? 'border-red-500 text-[#EA5555]' : 'border-[#f4f7fd]'}`}
      >
        <form
          onSubmit={handleSubmit(onSubmit)}
          className='flex justify-between'
        >
          <BorderlessInput
            id='title'
            type='title'
            placeholder='Input title here'
            name='title'
            register={register}
          />

          <button
            type='submit'
            className='mr-4 inline-flex items-center text-sm font-medium text-gray-600'
            data-cy='save-new-card-button'
          >
            <svg
              aria-hidden='true'
              focusable='false'
              role='img'
              viewBox='0 0 16 16'
              width='16'
              height='16'
              fill='currentColor'
            >
              <path d='M7.75 2a.75.75 0 0 1 .75.75V7h4.25a.75.75 0 0 1 0 1.5H8.5v4.25a.75.75 0 0 1-1.5 0V8.5H2.75a.75.75 0 0 1 0-1.5H7V2.75A.75.75 0 0 1 7.75 2Z' />
            </svg>
          </button>
        </form>
      </div>
    </div>
  )
}

export default NewCardView
