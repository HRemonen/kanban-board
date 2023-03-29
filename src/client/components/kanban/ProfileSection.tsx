import React from 'react'

const ProfileSection = () => (
  <div className="px-6 text-black h-screen overflow-auto">
    <div className="w-full px-4 mx-auto">
      <div className="relative flex flex-col min-w-0 break-words bg-white w-full mb-6 mt-16">
        <div className="px-6">
          <div className="flex flex-wrap justify-center">
            <div className="w-full px-4 flex justify-center">
              <div className="relative">
                <img
                  alt="..."
                  src="default.png"
                  className="shadow-xl rounded-full h-auto align-middle border-none absolute -m-16 -ml-20 lg:-ml-16 max-w-150-px"
                />
              </div>
            </div>
            <div className="w-full px-4 text-center mt-20">
              <div className="flex justify-center py-4 lg:pt-4 pt-8">
                <div className="p-3 text-center">
                  <span className="text-xl font-bold block uppercase tracking-wide text-gray-600">
                    22
                  </span>
                  <span className="text-sm text-gray-400">Friends</span>
                </div>
                <div className="p-3 text-center">
                  <span className="text-xl font-bold block uppercase tracking-wide text-gray-600">
                    10
                  </span>
                  <span className="text-sm text-gray-400">Photos</span>
                </div>
                <div className="p-3 text-center">
                  <span className="text-xl font-bold block uppercase tracking-wide text-gray-600">
                    89
                  </span>
                  <span className="text-sm text-gray-400">Comments</span>
                </div>
              </div>
            </div>
          </div>
          <div className="text-center mt-12">
            <h3 className="text-xl font-semibold leading-normal mb-2 text-gray-700 mb-2">
              Jenna Stones
            </h3>
            <div className="text-sm leading-normal mt-0 mb-2 text-gray-400 font-bold uppercase">
              <p className="fas fa-map-marker-alt mr-2 text-lg text-blueGray-400">
                Los Angeles, California
              </p>
            </div>
            <div className="mb-2 text-gray-600 mt-10">
              <p className="fas fa-briefcase mr-2 text-lg text-gray-400">
                Solution Manager - Creative Tim Officer
              </p>
            </div>
            <div className="mb-2 text-gray-600">
              <p className="fas fa-university mr-2 text-lg text-gray-400">
                University of Computer Science
              </p>
            </div>
          </div>
          <div className="mt-10 py-10 border-t border-gray-200 text-center">
            <div className="flex flex-wrap justify-center">
              <div className="w-full lg:w-9/12 px-4">
                <p className="mb-4 text-lg leading-relaxed text-gray-700">
                  An artist of considerable range, Jenna the name taken by
                  Melbourne-raised, Brooklyn-based Nick Murphy writes, performs
                  and records all of his own music, giving it a warm, intimate
                  feel with a solid groove structure. An artist of considerable
                  range.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
)

export default ProfileSection
