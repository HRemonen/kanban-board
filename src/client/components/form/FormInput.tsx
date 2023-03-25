import { InputType } from '../../types'

const FormInput = ({
  register,
  error,
  label,
  id,
  ...inputProps
}: InputType) => (
  <div className="mb-8">
    <label
      className="block mb-2 font-semibold text-gray-900 md:text-xl"
      htmlFor={id}
    >
      {label}
    </label>
    <input
      className={`block w-full bg-transparent outline-none border-b-2 py-2 px-4 placeholder-gray-500 
        ${error ? 'text-red-400 border-red-500' : 'border-gray-400'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className="text-red-500 text-sm mt-2">{error.message}</p>}
  </div>
)

export default FormInput
