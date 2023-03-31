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
                  className="w-40 h-40 p-1 rounded-full ring-2 ring-gray-300 dark:ring-gray-500"
                  src="/docs/images/people/profile-picture-5.jpg"
                  alt="Bordered avatar"
                />
              </div>
            </div>
            <div className="w-full px-4 text-center mt-20">
              <div className="flex justify-center py-4 lg:pt-4 pt-8">
                <div className="p-3 text-center">
                  <span className="text-xl font-bold block uppercase tracking-wide text-gray-600">
                    20
                  </span>
                  <span className="text-sm text-gray-400">Boards</span>
                </div>
                <div className="p-3 text-center">
                  <span className="text-xl font-bold block uppercase tracking-wide text-gray-600">
                    10
                  </span>
                  <span className="text-sm text-gray-400">Cards</span>
                </div>
                <div className="p-3 text-center">
                  <span className="text-xl font-bold block uppercase tracking-wide text-gray-600">
                    0
                  </span>
                  <span className="text-sm text-gray-400">Teams</span>
                </div>
              </div>
            </div>
          </div>
          <div className="text-center mt-12">
            <h3 className="text-xl font-semibold leading-normal mb-2 text-gray-700 mb-2">
              Username
            </h3>
            <div className="text-sm leading-normal mt-0 mb-2 text-gray-400 font-bold uppercase">
              <p className="fas fa-map-marker-alt mr-2 text-lg text-blueGray-400">
                EMAIL
              </p>
            </div>
          </div>
          <div className="mt-10 py-10 border-t border-gray-200 text-center">
            <div className="flex flex-wrap justify-center">
              <div className="w-full lg:w-9/12 px-4">
                <p className="mb-4 text-lg leading-relaxed text-gray-700">
                  Something diipa daapa information
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
