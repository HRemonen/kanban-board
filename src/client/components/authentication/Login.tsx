import { useForm } from 'react-hook-form'
import { Link } from 'react-router-dom'

import FormInput from '../form/FormInput'
import { LoginUser } from '../../types'

import login_illustration_image from '../../illustrations/login_illustration_image.png'

const Login = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginUser>({
    mode: 'onBlur',
  })

  const onSubmit = async ({ email, password }: LoginUser) => {
    console.log(email, password)
  }

  return (
    <section className="md:grid md:grid-cols-2 text-center">
      <div className="flex flex-col h-screen justify-center text-center items-center p-12 border-r-2 border-solid border-gray-300 md:shadow-lg">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col justify-center text-left w-[80%]"
        >
          <FormInput
            id="email"
            type="email"
            placeholder="hello@world.com"
            name="email"
            label="Email"
            register={register}
            error={errors.email}
          />

          <FormInput
            id="password"
            type="password"
            placeholder="Password"
            name="password"
            label="Password"
            register={register}
            error={errors.password}
          />

          <button
            id="login-button"
            data-cy="login-form-button"
            type="submit"
            className="inline-block text-white bg-gradient-to-br from-green-300 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2"
          >
            Login
          </button>
        </form>

        <p className="mt-4">
          Don&lsquo;t have an account?
          <Link
            to="/register"
            className="ml-2 text-blue-600 inline-flex items-center font-medium hover:underline"
          >
            Register here
          </Link>
        </p>
      </div>

      <div className="hidden md:flex flex-col justify-center ">
        <img className="scale-50" src={login_illustration_image} alt="" />
      </div>
    </section>
  )
}

export default Login
