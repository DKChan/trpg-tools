import { renderHook, act } from '@testing-library/react'
import { useAuthStore } from './authStore'

describe('AuthStore', () => {
  beforeEach(() => {
    // 清除 localStorage
    localStorage.clear()
    // 重置 store
    useAuthStore.setState({ user: null, token: null })
  })

  it('初始化时用户和 token 为 null', () => {
    const { result } = renderHook(() => useAuthStore())

    expect(result.current.user).toBeNull()
    expect(result.current.token).toBeNull()
  })

  it('setAuth 正确设置用户和 token', () => {
    const { result } = renderHook(() => useAuthStore())

    const mockUser = {
      user_id: 1,
      email: 'test@example.com',
      nickname: 'Test User',
    }
    const mockToken = 'test-token'

    act(() => {
      result.current.setAuth(mockUser as any, mockToken)
    })

    expect(result.current.user).toEqual(mockUser)
    expect(result.current.token).toBe(mockToken)
  })

  it('setAuth 将状态持久化到 localStorage', () => {
    const { result } = renderHook(() => useAuthStore())

    const mockUser = {
      user_id: 1,
      email: 'test@example.com',
      nickname: 'Test User',
    }
    const mockToken = 'test-token'

    act(() => {
      result.current.setAuth(mockUser as any, mockToken)
    })

    const storedData = localStorage.getItem('auth-storage')
    expect(storedData).toBeTruthy()

    const parsed = JSON.parse(storedData!)
    expect(parsed.state.user).toEqual(mockUser)
    expect(parsed.state.token).toBe(mockToken)
  })

  it('logout 清除用户和 token', () => {
    const { result } = renderHook(() => useAuthStore())

    // 先设置数据
    const mockUser = {
      user_id: 1,
      email: 'test@example.com',
      nickname: 'Test User',
    }
    const mockToken = 'test-token'

    act(() => {
      result.current.setAuth(mockUser as any, mockToken)
    })

    expect(result.current.user).toEqual(mockUser)
    expect(result.current.token).toBe(mockToken)

    // 执行 logout
    act(() => {
      result.current.logout()
    })

    expect(result.current.user).toBeNull()
    expect(result.current.token).toBeNull()
  })

  it('logout 清除 localStorage 中的数据', () => {
    const { result } = renderHook(() => useAuthStore())

    const mockUser = {
      user_id: 1,
      email: 'test@example.com',
      nickname: 'Test User',
    }
    const mockToken = 'test-token'

    act(() => {
      result.current.setAuth(mockUser as any, mockToken)
    })

    act(() => {
      result.current.logout()
    })

    const storedData = localStorage.getItem('auth-storage')
    const parsed = JSON.parse(storedData!)
    expect(parsed.state.user).toBeNull()
    expect(parsed.state.token).toBeNull()
  })

  it('从 localStorage 恢复状态', () => {
    // 先设置 localStorage
    const mockUser = {
      user_id: 1,
      email: 'test@example.com',
      nickname: 'Test User',
    }
    const mockToken = 'test-token'

    localStorage.setItem(
      'auth-storage',
      JSON.stringify({
        state: { user: mockUser, token: mockToken },
        version: 0,
      })
    )

    // 重新创建 store
    const { result } = renderHook(() => useAuthStore())

    expect(result.current.user).toEqual(mockUser)
    expect(result.current.token).toBe(mockToken)
  })
})
