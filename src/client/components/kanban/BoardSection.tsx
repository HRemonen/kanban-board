import React from 'react'
import { DragDropContext } from 'react-beautiful-dnd'

const BoardSection = () => (
  <div className="px-6 text-black h-screen overflow-auto">
    <div className="overflow-y-hidden scrollbar-thin scrollbar-thumb-mainPurple scrollbar-track-transparent flex-1 p-4 space-x-4 flex">
      <p>Hello word</p>
    </div>
  </div>
)

export default BoardSection
