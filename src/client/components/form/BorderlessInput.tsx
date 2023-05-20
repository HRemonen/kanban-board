import { InputType } from '../../types'

const BorderlessInput = ({ register, error, id, ...inputProps }: InputType) => (
  <div>
    <input
      className={`w-full resize-none bg-transparent text-sm font-light outline outline-0 transition-all focus:outline-0
        ${error ? 'border-red-500 text-[#EA5555]' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className='mt-2 text-sm text-[#EA5555]'>{error.message}</p>}
  </div>
)

export default BorderlessInput
