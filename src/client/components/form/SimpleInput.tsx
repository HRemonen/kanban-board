import { InputType } from '../../types'

const SimpleInput = ({
  register,
  error,
  label,
  id,
  ...inputProps
}: InputType) => (
  <div className="mb-8">
    <label
      className="block mb-2 text-sm font-medium text-gray-900"
      htmlFor={id}
    >
      {label}
    </label>
    <input
      className={`bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5
        ${error ? 'text-[#EA5555] border-red-500' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className="text-[#EA5555] text-sm mt-2">{error.message}</p>}
  </div>
)

export default SimpleInput
