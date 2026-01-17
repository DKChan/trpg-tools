import { describe, it, expect, vi, beforeEach } from 'vitest'
import { roomService, characterService } from './index'
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
      rule_system: 'DND5e',
    }

    roomService.createRoom(mockData)

    expect(api.post).toHaveBeenCalledWith('/rooms', mockData)
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
