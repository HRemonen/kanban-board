import React from 'react'
import NewsSection from './NewsSection'
import StatsSection from './StatsSection'

const HomeSection = () => (
  <section className="ml-60 pt-14 text-black h-screen overflow-auto">
    <div className="px-6 py-8">
      <div className="max-w-4xl mx-auto">
        <div className="bg-white rounded-3xl p-8 mb-5">
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
  </section>
)

export default HomeSection
