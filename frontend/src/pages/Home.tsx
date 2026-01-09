import { useState } from 'react'
import { Card, Row, Col, Button, Input, Modal, Form, message } from 'antd'
import { PlusOutlined, SearchOutlined } from '@ant-design/icons'
import { roomService } from '../services'

function Home() {
  const [rooms, setRooms] = useState<any[]>([])
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [form] = Form.useForm()

  const handleCreateRoom = async (values: any) => {
    try {
      const response = await roomService.createRoom(values)
      if (response.data.code === 200) {
        message.success('房间创建成功')
        setIsModalOpen(false)
        form.resetFields()
      }
    } catch (error) {
      message.error('房间创建失败')
    }
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
        <Col xs={24} sm={12} md={8} lg={6}>
          <Card
            title="示例房间"
            extra={<span className="text-sm text-gray-500">DND5e</span>}
            hoverable
            className="h-full"
          >
            <p className="text-gray-600 mb-4">这是一个示例房间</p>
            <div className="flex justify-between text-sm text-gray-500">
              <span>DM: 示例用户</span>
              <span>3/10 人</span>
            </div>
          </Card>
        </Col>
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
