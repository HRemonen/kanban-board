import { useContext, useState } from 'react'
import axios, { AxiosError } from 'axios'
import { useForm } from 'react-hook-form'
import { Link, useNavigate } from 'react-router-dom'
import { enqueueSnackbar } from 'notistack'

import loginService from '../../services/authService'
import { AuthContext } from '../../contexts/AuthContext'

import LoginInput from '../form/LoginInput'
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
        navigate('/')
        enqueueSnackbar(`Login success ${response.data.user.username}`, {
          variant: 'success',
        })
      })
      .catch((err: Error | AxiosError) => {
        if (!axios.isAxiosError(err)) {
          enqueueSnackbar('Could not log in at the moment', {
            variant: 'error',
          })
          return
        }
        const { response } = err
        const responseData: APIFailure = response?.data

        if (
          responseData.data?.Email ||
          responseData.message === 'record not found'
        )
          setError('email', {
            type: 'custom',
            message: 'Invalid email',
          })
        if (
          responseData.data?.Password ||
          responseData.message === 'Invalid password'
        )
          setError('password', {
            type: 'custom',
            message: 'Invalid password',
          })
      })
      .finally(() => setLoading(false))
  }

  return (
    <section className='bg-[#A8A4FF] text-center md:grid md:grid-cols-2'>
      <div className='flex h-screen flex-col items-center justify-center p-12 text-center'>
        <form
          onSubmit={handleSubmit(onLogin)}
          className='flex w-[80%] flex-col justify-center text-left'
        >
          <LoginInput
            id='email'
            type='email'
            placeholder='hello@world.com'
            name='email'
            label='Email'
            register={register}
            error={errors.email}
          />

          <LoginInput
            id='password'
            type='password'
            placeholder='Password'
            name='password'
            label='Password'
            register={register}
            error={errors.password}
          />

          {!loading ? (
            <button
              id='login-button'
              data-cy='login-form-button'
              type='submit'
              className='mb-2 mr-2 inline-block items-center rounded-lg bg-[#24292F] px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-[#24292F]/90 focus:outline-none focus:ring-4 focus:ring-[#24292F]/50'
            >
              Login
            </button>
          ) : (
            <div className='inline-flex flex-col items-center'>
              <LoadingSpinner />
            </div>
          )}
        </form>

        {errors.root && (
          <p className='mt-2 text-sm text-red-500'>{errors.root.message}</p>
        )}

        <p className='mt-4'>
          Don&lsquo;t have an account?
          <Link
            to='/register'
            className='ml-2 inline-flex items-center font-medium text-blue-600 hover:underline'
          >
            Register here
          </Link>
        </p>
      </div>

      <div className='hidden flex-col justify-center md:flex '>
        <img className='scale-75' src={login_illustration_image} alt='' />
      </div>
    </section>
  )
}

export default Login
