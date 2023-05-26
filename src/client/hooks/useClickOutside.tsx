import React, { useEffect } from 'react'

const useClickOutside = (
  ref: React.MutableRefObject<any>,
  display: React.Dispatch<React.SetStateAction<boolean>>
) => {
  const handleClickOutside = (event: any) => {
    if (!ref.current) return
    if (!ref.current.contains(event.target)) {
      display(false)
    }
  }
  useEffect(() => {
    document.addEventListener('click', handleClickOutside, true)
  }, [])
}

export default useClickOutside
