import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { QueryClientProvider } from 'react-query'
import { SnackbarProvider } from 'notistack'

import App from './App'

import { AuthProvider } from './contexts/AuthContext'
import queryClient from './util/queryClient'

import './index.css'

const PUBLIC_URL = ''

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <BrowserRouter basename={PUBLIC_URL}>
    <AuthProvider>
      <QueryClientProvider client={queryClient}>
        <SnackbarProvider preventDuplicate>
          <App />
        </SnackbarProvider>
      </QueryClientProvider>
    </AuthProvider>
  </BrowserRouter>
)
