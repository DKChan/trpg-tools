import { describe, it, expect, vi, beforeEach } from 'vitest'
import { authService, userService, roomService, characterService } from './index'
import api from './api'

// Mock api
vi.mock('./api', () => ({
  default: {
    get: vi.fn(),
    post: vi.fn(),
    put: vi.fn(),
    delete: vi.fn(),
  },
}))

describe('Auth Service', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('register 调用正确的 API', () => {
    const mockData = {
      email: 'test@example.com',
      password: 'password123',
      nickname: 'Test User',
    }

    authService.register(mockData)

    expect(api.post).toHaveBeenCalledWith('/auth/register', mockData)
  })

  it('login 调用正确的 API', () => {
    const mockData = {
      email: 'test@example.com',
      password: 'password123',
    }

    authService.login(mockData)

    expect(api.post).toHaveBeenCalledWith('/auth/login', mockData)
  })
})

describe('User Service', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('getProfile 调用正确的 API', () => {
    userService.getProfile()

    expect(api.get).toHaveBeenCalledWith('/user/profile')
  })

  it('updateProfile 调用正确的 API', () => {
    const mockData = {
      nickname: 'Updated Name',
      avatar: 'https://example.com/avatar.png',
    }

    userService.updateProfile(mockData)

    expect(api.put).toHaveBeenCalledWith('/user/profile', mockData)
  })

  it('updatePassword 调用正确的 API', () => {
    const mockData = {
      old_password: 'oldpassword',
      new_password: 'newpassword123',
    }

    userService.updatePassword(mockData)

    expect(api.put).toHaveBeenCalledWith('/user/password', mockData)
  })
})

describe('Room Service', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('getRooms 调用正确的 API', () => {
    roomService.getRooms()

    expect(api.get).toHaveBeenCalledWith('/rooms')
  })

  it('getRoom 调用正确的 API', () => {
    const roomId = 1
    roomService.getRoom(roomId)

    expect(api.get).toHaveBeenCalledWith(`/rooms/${roomId}`)
  })

  it('createRoom 调用正确的 API', () => {
    const mockData = {
      name: 'Test Room',
      description: 'Test description',
      max_players: 5,
      is_public: true,
    }

    roomService.createRoom(mockData)

    expect(api.post).toHaveBeenCalledWith('/rooms', mockData)
  })

  it('joinRoom 调用正确的 API', () => {
    const roomId = 1
    roomService.joinRoom(roomId)

    expect(api.post).toHaveBeenCalledWith(`/rooms/${roomId}/join`)
  })

  it('leaveRoom 调用正确的 API', () => {
    const roomId = 1
    roomService.leaveRoom(roomId)

    expect(api.post).toHaveBeenCalledWith(`/rooms/${roomId}/leave`)
  })

  it('deleteRoom 调用正确的 API', () => {
    const roomId = 1
    roomService.deleteRoom(roomId)

    expect(api.delete).toHaveBeenCalledWith(`/rooms/${roomId}`)
  })
})

describe('Character Service', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('getCharacters 调用正确的 API', () => {
    const roomId = 1
    characterService.getCharacters(roomId)

    expect(api.get).toHaveBeenCalledWith(`/rooms/${roomId}/characters`)
  })

  it('getCharacter 调用正确的 API', () => {
    const roomId = 1
    const characterId = 1
    characterService.getCharacter(roomId, characterId)

    expect(api.get).toHaveBeenCalledWith(`/rooms/${roomId}/characters/${characterId}`)
  })

  it('createCharacter 调用正确的 API', () => {
    const roomId = 1
    const mockData = {
      name: 'Test Character',
      race: 'Human',
      class: 'Fighter',
    }

    characterService.createCharacter(roomId, mockData)

    expect(api.post).toHaveBeenCalledWith(`/rooms/${roomId}/characters`, mockData)
  })

  it('updateCharacter 调用正确的 API', () => {
    const roomId = 1
    const characterId = 1
    const mockData = {
      name: 'Updated Character',
    }

    characterService.updateCharacter(roomId, characterId, mockData)

    expect(api.put).toHaveBeenCalledWith(
      `/rooms/${roomId}/characters/${characterId}`,
      mockData
    )
  })

  it('deleteCharacter 调用正确的 API', () => {
    const roomId = 1
    const characterId = 1
    characterService.deleteCharacter(roomId, characterId)

    expect(api.delete).toHaveBeenCalledWith(`/rooms/${roomId}/characters/${characterId}`)
  })
})
