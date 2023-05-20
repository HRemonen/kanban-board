import React from 'react'
import NewsSection from './NewsSection'
import StatsSection from './StatsSection'

const HomeSection = () => (
  <div className='h-screen overflow-auto px-6 text-black'>
    <div className='mx-auto w-full px-4'>
      <div className='relative mb-6 mt-16 flex w-full min-w-0 flex-col break-words bg-white p-8'>
        <h1 className='mb-10 text-3xl font-bold'>
          A Kanban board is a visual tool used to manage and track work in
          progress.
        </h1>

        <div className='my-10'>
          <div className='grid grid-cols-2 gap-x-20'>
            <StatsSection />
            <NewsSection />
          </div>
        </div>
      </div>
    </div>
  </div>
)

export default HomeSection
