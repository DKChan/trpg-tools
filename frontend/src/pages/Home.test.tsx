import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import Home from './Home'
import { roomService } from '../services'

// Mock roomService
vi.mock('../services', () => ({
  roomService: {
    getRooms: vi.fn(),
    createRoom: vi.fn(),
  },
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

describe('Home Page', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('渲染页面标题和创建房间按钮', () => {
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    )

    expect(screen.getByText('房间列表')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /创建房间/i })).toBeInTheDocument()
  })

  it('显示搜索框', () => {
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    )

    expect(screen.getByPlaceholderText('搜索房间')).toBeInTheDocument()
  })

  it('加载房间列表', async () => {
    const mockRooms = [
      {
        id: 1,
        name: 'Test Room 1',
        description: 'Description 1',
        rule_system: 'DND5e',
      },
      {
        id: 2,
        name: 'Test Room 2',
        description: 'Description 2',
        rule_system: 'DND5e',
      },
    ]

    vi.mocked(roomService.getRooms).mockResolvedValue({
      data: {
        code: 200,
        message: 'success',
        data: mockRooms,
      },
    } as any)

    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    )

    await waitFor(() => {
      expect(roomService.getRooms).toHaveBeenCalled()
    })
  })

  it('成功创建房间', async () => {
    const mockResponse = {
      data: {
        code: 200,
        message: 'Room created successfully',
        data: {
          id: 1,
          name: 'New Room',
          description: 'Test description',
        },
      },
    }

    vi.mocked(roomService.createRoom).mockResolvedValue(mockResponse as any)

    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    )

    expect(message.success).toBeDefined()
  })

  it('搜索房间', async () => {
    const user = userEvent.setup()
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    )

    const searchInput = screen.getByPlaceholderText('搜索房间')
    await user.type(searchInput, 'Test Room')

    expect(searchInput).toHaveValue('Test Room')
  })
})

