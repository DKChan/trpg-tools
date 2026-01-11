import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import Register from './Register'
import { authService } from '../services'

// Mock authService
vi.mock('../services', () => ({
  authService: {
    register: vi.fn(),
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

describe('Register Page', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('渲染注册表单', () => {
    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    )

    expect(screen.getByText('注册')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('邮箱')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('昵称')).toBeInTheDocument()
    expect(screen.getByPlaceholderText('密码', { exact: false })).toBeInTheDocument()
    expect(screen.getByPlaceholderText('确认密码')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: '注册' })).toBeInTheDocument()
  })

  it('验证密码一致性', async () => {
    const user = userEvent.setup()
    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    )

    const passwordInput = screen.getByPlaceholderText('密码', { exact: false })
    const confirmPasswordInput = screen.getByPlaceholderText('确认密码')
    const registerButton = screen.getByRole('button', { name: '注册' })

    // 输入不一致的密码
    await user.type(passwordInput, 'password123')
    await user.type(confirmPasswordInput, 'password456')
    await user.click(registerButton)

    // 应该显示密码不一致提示
    await waitFor(() => {
      const errorText = screen.queryByText(/两次输入的密码不一致/i)
      expect(errorText).toBeInTheDocument()
    })
  })

  it('验证密码最少6位', async () => {
    const user = userEvent.setup()
    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    )

    const passwordInput = screen.getByPlaceholderText('密码', { exact: false })
    const confirmPasswordInput = screen.getByPlaceholderText('确认密码')
    const registerButton = screen.getByRole('button', { name: '注册' })

    // 输入少于6位的密码
    await user.type(passwordInput, '123')
    await user.type(confirmPasswordInput, '123')
    await user.click(registerButton)

    // 应该显示密码长度提示
    await waitFor(() => {
      const errorText = screen.queryByText(/密码至少6位/i)
      expect(errorText).toBeInTheDocument()
    })
  })

  it('成功注册', async () => {
    const user = userEvent.setup()
    const mockResponse = {
      data: {
        code: 200,
        message: 'User registered successfully',
        data: {
          user_id: 1,
          email: 'test@example.com',
          nickname: 'Test User',
        },
      },
    }

    vi.mocked(authService.register).mockResolvedValue(mockResponse as any)

    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    )

    const emailInput = screen.getByPlaceholderText('邮箱')
    const nicknameInput = screen.getByPlaceholderText('昵称')
    const passwordInput = screen.getByPlaceholderText('密码', { exact: false })
    const confirmPasswordInput = screen.getByPlaceholderText('确认密码')
    const registerButton = screen.getByRole('button', { name: '注册' })

    // 输入有效数据
    await user.type(emailInput, 'test@example.com')
    await user.type(nicknameInput, 'Test User')
    await user.type(passwordInput, 'password123')
    await user.type(confirmPasswordInput, 'password123')
    await user.click(registerButton)

    await waitFor(() => {
      expect(authService.register).toHaveBeenCalledWith({
        email: 'test@example.com',
        nickname: 'Test User',
        password: 'password123',
      })
      expect(message.success).toHaveBeenCalledWith('注册成功')
    })
  })

  it('注册失败显示错误提示', async () => {
    const user = userEvent.setup()
    vi.mocked(authService.register).mockRejectedValue(new Error('Registration failed'))

    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    )

    const emailInput = screen.getByPlaceholderText('邮箱')
    const nicknameInput = screen.getByPlaceholderText('昵称')
    const passwordInput = screen.getByPlaceholderText('密码', { exact: false })
    const confirmPasswordInput = screen.getByPlaceholderText('确认密码')
    const registerButton = screen.getByRole('button', { name: '注册' })

    await user.type(emailInput, 'test@example.com')
    await user.type(nicknameInput, 'Test User')
    await user.type(passwordInput, 'password123')
    await user.type(confirmPasswordInput, 'password123')
    await user.click(registerButton)

    await waitFor(() => {
      expect(message.error).toHaveBeenCalledWith('注册失败')
    })
  })
})
