import { Card, Descriptions, Button, message, Spin } from 'antd'
import { ArrowLeftOutlined } from '@ant-design/icons'
import { useNavigate, useParams } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { roomService, characterService } from '../services'

function RoomDetail() {
  const navigate = useNavigate()
  const { id } = useParams()
  const [room, setRoom] = useState<any>(null)
  const [characters, setCharacters] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchRoomDetail()
    fetchCharacters()
  }, [id])

  const fetchRoomDetail = async () => {
    try {
      const response = await roomService.getRoom(Number(id))
      if (response.data.code === 200) {
        setRoom(response.data.data)
      } else {
        message.error('获取房间详情失败')
      }
    } catch (error) {
      message.error('获取房间详情失败')
    }
  }

  const fetchCharacters = async () => {
    try {
      const response = await characterService.getCharacters(Number(id))
      if (response.data.code === 200) {
        setCharacters(response.data.data)
      } else {
        message.error('获取人物卡失败')
      }
    } catch (error) {
      message.error('获取人物卡失败')
    } finally {
      setLoading(false)
    }
  }

  const handleDeleteRoom = async () => {
    try {
      const response = await roomService.deleteRoom(Number(id))
      if (response.data.code === 200) {
        message.success('房间已删除')
        navigate('/')
      } else {
        message.error('删除房间失败')
      }
    } catch (error) {
      message.error('删除房间失败')
    }
  }

  const handleCreateCharacter = () => {
    navigate(`/rooms/${id}/characters/new`)
  }

  if (loading) {
    return <Spin size="large" className="flex justify-center items-center h-64" />
  }

  if (!room) {
    return <div>房间不存在</div>
  }

  return (
    <div>
      <Button
        icon={<ArrowLeftOutlined />}
        onClick={() => navigate('/')}
        className="mb-4"
      >
        返回
      </Button>

      <Card className="mb-6">
        <Descriptions title="房间信息" bordered column={2}>
          <Descriptions.Item label="房间名称">{room.name}</Descriptions.Item>
          <Descriptions.Item label="规则系统">{room.rule_system}</Descriptions.Item>
          <Descriptions.Item label="人物卡数量">{characters.length}</Descriptions.Item>
          <Descriptions.Item label="房间描述" span={2}>
            {room.description || '暂无描述'}
          </Descriptions.Item>
        </Descriptions>
        <div className="mt-4">
          <Button type="primary" onClick={handleCreateCharacter} className="mr-2">
            创建人物卡
          </Button>
          <Button danger onClick={handleDeleteRoom}>
            删除房间
          </Button>
        </div>
      </Card>

      <Card title="人物卡列表">
        {characters.length === 0 ? (
          <div className="text-center text-gray-500 py-8">暂无人物卡</div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {characters.map((char) => (
              <Card
                key={char.id}
                hoverable
                onClick={() => navigate(`/rooms/${id}/characters/${char.id}`)}
              >
                <h3 className="font-bold mb-2">{char.name}</h3>
                <div className="text-sm text-gray-600">
                  <p>种族: {char.race || '-'}</p>
                  <p>职业: {char.class || '-'}</p>
                  <p>等级: {char.level || 1}</p>
                </div>
              </Card>
            ))}
          </div>
        )}
      </Card>
    </div>
  )
}

export default RoomDetail
