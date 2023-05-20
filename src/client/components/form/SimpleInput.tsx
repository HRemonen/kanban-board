import { InputType } from '../../types'

const SimpleInput = ({
  register,
  error,
  label,
  id,
  ...inputProps
}: InputType) => (
  <div className='mb-8'>
    <label
      className='mb-2 block text-sm font-medium text-gray-900'
      htmlFor={id}
    >
      {label}
    </label>
    <input
      className={`block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500
        ${error ? 'border-red-500 text-[#EA5555]' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className='mt-2 text-sm text-[#EA5555]'>{error.message}</p>}
  </div>
)

export default SimpleInput
