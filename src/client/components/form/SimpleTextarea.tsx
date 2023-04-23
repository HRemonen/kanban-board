import { TextareaType } from '../../types'

const SimpleTextarea = ({
  register,
  error,
  label,
  id,
  ...inputProps
}: TextareaType) => (
  <div className="mb-8">
    <label
      className="block mb-2 text-sm font-medium text-gray-900"
      htmlFor={id}
    >
      {label}
    </label>
    <textarea
      className={`block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500
        ${error ? 'text-[#EA5555] border-red-500' : 'border-[#f4f7fd]'}`}
      rows={10}
      id={id}
      {...register(id)}
      {...inputProps}
    />
    {error && <p className="text-[#EA5555] text-sm mt-2">{error.message}</p>}
  </div>
)

export default SimpleTextarea
