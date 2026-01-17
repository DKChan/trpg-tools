import { useState, useEffect } from 'react'
import { Card, Row, Col, Button, Input, Modal, Form, message, Spin } from 'antd'
import { PlusOutlined, SearchOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { roomService } from '../services'

function Home() {
  const navigate = useNavigate()
  const [rooms, setRooms] = useState<any[]>([])
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const [form] = Form.useForm()

  useEffect(() => {
    fetchRooms()
  }, [])

  const fetchRooms = async () => {
    try {
      setLoading(true)
      const response = await roomService.getRooms()
      if (response.data.code === 200) {
        setRooms(response.data.data)
      }
    } catch (error) {
      message.error('获取房间列表失败')
    } finally {
      setLoading(false)
    }
  }

  const handleCreateRoom = async (values: any) => {
    try {
      const response = await roomService.createRoom(values)
      if (response.data.code === 200) {
        message.success('房间创建成功')
        setIsModalOpen(false)
        form.resetFields()
        fetchRooms()
      }
    } catch (error) {
      message.error('房间创建失败')
    }
  }

  if (loading) {
    return <Spin size="large" className="flex justify-center items-center h-64" />
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">房间列表</h1>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => setIsModalOpen(true)}
        >
          创建房间
        </Button>
      </div>

      <Input
        placeholder="搜索房间"
        prefix={<SearchOutlined />}
        className="mb-6 max-w-md"
      />

      <Row gutter={[16, 16]}>
        {rooms.map((room) => (
          <Col xs={24} sm={12} md={8} lg={6} key={room.id}>
            <Card
              title={room.name}
              extra={<span className="text-sm text-gray-500">{room.rule_system}</span>}
              hoverable
              className="h-full"
              onClick={() => navigate(`/rooms/${room.id}`)}
            >
              <p className="text-gray-600 mb-4">{room.description || '暂无描述'}</p>
            </Card>
          </Col>
        ))}
      </Row>

      <Modal
        title="创建房间"
        open={isModalOpen}
        onCancel={() => setIsModalOpen(false)}
        footer={null}
      >
        <Form form={form} onFinish={handleCreateRoom} layout="vertical">
          <Form.Item
            name="name"
            label="房间名称"
            rules={[{ required: true, message: '请输入房间名称' }]}
          >
            <Input placeholder="请输入房间名称" />
          </Form.Item>

          <Form.Item name="description" label="房间描述">
            <Input.TextArea placeholder="请输入房间描述" rows={3} />
          </Form.Item>

          <Form.Item name="rule_system" label="规则系统" initialValue="DND5e">
            <Input disabled />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" className="w-full">
              创建
            </Button>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

export default Home

