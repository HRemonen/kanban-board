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
