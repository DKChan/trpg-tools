import api from './api'
import { ApiResponse } from '../types'

export const roomService = {
  getRooms: () => api.get<ApiResponse<any[]>>('/rooms'),

  getRoom: (id: number) => api.get<ApiResponse<any>>(`/rooms/${id}`),

  createRoom: (data: {
    name: string
    description?: string
    rule_system?: string
  }) => api.post<ApiResponse<any>>('/rooms', data),

  deleteRoom: (id: number) => api.delete<ApiResponse<null>>(`/rooms/${id}`),
}

export const characterService = {
  getCharacters: (roomId: number) =>
    api.get<ApiResponse<any[]>>(`/rooms/${roomId}/characters`),

  getCharacter: (roomId: number, id: number) =>
    api.get<ApiResponse<any>>(`/rooms/${roomId}/characters/${id}`),

  createCharacter: (roomId: number, data: any) =>
    api.post<ApiResponse<any>>(`/rooms/${roomId}/characters`, data),

  updateCharacter: (roomId: number, id: number, data: any) =>
    api.put<ApiResponse<null>>(`/rooms/${roomId}/characters/${id}`, data),

  deleteCharacter: (roomId: number, id: number) =>
    api.delete<ApiResponse<null>>(`/rooms/${roomId}/characters/${id}`),
}
