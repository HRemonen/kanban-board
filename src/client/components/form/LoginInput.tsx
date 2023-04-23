import { InputType } from '../../types'

const LoginInput = ({
  register,
  error,
  label,
  id,
  ...inputProps
}: InputType) => (
  <div className="mb-8">
    <label
      className="block mb-2 font-semibold text-[#f4f7fd] md:text-xl"
      htmlFor={id}
    >
      {label}
    </label>
    <input
      className={`block w-full bg-transparent outline-none border-b-2 py-2 px-4 placeholder-gray-700 
        ${error ? 'text-[#EA5555] border-red-500' : 'border-[#f4f7fd]'}`}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className="text-[#EA5555] text-sm mt-2">{error.message}</p>}
  </div>
)

export default LoginInput
