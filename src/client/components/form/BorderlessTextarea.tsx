import { TextareaType } from '../../types'

const BorderlessTextarea = ({
  register,
  error,
  id,
  ...inputProps
}: TextareaType) => (
  <div className="mb-8">
    <textarea
      className={`h-[70vh] w-full resize-none bg-transparent pt-4 pb-1.5 text-sm font-light outline outline-0 transition-all focus:outline-0
        ${error ? 'text-[#EA5555] border-red-500' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className="text-[#EA5555] text-sm mt-2">{error.message}</p>}
  </div>
)

export default BorderlessTextarea
