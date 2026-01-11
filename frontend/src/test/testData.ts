/**
 * 测试数据生成器
 */

// Mock 用户数据
export const mockUser = {
  user_id: 1,
  email: 'test@example.com',
  nickname: 'Test User',
  avatar: 'https://example.com/avatar.png',
}

// Mock 房间数据
export const mockRoom = {
  id: 1,
  name: 'Test Room',
  description: 'A test room for testing',
  rule_system: 'DND5e',
  dm_id: 1,
  max_players: 5,
  is_public: true,
  created_at: '2024-01-01T00:00:00Z',
  updated_at: '2024-01-01T00:00:00Z',
}

// Mock 人物卡数据
export const mockCharacter = {
  id: 1,
  user_id: 1,
  room_id: 1,
  name: 'Test Character',
  race: 'Human',
  class: 'Fighter',
  level: 1,
  background: 'Soldier',
  alignment: 'Lawful Good',
  strength: 16,
  dexterity: 14,
  constitution: 15,
  intelligence: 10,
  wisdom: 12,
  charisma: 8,
  hp: 10,
  max_hp: 10,
  ac: 10,
  speed: 30,
  skills: [],
  equipment: [],
  spells: [],
  created_at: '2024-01-01T00:00:00Z',
  updated_at: '2024-01-01T00:00:00Z',
}

// Mock 房间列表
export const mockRooms = [
  mockRoom,
  {
    ...mockRoom,
    id: 2,
    name: 'Test Room 2',
  },
]

// Mock 人物卡列表
export const mockCharacters = [
  mockCharacter,
  {
    ...mockCharacter,
    id: 2,
    name: 'Test Character 2',
  },
]
