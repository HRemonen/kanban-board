import React, { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { useCreateNewList } from '../../services/listService'

import BorderlessTextarea from '../form/BorderlessTextarea'
import SaveButton from '../common/SaveButton'
import CloseMenu from '../common/CloseMenu'

import { Board } from '../../types'
import { NewList, NewListZod } from '../../validators/validators'

type CreateListProps = {
  setShowCreateList: React.Dispatch<React.SetStateAction<boolean>>
}

const AddListButton = ({ setShowCreateList }: CreateListProps) => (
  <div className='mt-[84px]'>
    <button
      type='button'
      aria-label='Add a new list to the board'
      data-cy='add-new-list-button'
      className='inline-block select-none overflow-visible rounded-lg border border-gray-200 p-3 align-text-bottom shadow '
      onClick={() => setShowCreateList(true)}
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
  </div>
)

const NewListView = ({ board }: { board: Board }) => {
  const mutateList = useCreateNewList()
  const [showCreateList, setShowCreateList] = useState(false)

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<NewList>({
    shouldUnregister: true,
    resolver: zodResolver(NewListZod),
    defaultValues: {
      name: '',
    },
  })

  const onSubmit = (data: NewList) => {
    mutateList.mutateAsync({
      boardID: board.ID,
      list: data,
    })

    setShowCreateList(false)
  }

  if (!showCreateList) {
    return <AddListButton setShowCreateList={setShowCreateList} />
  }

  return (
    <div>
      <div
        data-cy='new-list-form'
        className={`relative mt-[84px] h-[300px] w-[280px] shrink-0 rounded-lg border border-gray-200 p-6 shadow
        ${errors.name ? 'border-red-500 text-[#EA5555]' : 'border-[#f4f7fd]'}`}
      >
        <form
          onSubmit={handleSubmit(onSubmit)}
          className='flex flex-col text-left'
        >
          <BorderlessTextarea
            id='name'
            type='name'
            name='name'
            label='Add list'
            placeholder='Input list name'
            register={register}
            error={errors.name}
          />
          <div className='absolute bottom-2'>
            <SaveButton />
          </div>
        </form>
        <CloseMenu onClick={() => setShowCreateList(false)} />
      </div>
      <div>
        {errors.name && (
          <p className='mt-2 text-sm text-[#EA5555]'>{errors.name.message}</p>
        )}
      </div>
    </div>
  )
}

export default NewListView
