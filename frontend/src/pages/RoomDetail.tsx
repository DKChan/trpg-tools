import { Card, Descriptions, Button, List, Avatar, message, Spin } from 'antd'
import { ArrowLeftOutlined, UserOutlined } from '@ant-design/icons'
import { useNavigate, useParams } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { roomService } from '../services'
import { useAuthStore } from '../store/authStore'

function RoomDetail() {
  const navigate = useNavigate()
  const { id } = useParams()
  const [room, setRoom] = useState<any>(null)
  const [members, setMembers] = useState<any[]>([])
  const [loading, setLoading] = useState(true)
  const { user } = useAuthStore()

  useEffect(() => {
    fetchRoomDetail()
    fetchRoomMembers()
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

  const fetchRoomMembers = async () => {
    try {
      setLoading(true)
      const response = await roomService.getRoomMembers(Number(id))
      if (response.data.code === 200) {
        setMembers(response.data.data)
      } else {
        message.error('获取房间成员失败')
      }
    } catch (error) {
      message.error('获取房间成员失败')
    } finally {
      setLoading(false)
    }
  }

  const handleLeaveRoom = async () => {
    try {
      const response = await roomService.leaveRoom(Number(id))
      if (response.data.code === 200) {
        message.success('已退出房间')
        navigate('/')
      } else {
        message.error('退出房间失败')
      }
    } catch (error) {
      message.error('退出房间失败')
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

  if (loading) {
    return <Spin size="large" className="flex justify-center items-center h-64" />
  }

  if (!room) {
    return <div>房间不存在</div>
  }

  const isDM = user?.id === room.dmid

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
          <Descriptions.Item label="DM">{room.dm?.nickname || '未知'}</Descriptions.Item>
          <Descriptions.Item label="当前人数">{members.length}/{room.max_players}</Descriptions.Item>
          <Descriptions.Item label="邀请码">{room.invite_code}</Descriptions.Item>
          <Descriptions.Item label="房间描述" span={2}>
            {room.description || '暂无描述'}
          </Descriptions.Item>
        </Descriptions>
      </Card>

      <Card
        title="房间成员"
        extra={
          isDM ? (
            <Button danger onClick={handleDeleteRoom}>
              删除房间
            </Button>
          ) : (
            <Button onClick={handleLeaveRoom}>离开房间</Button>
          )
        }
      >
        <List
          dataSource={members}
          renderItem={(member: any) => (
            <List.Item>
              <List.Item.Meta
                avatar={<Avatar icon={<UserOutlined />} />}
                title={member.user?.nickname || '未知用户'}
                description={member.role === 'dm' ? 'DM' : '玩家'}
              />
            </List.Item>
          )}
        />
      </Card>
    </div>
  )
}

export default RoomDetail
