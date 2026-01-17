export interface Room {
  id: number
  name: string
  description: string
  rule_system: string
  created_at: string
  updated_at: string
}

export interface CharacterCard {
  id: number
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
