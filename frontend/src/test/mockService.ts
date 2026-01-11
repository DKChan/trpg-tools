/**
 * Mock API 服务
 */
import { vi } from 'vitest'

// Mock axios instance
export const mockApi = {
  get: vi.fn(),
  post: vi.fn(),
  put: vi.fn(),
  delete: vi.fn(),
}

// Mock 成功响应
export const mockSuccessResponse = (data: any) => ({
  data: {
    code: 200,
    message: 'success',
    data,
  },
})

// Mock 错误响应
export const mockErrorResponse = (message: string, code: number = 400) => ({
  response: {
    data: {
      code,
      message,
    },
  },
})
