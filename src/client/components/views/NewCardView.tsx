import React, { useRef, useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import useClickOutside from '../../hooks/useClickOutside'

import { useCreateNewCard } from '../../services/cardService'

import BorderlessInput from '../form/BorderlessInput'

import { NewCardZod, NewCard } from '../../validators/validators'
import { List } from '../../types'

type CreateCardProps = {
  setShowCardBones: React.Dispatch<React.SetStateAction<boolean>>
}

const AddCardButton = ({ setShowCardBones }: CreateCardProps) => (
  <div className='mb-2 px-2 hover:rounded-lg hover:bg-gray-300'>
    <button
      type='button'
      className='inline-flex items-center py-2 text-sm font-medium text-gray-600'
      data-cy='add-card-button'
      onClick={() => setShowCardBones(true)}
    >
      <span className='mr-2 inline-block select-text overflow-visible align-text-bottom'>
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
      </span>
      Add card
    </button>
  </div>
)

const NewCardView = ({ list }: { list: List }) => {
  const newCardRef = useRef(null)
  const [showCardBones, setShowCardBones] = useState(false)
  const mutateCard = useCreateNewCard()

  useClickOutside(newCardRef, setShowCardBones)

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<NewCard>({
    shouldUnregister: true,
    resolver: zodResolver(NewCardZod),
    defaultValues: {
      title: '',
    },
  })

  const onSubmit = (data: NewCard) => {
    mutateCard.mutateAsync({
      listID: list.ID,
      card: data,
    })

    setShowCardBones(false)
  }

  if (!showCardBones) {
    return <AddCardButton setShowCardBones={setShowCardBones} />
  }

  return (
    <div className='mb-2 mt-[44px]'>
      <div
        ref={newCardRef}
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
      <div>
        {errors.title && (
          <p className='text-sm text-[#EA5555]'>{errors.title.message}</p>
        )}
      </div>
    </div>
  )
}

export default NewCardView
