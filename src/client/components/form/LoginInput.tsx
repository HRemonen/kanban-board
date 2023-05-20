import { InputType } from '../../types'

const LoginInput = ({
  register,
  error,
  label,
  id,
  ...inputProps
}: InputType) => (
  <div className='mb-8'>
    <label
      className='mb-2 block font-semibold text-[#f4f7fd] md:text-xl'
      htmlFor={id}
    >
      {label}
    </label>
    <input
      className={`block w-full border-b-2 bg-transparent px-4 py-2 placeholder-gray-700 outline-none 
        ${error ? 'border-red-500 text-[#EA5555]' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className='mt-2 text-sm text-[#EA5555]'>{error.message}</p>}
  </div>
)

export default LoginInput
