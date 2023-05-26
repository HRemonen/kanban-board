import { TextareaType } from '../../types'

const BorderlessTextarea = ({
  register,
  id,
  label,
  ...inputProps
}: TextareaType) => (
  <div className='mb-8'>
    <h3 className='text-lg font-semibold'>{label}</h3>
    <textarea
      className='h-[200px] w-full resize-none bg-transparent pb-1.5 pt-4 text-sm font-light outline outline-0 transition-all focus:outline-0'
      id={id}
      {...register(id)}
      {...inputProps}
    />
  </div>
)

export default BorderlessTextarea
