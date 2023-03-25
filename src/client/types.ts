import { InputHTMLAttributes } from 'react'
import { FieldError, UseFormRegister, FieldValues } from 'react-hook-form'

// USER RELATED

export interface RegisterUser {
  email: string
  name?: string
  password: string
  confirmPassword: string
}

export interface LoginUser {
  email: string
  password: string
}

export interface PublicUser {
  id: string
  name?: string
  email: string
  password: string
  boards: any[]
}

export interface PrivateUser extends PublicUser {
  role: 'user' | string
  provider: 'local' | 'google'
  photo: string
  Boards: any[]
  CreatedAt: Date
  UpdatedAt: Date
}

export interface APIResponse {
  data: PublicUser[] | PrivateUser | null
  message: string
  status: 'success' | 'fail'
}

// FORM RELATED

export interface InputType extends InputHTMLAttributes<HTMLInputElement> {
  register: UseFormRegister<FieldValues> | UseFormRegister<any>
  error: FieldError | undefined
  label: string
  id: string
}
