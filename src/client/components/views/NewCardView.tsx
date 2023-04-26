import React from 'react'
import { useForm } from 'react-hook-form'
import { useCreateNewCard } from '../../services/cardService'
import SimpleInput from '../form/SimpleInput'
import SimpleTextarea from '../form/SimpleTextarea'
import { List, NewCard } from '../../types'
import SaveButton from '../common/SaveButton'
import CloseMenu from '../common/CloseMenu'

type ModalProps = {
  list: List
  setShowModal: React.Dispatch<React.SetStateAction<boolean>>
}

const NewCardView = ({ list, setShowModal }: ModalProps) => {
  const mutateCard = useCreateNewCard()

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<NewCard>()

  const onSubmit = (data: NewCard) => {
    mutateCard.mutateAsync({
      listID: list.ID,
      card: data,
    })

    setShowModal(false)
  }

  return (
    <div className="fixed top-0 right-0 z-40 h-screen p-4 bg-white w-[60%] border-l-2">
      <h5
        id="drawer-label"
        className="inline-flex items-center mb-6 text-base font-semibold text-gray-500 uppercase"
      >
        <svg
          className="w-5 h-5 mr-2"
          aria-hidden="true"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fillRule="evenodd"
            d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
            clipRule="evenodd"
          />
        </svg>
        New card form
      </h5>
      <CloseMenu onClick={() => setShowModal(false)} />
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col justify-center text-left w-[80%]"
      >
        <div className="grid grid-cols-12 gap-2">
          <div className="col-span-12 border-b-2">
            <SimpleInput
              id="title"
              type="title"
              placeholder="Input title here"
              name="title"
              label="Title"
              register={register}
              error={errors.title}
            />
          </div>
          <div className="col-span-7">
            <SimpleTextarea
              id="description"
              type="description"
              placeholder="Input description here"
              name="description"
              label="Description"
              register={register}
              error={errors.description}
            />
          </div>
          <div className="col-span-5 row-span-6 border-l-2">
            <div className="my-4 mx-2 grid grid-cols-2 gap-4">
              <div className="font-normal text-sm text-gray-700">Assignees</div>
              <div className="font-light text-sm text-gray-400">
                Add assignees
              </div>
              <div className="font-normal text-sm text-gray-700">Labels</div>
              <div className="font-light text-sm text-gray-400">Add labels</div>
              <div className="font-normal text-sm text-gray-700">Status</div>
              <div className="font-light text-sm text-gray-400">
                <button
                  type="button"
                  className="px-4 bg-gray-100 hover:bg-gray-200 text-gray-800 text-sm font-medium rounded-full"
                >
                  Add status
                </button>
              </div>
              <div className="my-4 border-t-2 col-span-2">
                <SaveButton />
              </div>
            </div>
          </div>
        </div>
      </form>
    </div>
  )
}

export default NewCardView
