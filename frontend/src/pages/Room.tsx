import { useParams } from 'react-router-dom'
import { Card, Typography, Tabs, List, Avatar, Button } from 'antd'
import { UserOutlined, PlusOutlined } from '@ant-design/icons'

const { Title } = Typography
const { TabPane } = Tabs

const Room: React.FC = () => {
  const { id } = useParams<{ id: string }>()

  // TODO: 从API获取房间信息
  const room = {
    id: id || '',
    name: '测试房间',
    description: '这是一个测试房间，用于演示功能',
    ruleType: 'DND5e',
    creator: 'DM张三',
    createdAt: '2026-01-05'
  }

  // TODO: 从API获取房间成员列表
  const members = [
    { id: '1', name: 'DM张三', role: 'dm', avatar: '' },
    { id: '2', name: '玩家李四', role: 'player', avatar: '' },
    { id: '3', name: '玩家王五', role: 'player', avatar: '' },
  ]

  // TODO: 从API获取人物卡列表
  const characterSheets = [
    { id: '1', name: '冒险者小明', user: '玩家李四', race: '人类', class: '战士' },
    { id: '2', name: '法师小红', user: '玩家王五', race: '精灵', class: '法师' },
  ]

  return (
    <div>
      <Title level={2}>{room.name}</Title>
      <div className="mb-6">
        <p>规则类型：{room.ruleType}</p>
        <p>创建者：{room.creator}</p>
        <p>创建时间：{room.createdAt}</p>
        <p className="mt-2">{room.description}</p>
      </div>

      <Tabs defaultActiveKey="1" className="mb-6">
        <TabPane tab="成员列表" key="1">
          <Card>
            <List
              dataSource={members}
              renderItem={(member) => (
                <List.Item
                  actions={[
                    member.role === 'dm' ? 'DM' : '玩家'
                  ]}
                >
                  <List.Item.Meta
                    avatar={<Avatar icon={<UserOutlined />} />}
                    title={member.name}
                  />
                </List.Item>
              )}
            />
            <div className="mt-4 text-center">
              <Button type="primary" icon={<PlusOutlined />}>
                邀请成员
              </Button>
            </div>
          </Card>
        </TabPane>
        <TabPane tab="人物卡" key="2">
          <Card>
            <List
              dataSource={characterSheets}
              renderItem={(sheet) => (
                <List.Item
                  actions={[
                    <Button type="link" href={`/character-sheet/${sheet.id}`}>
                      查看
                    </Button>
                  ]}
                >
                  <List.Item.Meta
                    title={sheet.name}
                    description={`${sheet.race} ${sheet.class} - ${sheet.user}`}
                  />
                </List.Item>
              )}
            />
            <div className="mt-4 text-center">
              <Button type="primary" icon={<PlusOutlined />}>
                创建人物卡
              </Button>
            </div>
          </Card>
        </TabPane>
        <TabPane tab="房间设置" key="3">
          <Card>
            <p>房间设置内容</p>
            {/* TODO: 实现房间设置功能 */}
          </Card>
        </TabPane>
      </Tabs>
    </div>
  )
}

export default Room
