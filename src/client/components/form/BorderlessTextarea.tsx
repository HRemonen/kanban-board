import { TextareaType } from '../../types'

const BorderlessTextarea = ({
  register,
  error,
  id,
  ...inputProps
}: TextareaType) => (
  <div className='mb-8'>
    <textarea
      className={`h-[70vh] w-full resize-none bg-transparent pb-1.5 pt-4 text-sm font-light outline outline-0 transition-all focus:outline-0
        ${error ? 'border-red-500 text-[#EA5555]' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className='mt-2 text-sm text-[#EA5555]'>{error.message}</p>}
  </div>
)

export default BorderlessTextarea
