import { InputType } from '../../types'

const BorderlessInput = ({ register, error, id, ...inputProps }: InputType) => (
  <div>
    <input
      className='w-full resize-none bg-transparent text-sm font-light outline outline-0 transition-all focus:outline-0'
      id={id}
      {...register(id)}
      {...inputProps}
    />
  </div>
)

export default BorderlessInput
