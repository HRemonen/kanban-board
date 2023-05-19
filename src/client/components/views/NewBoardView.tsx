import React from 'react'
import { useNavigate } from 'react-router-dom'
import { useForm } from 'react-hook-form'

import { useCreateBoard } from '../../services/boardService'

import SaveButton from '../common/SaveButton'
import SimpleTextarea from '../form/SimpleTextarea'
import SimpleInput from '../form/SimpleInput'

import { NewBoard } from '../../types'

const NewBoardView = () => {
  const mutateBoard = useCreateBoard()
  const navigate = useNavigate()

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<NewBoard>()

  const onSubmit = (data: NewBoard) => {
    mutateBoard.mutateAsync({
      board: data,
    })
    navigate('/boards')
  }

  return (
    <div className="px-6 text-black h-screen overflow-auto">
      <div className="w-full px-4 mx-auto">
        <div className="relative min-w-0 break-words bg-white w-full p-8 mb-6 mt-16">
          <div className="flex flex-col justify-between">
            <h1 className="text-3xl font-bold">Create a New Board</h1>
            <form
              onSubmit={handleSubmit(onSubmit)}
              className="flex flex-col justify-center text-left w-full mt-8"
            >
              <div className="grid grid-cols-12 gap-2">
                <div className="col-span-12 border-b-2">
                  <SimpleInput
                    id="name"
                    type="name"
                    placeholder="Input board name here"
                    name="name"
                    label="Name"
                    register={register}
                    error={errors.name}
                  />
                </div>
                <div className="col-span-7 my-4">
                  <SimpleTextarea
                    id="description"
                    type="description"
                    placeholder="Input board description here"
                    name="description"
                    label="Description"
                    register={register}
                    error={errors.description}
                  />
                </div>
                <div className="col-span-5 row-span-6 border-l-2">
                  <div className="my-4 mx-2 grid grid-cols-2 gap-4">
                    <div className="font-normal text-sm text-gray-700">
                      Team members
                    </div>
                    <div className="font-light text-sm text-gray-400">
                      Add members
                    </div>
                    <div className="my-4 border-t-2 col-span-2">
                      <SaveButton />
                    </div>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  )
}

export default NewBoardView
