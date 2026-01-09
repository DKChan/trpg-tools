import { Form, Input, Button, Card, message } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { authService } from '../services'
import { useAuthStore } from '../store/authStore'

function Login() {
  const navigate = useNavigate()
  const setAuth = useAuthStore((state) => state.setAuth)
  const [form] = Form.useForm()

  const onFinish = async (values: any) => {
    try {
      const response = await authService.login(values)
      if (response.data.code === 200) {
        const { user_id, email, nickname } = response.data.data
        setAuth({ user_id, email, nickname }, 'mock-token')
        message.success('登录成功')
        navigate('/')
      }
    } catch (error) {
      message.error('登录失败')
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <Card className="w-full max-w-md shadow-lg">
        <h1 className="text-2xl font-bold text-center mb-6">登录</h1>
        <Form
          form={form}
          name="login"
          onFinish={onFinish}
          autoComplete="off"
          size="large"
        >
          <Form.Item
            name="email"
            rules={[
              { required: true, message: '请输入邮箱' },
              { type: 'email', message: '请输入有效的邮箱地址' },
            ]}
          >
            <Input prefix={<UserOutlined />} placeholder="邮箱" />
          </Form.Item>

          <Form.Item
            name="password"
            rules={[{ required: true, message: '请输入密码' }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="密码" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" className="w-full">
              登录
            </Button>
          </Form.Item>

          <div className="text-center">
            <span className="text-gray-600">还没有账号？</span>
            <Button
              type="link"
              onClick={() => navigate('/register')}
              className="p-0 ml-1"
            >
              立即注册
            </Button>
          </div>
        </Form>
      </Card>
    </div>
  )
}

export default Login
