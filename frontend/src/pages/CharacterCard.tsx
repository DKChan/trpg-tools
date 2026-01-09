import { useState } from 'react'
import { Card, Form, Input, InputNumber, Button, message } from 'antd'
import { SaveOutlined, ArrowLeftOutlined } from '@ant-design/icons'
import { useNavigate, useParams } from 'react-router-dom'

function CharacterCard() {
  const navigate = useNavigate()
  const { roomId, id } = useParams()
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)

  const handleSave = async (values: any) => {
    setLoading(true)
    try {
      message.success('保存成功')
    } catch (error) {
      message.error('保存失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <Button
          icon={<ArrowLeftOutlined />}
          onClick={() => navigate(`/rooms/${roomId}`)}
        >
          返回
        </Button>
        <Button
          type="primary"
          icon={<SaveOutlined />}
          onClick={() => form.submit()}
          loading={loading}
        >
          保存
        </Button>
      </div>

      <Form
        form={form}
        layout="vertical"
        onFinish={handleSave}
        initialValues={{
          level: 1,
          strength: 10,
          dexterity: 10,
          constitution: 10,
          intelligence: 10,
          wisdom: 10,
          charisma: 10,
          ac: 10,
          hp: 10,
          max_hp: 10,
          speed: 30,
          proficiency: 2,
        }}
      >
        <Card title="基本信息" className="mb-4">
          <Form.Item name="name" label="角色名称" rules={[{ required: true }]}>
            <Input placeholder="请输入角色名称" />
          </Form.Item>
          <Form.Item name="race" label="种族">
            <Input placeholder="请输入种族" />
          </Form.Item>
          <Form.Item name="class" label="职业">
            <Input placeholder="请输入职业" />
          </Form.Item>
          <Form.Item name="level" label="等级">
            <InputNumber min={1} max={20} className="w-full" />
          </Form.Item>
          <Form.Item name="background" label="背景">
            <Input.TextArea rows={2} placeholder="请输入背景" />
          </Form.Item>
          <Form.Item name="alignment" label="阵营">
            <Input placeholder="请输入阵营" />
          </Form.Item>
        </Card>

        <Card title="属性值" className="mb-4">
          <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
            <Form.Item name="strength" label="力量">
              <InputNumber min={1} max={30} className="w-full" />
            </Form.Item>
            <Form.Item name="dexterity" label="敏捷">
              <InputNumber min={1} max={30} className="w-full" />
            </Form.Item>
            <Form.Item name="constitution" label="体质">
              <InputNumber min={1} max={30} className="w-full" />
            </Form.Item>
            <Form.Item name="intelligence" label="智力">
              <InputNumber min={1} max={30} className="w-full" />
            </Form.Item>
            <Form.Item name="wisdom" label="感知">
              <InputNumber min={1} max={30} className="w-full" />
            </Form.Item>
            <Form.Item name="charisma" label="魅力">
              <InputNumber min={1} max={30} className="w-full" />
            </Form.Item>
          </div>
        </Card>

        <Card title="战斗属性" className="mb-4">
          <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
            <Form.Item name="ac" label="护甲等级">
              <InputNumber min={1} max={30} className="w-full" />
            </Form.Item>
            <Form.Item name="hp" label="当前生命值">
              <InputNumber min={0} max={1000} className="w-full" />
            </Form.Item>
            <Form.Item name="max_hp" label="最大生命值">
              <InputNumber min={1} max={1000} className="w-full" />
            </Form.Item>
            <Form.Item name="speed" label="速度">
              <InputNumber min={0} max={100} className="w-full" />
            </Form.Item>
            <Form.Item name="proficiency" label="熟练加值">
              <InputNumber min={0} max={10} className="w-full" />
            </Form.Item>
          </div>
        </Card>

        <Card title="装备">
          <Form.Item name="equipment" label="装备列表">
            <Input.TextArea rows={4} placeholder="请输入装备信息" />
          </Form.Item>
        </Card>
      </Form>
    </div>
  )
}

export default CharacterCard
