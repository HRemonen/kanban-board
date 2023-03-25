import { useContext } from 'react'
import axios, { AxiosError } from 'axios'
import { useForm } from 'react-hook-form'
import { Link } from 'react-router-dom'

import loginService from '../../services/authService'
import { AuthContext } from '../../contexts/AuthContext'

import FormInput from '../form/FormInput'
import { APIFailure, LoginUser } from '../../types'

import login_illustration_image from '../../illustrations/login_illustration_image.png'

const Login = () => {
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm<LoginUser>({
    mode: 'onBlur',
  })

  const { user, login } = useContext(AuthContext)

  const onSubmit = async (loginInput: LoginUser) => {
    await loginService(loginInput)
      .then((response) => {
        login(response.data.token, response.data.user)
      })
      .catch((err: Error | AxiosError) => {
        if (axios.isAxiosError(err)) {
          const { response } = err
          const responseData: APIFailure = response?.data

          if (responseData.data?.Email)
            setError('email', {
              type: 'custom',
              message: 'Invalid email or password',
            })
          if (responseData.data?.Password)
            setError('password', {
              type: 'custom',
              message: 'Invalid email or password',
            })
        } else {
          console.log('Something unexpected happened', err)
        }
      })
  }

  console.log(user)

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
