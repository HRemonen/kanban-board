import { InputType } from '../../types'

const BorderlessInput = ({ register, error, id, ...inputProps }: InputType) => (
  <div>
    <input
      className={`w-full resize-none bg-transparent text-sm font-light outline outline-0 transition-all focus:outline-0
        ${error ? 'text-[#EA5555] border-red-500' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className="text-[#EA5555] text-sm mt-2">{error.message}</p>}
  </div>
)

export default BorderlessInput
