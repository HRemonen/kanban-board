import React, { useState } from 'react'
import { useForm } from 'react-hook-form'

import SimpleInput from '../form/SimpleInput'
import { NewList } from '../../types'
import SaveButton from '../common/SaveButton'

type CreateListProps = {
  setShowCreateList: React.Dispatch<React.SetStateAction<boolean>>
}

const AddListButton = ({ setShowCreateList }: CreateListProps) => (
  <div className="mt-[76px]">
    <button
      type="button"
      aria-label="Add a new list to the board"
      data-cy="add-new-list-button"
      className="p-3 inline-block align-text-bottom select-none overflow-visible border border-gray-200 rounded-lg shadow "
      onClick={() => setShowCreateList(true)}
    >
      <svg
        aria-hidden="true"
        focusable="false"
        role="img"
        viewBox="0 0 16 16"
        width="16"
        height="16"
        fill="currentColor"
      >
        <path d="M7.75 2a.75.75 0 0 1 .75.75V7h4.25a.75.75 0 0 1 0 1.5H8.5v4.25a.75.75 0 0 1-1.5 0V8.5H2.75a.75.75 0 0 1 0-1.5H7V2.75A.75.75 0 0 1 7.75 2Z" />
      </svg>
    </button>
  </div>
)

const NewListView = () => {
  const [showCreateList, setShowCreateList] = useState(false)

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<NewList>()

  const onSubmit = (data: NewList) => {
    console.log(data)

    setShowCreateList(false)
  }

  if (!showCreateList) {
    return <AddListButton setShowCreateList={setShowCreateList} />
  }

  return (
    <div
      data-cy="new-list-form"
      className="mt-[76px] w-[280px] shrink-0 p-6 border border-gray-200 rounded-lg shadow "
    >
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col text-left"
      >
        <div className="">
          <div className="">
            <SimpleInput
              id="name"
              type="name"
              placeholder="Input list name"
              name="name"
              label="List name"
              register={register}
              error={errors.name}
            />
          </div>
          <SaveButton />
        </div>
      </form>
    </div>
  )
}

export default NewListView
