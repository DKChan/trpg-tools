export interface User {
  user_id: number
  email: string
  nickname: string
  avatar?: string
}

export interface Room {
  id: number
  name: string
  description: string
  rule_system: string
  invite_code: string
  dm_id: number
  max_players: number
  is_public: boolean
  created_at: string
  updated_at: string
}

export interface CharacterCard {
  id: number
  user_id: number
  room_id: number
  name: string
  race: string
  class: string
  level: number
  background: string
  alignment: string
  strength: number
  dexterity: number
  constitution: number
  intelligence: number
  wisdom: number
  charisma: number
  ac: number
  hp: number
  max_hp: number
  speed: number
  proficiency: number
  skills: string
  saves: string
  equipment: string
  spells: string
  created_at: string
  updated_at: string
}

export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}
