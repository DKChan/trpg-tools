import { Card, Descriptions, Button, List, Avatar } from 'antd'
import { ArrowLeftOutlined, UserOutlined } from '@ant-design/icons'
import { useNavigate, useParams } from 'react-router-dom'

function RoomDetail() {
  const navigate = useNavigate()
  const { id } = useParams()

  const members = [
    { id: 1, name: '示例玩家1', role: 'player' },
    { id: 2, name: '示例玩家2', role: 'player' },
    { id: 3, name: '示例玩家3', role: 'player' },
  ]

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
          <Descriptions.Item label="房间名称">示例房间</Descriptions.Item>
          <Descriptions.Item label="规则系统">DND5e</Descriptions.Item>
          <Descriptions.Item label="DM">示例用户</Descriptions.Item>
          <Descriptions.Item label="当前人数">3/10</Descriptions.Item>
          <Descriptions.Item label="房间描述" span={2}>
            这是一个示例房间
          </Descriptions.Item>
        </Descriptions>
      </Card>

      <Card title="房间成员">
        <List
          dataSource={members}
          renderItem={(member) => (
            <List.Item>
              <List.Item.Meta
                avatar={<Avatar icon={<UserOutlined />} />}
                title={member.name}
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
