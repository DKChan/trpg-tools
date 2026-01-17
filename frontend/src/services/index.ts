import api from './api'
import { ApiResponse, User } from '../types'

export const authService = {
  register: (data: { email: string; password: string; nickname: string }) =>
    api.post<ApiResponse<User>>('/auth/register', data),

  login: (data: { email: string; password: string }) =>
    api.post<ApiResponse<User>>('/auth/login', data),
}

export const userService = {
  getProfile: () => api.get<ApiResponse<User>>('/user/profile'),

  updateProfile: (data: { nickname?: string; avatar?: string }) =>
    api.put<ApiResponse<null>>('/user/profile', data),

  updatePassword: (data: { old_password: string; new_password: string }) =>
    api.put<ApiResponse<null>>('/user/password', data),
}

export const roomService = {
  getRooms: () => api.get<ApiResponse<any[]>>('/rooms'),

  getRoom: (id: number) => api.get<ApiResponse<any>>(`/rooms/${id}`),

  createRoom: (data: {
    name: string
    description?: string
    rule_system?: string
    password?: string
    max_players?: number
    is_public?: boolean
  }) => api.post<ApiResponse<any>>('/rooms', data),

  joinRoom: (id: number) => api.post<ApiResponse<null>>(`/rooms/${id}/join`),

  leaveRoom: (id: number) => api.post<ApiResponse<null>>(`/rooms/${id}/leave`),

  deleteRoom: (id: number) => api.delete<ApiResponse<null>>(`/rooms/${id}`),

  getRoomMembers: (id: number) => api.get<ApiResponse<any[]>>(`/rooms/${id}/members`),

  kickMember: (id: number, userId: number) =>
    api.put<ApiResponse<null>>(`/rooms/${id}/members/${userId}/kick`),

  transferDM: (id: number, userId: number) =>
    api.put<ApiResponse<null>>(`/rooms/${id}/transfer-dm`, { user_id: userId }),
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
