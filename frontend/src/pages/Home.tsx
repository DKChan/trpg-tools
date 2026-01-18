import { useState, useEffect } from 'react'
import { Card, Row, Col, Button, Input, Modal, Form, message, Spin, Popconfirm } from 'antd'
import { PlusOutlined, SearchOutlined, DeleteOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { roomService } from '../services'

function Home() {
  const navigate = useNavigate()
  const [rooms, setRooms] = useState<any[]>([])
  const [filteredRooms, setFilteredRooms] = useState<any[]>([])
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [deleteModalOpen, setDeleteModalOpen] = useState(false)
  const [deleteTarget, setDeleteTarget] = useState<{id: number; name: string} | null>(null)
  const [deleteConfirmName, setDeleteConfirmName] = useState('')
  const [loading, setLoading] = useState(false)
  const [deleting, setDeleting] = useState(false)
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
        setFilteredRooms(response.data.data)
      }
    } catch (error) {
      message.error('获取房间列表失败')
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value
    if (!value) {
      setFilteredRooms(rooms)
    } else {
      const filtered = rooms.filter((room) =>
        room.name.toLowerCase().includes(value.toLowerCase()) ||
        (room.description && room.description.toLowerCase().includes(value.toLowerCase()))
      )
      setFilteredRooms(filtered)
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

  const handleDeleteRoom = async () => {
    if (!deleteTarget) return

    setDeleting(true)
    try {
      const response = await roomService.deleteRoom(deleteTarget.id)
      if (response.data.code === 200) {
        message.success('房间删除成功')
        setDeleteModalOpen(false)
        setDeleteTarget(null)
        setDeleteConfirmName('')
        fetchRooms()
      } else {
        message.error('房间删除失败')
      }
    } catch (error) {
      message.error('房间删除失败')
    } finally {
      setDeleting(false)
    }
  }

  const openDeleteModal = (room: any, e: React.MouseEvent) => {
    e.stopPropagation()
    setDeleteTarget({ id: room.id, name: room.name })
    setDeleteConfirmName('')
    setDeleteModalOpen(true)
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
        onChange={handleSearch}
        allowClear
      />

      <Row gutter={[16, 16]}>
        {filteredRooms.map((room) => (
          <Col xs={24} sm={12} md={8} lg={6} key={room.id}>
            <Card
              title={room.name}
              extra={
                <div className="flex items-center gap-2">
                  <span className="text-sm text-gray-500">{room.rule_system}</span>
                  <Button
                    type="text"
                    danger
                    icon={<DeleteOutlined />}
                    size="small"
                    onClick={(e) => openDeleteModal(room, e)}
                  />
                </div>
              }
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

      <Modal
        title="删除房间"
        open={deleteModalOpen}
        onCancel={() => setDeleteModalOpen(false)}
        onOk={handleDeleteRoom}
        okText="确认删除"
        cancelText="取消"
        okButtonProps={{ disabled: deleteConfirmName !== deleteTarget?.name, loading: deleting }}
        width={520}
      >
        <div className="space-y-4">
          <p className="text-red-500">
            警告：此操作无法撤销。请谨慎操作。
          </p>
          <div>
            <p className="mb-2">
              输入 <strong className="text-red-600">{deleteTarget?.name}</strong> 以确认删除
            </p>
            <Input
              placeholder={`请输入房间名称: ${deleteTarget?.name}`}
              value={deleteConfirmName}
              onChange={(e) => setDeleteConfirmName(e.target.value)}
              autoFocus
            />
          </div>
        </div>
      </Modal>
    </div>
  )
}

export default Home

