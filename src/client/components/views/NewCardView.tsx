import React from 'react'
import { useForm } from 'react-hook-form'
import { useCreateNewCard } from '../../services/cardService'
import SimpleInput from '../form/SimpleInput'
import SimpleTextarea from '../form/SimpleTextarea'
import { List, NewCard } from '../../types'

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
      <button
        type="button"
        className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 absolute top-2.5 right-2.5 inline-flex items-center"
        onClick={() => setShowModal(false)}
      >
        <svg
          aria-hidden="true"
          className="w-5 h-5"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fillRule="evenodd"
            d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
            clipRule="evenodd"
          />
        </svg>
        <span className="sr-only">Close menu</span>
      </button>
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
                <div className="mt-2 px-2 hover:bg-green-300 hover:rounded-lg">
                  <button
                    type="submit"
                    className="inline-flex items-center py-2 text-green-500 text-sm font-medium"
                  >
                    <span className="mr-2 inline-block align-text-bottom select-text overflow-visible">
                      <svg
                        aria-hidden="true"
                        focusable="false"
                        role="img"
                        viewBox="0 0 16 16"
                        width="16"
                        height="16"
                        fill="currentColor"
                      >
                        <path d="M0 2.75C0 1.784.784 1 1.75 1h12.5c.966 0 1.75.784 1.75 1.75v1.5A1.75 1.75 0 0 1 14.25 6H1.75A1.75 1.75 0 0 1 0 4.25ZM1.75 7a.75.75 0 0 1 .75.75v5.5c0 .138.112.25.25.25h10.5a.25.25 0 0 0 .25-.25v-5.5a.75.75 0 0 1 1.5 0v5.5A1.75 1.75 0 0 1 13.25 15H2.75A1.75 1.75 0 0 1 1 13.25v-5.5A.75.75 0 0 1 1.75 7Zm0-4.5a.25.25 0 0 0-.25.25v1.5c0 .138.112.25.25.25h12.5a.25.25 0 0 0 .25-.25v-1.5a.25.25 0 0 0-.25-.25ZM6.25 8h3.5a.75.75 0 0 1 0 1.5h-3.5a.75.75 0 0 1 0-1.5Z" />
                      </svg>
                    </span>
                    Save changes
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </form>
    </div>
  )
}

export default NewCardView
