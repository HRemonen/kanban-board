import { useContext, useState } from 'react'
import axios, { AxiosError } from 'axios'
import { useForm } from 'react-hook-form'
import { Link, useNavigate } from 'react-router-dom'

import loginService from '../../services/authService'
import { AuthContext } from '../../contexts/AuthContext'

import FormInput from '../form/FormInput'
import { APIFailure, LoginUser } from '../../types'

import login_illustration_image from '../../illustrations/login_illustration_image.svg'
import LoadingSpinner from '../common/Loading'

const Login = () => {
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors },
  } = useForm<LoginUser>({
    mode: 'onBlur',
  })

  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  const { login } = useContext(AuthContext)

  const onLogin = async (loginInput: LoginUser) => {
    setLoading(true)
    await loginService(loginInput)
      .then((response) => {
        login(response.data.token, response.data.user)
        setLoading(false)
        navigate('/')
      })
      .catch((err: Error | AxiosError) => {
        setLoading(false)
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
          else {
            setError('root', {
              type: 'custom',
              message: responseData.message,
            })
          }
        } else {
          console.log('Something unexpected happened', err)
        }
      })
  }

  return (
    <section className="bg-[#A8A4FF] md:grid md:grid-cols-2 text-center">
      <div className="flex flex-col h-screen justify-center text-center items-center p-12">
        <form
          onSubmit={handleSubmit(onLogin)}
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

          {!loading ? (
            <button
              id="login-button"
              data-cy="login-form-button"
              type="submit"
              className="inline-block text-white bg-[#24292F] hover:bg-[#24292F]/90 focus:ring-4 focus:outline-none focus:ring-[#24292F]/50 font-medium rounded-lg text-sm px-5 py-2.5 text-center items-center mr-2 mb-2"
            >
              Login
            </button>
          ) : (
            <div className="flex-col inline-flex items-center">
              <LoadingSpinner />
            </div>
          )}
        </form>

        {errors.root && (
          <p className="text-red-500 text-sm mt-2">{errors.root.message}</p>
        )}

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
        <img className="scale-75" src={login_illustration_image} alt="" />
      </div>
    </section>
  )
}

export default Login
