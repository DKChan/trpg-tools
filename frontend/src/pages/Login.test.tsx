import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import Login from './Login'
import { authService } from '../services'

// Mock authService
vi.mock('../services', () => ({
  authService: {
    login: vi.fn(),
  },
}))

// Mock react-router-dom
vi.mock('react-router-dom', async () => ({
  ...((await vi.importActual('react-router-dom')) as any),
  useNavigate: () => vi.fn(),
}))

// Mock antd message
vi.mock('antd', async () => ({
  ...(await vi.importActual('antd')),
  message: {
    success: vi.fn(),
    error: vi.fn(),
  },
}))

import { message } from 'antd'

describe('Login Page', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('渲染登录表单', () => {
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )

    expect(screen.getByText('登录')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('邮箱')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('密码')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: '登录' })).toBeInTheDocument()
  })

  it('显示"还没有账号？"提示和注册链接', () => {
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )

    expect(screen.getByText('还没有账号？')).toBeInTheDocument()
    expect(screen.getByText('立即注册')).toBeInTheDocument()
  })

  it('验证邮箱格式', async () => {
    const user = userEvent.setup()
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )

    const emailInput = screen.getByPlaceholderText('邮箱')
    const loginButton = screen.getByRole('button', { name: '登录' })

    // 输入无效邮箱
    await user.type(emailInput, 'invalid-email')
    await user.click(loginButton)

    // 应该显示邮箱格式错误提示
    await waitFor(() => {
      const errorText = screen.queryByText('请输入有效的邮箱地址')
      expect(errorText).toBeInTheDocument()
    })
  })

  it('验证必填字段', async () => {
    const user = userEvent.setup()
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )

    const loginButton = screen.getByRole('button', { name: '登录' })

    // 不输入任何字段直接点击登录
    await user.click(loginButton)

    // 应该显示必填字段提示
    await waitFor(() => {
      const emailError = screen.queryByText('请输入邮箱')
      const passwordError = screen.queryByText('请输入密码')
      expect(emailError || passwordError).toBeTruthy()
    })
  })

  it('成功登录', async () => {
    const user = userEvent.setup()
    const mockResponse = {
      data: {
        code: 200,
        message: 'success',
        data: {
          user_id: 1,
          email: 'test@example.com',
          nickname: 'Test User',
        },
      },
    }

    vi.mocked(authService.login).mockResolvedValue(mockResponse as any)

    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )

    const emailInput = screen.getByPlaceholderText('邮箱')
    const passwordInput = screen.getByPlaceholderText('密码')
    const loginButton = screen.getByRole('button', { name: '登录' })

    // 输入有效数据
    await user.type(emailInput, 'test@example.com')
    await user.type(passwordInput, 'password123')
    await user.click(loginButton)

    await waitFor(() => {
      expect(authService.login).toHaveBeenCalledWith({
        email: 'test@example.com',
        password: 'password123',
      })
      expect(message.success).toHaveBeenCalledWith('登录成功')
    })
  })

  it('登录失败显示错误提示', async () => {
    const user = userEvent.setup()
    vi.mocked(authService.login).mockRejectedValue(new Error('Login failed'))

    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    )

    const emailInput = screen.getByPlaceholderText('邮箱')
    const passwordInput = screen.getByPlaceholderText('密码')
    const loginButton = screen.getByRole('button', { name: '登录' })

    await user.type(emailInput, 'test@example.com')
    await user.type(passwordInput, 'wrongpassword')
    await user.click(loginButton)

    await waitFor(() => {
      expect(message.error).toHaveBeenCalledWith('登录失败')
    })
  })
})
