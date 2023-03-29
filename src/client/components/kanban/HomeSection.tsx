import React from 'react'
import NewsSection from './NewsSection'
import StatsSection from './StatsSection'

const HomeSection = () => (
  <div className="px-6 text-black h-screen overflow-auto">
    <div className="w-full px-4 mx-auto">
      <div className="relative flex flex-col min-w-0 break-words bg-white w-full p-8 mb-6 mt-16">
        <h1 className="text-3xl font-bold mb-10">
          A Kanban board is a visual tool used to manage and track work in
          progress.
        </h1>

        <div className="my-10">
          <div className="grid grid-cols-2 gap-x-20">
            <StatsSection />
            <NewsSection />
          </div>
        </div>
      </div>
    </div>
  </div>
)

export default HomeSection
