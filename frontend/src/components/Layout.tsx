import { Outlet } from 'react-router-dom'
import { Layout as AntLayout, Menu, Button, Avatar, Dropdown } from 'antd'
import { UserOutlined, LogoutOutlined, HomeOutlined } from '@ant-design/icons'
import { useAuthStore } from '../store/authStore'
import { useNavigate } from 'react-router-dom'

const { Header, Content } = AntLayout

function Layout() {
  const { user, logout } = useAuthStore()
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  const menuItems = [
    {
      key: 'home',
      icon: <HomeOutlined />,
      label: '首页',
      onClick: () => navigate('/'),
    },
  ]

  const userMenuItems = [
    {
      key: 'logout',
      icon: <LogoutOutlined />,
      label: '退出登录',
      onClick: handleLogout,
    },
  ]

  return (
    <AntLayout className="min-h-screen">
      <Header className="flex items-center justify-between bg-white shadow-md px-6">
        <div className="flex items-center gap-4">
          <div className="text-xl font-bold text-blue-600">TRPG Sync</div>
          <Menu
            theme="light"
            mode="horizontal"
            selectedKeys={[]}
            items={menuItems}
            className="border-none flex-1"
          />
        </div>
        <div className="flex items-center gap-4">
          <span className="text-gray-700">{user?.nickname}</span>
          <Dropdown menu={{ items: userMenuItems }} placement="bottomRight">
            <Avatar icon={<UserOutlined />} className="cursor-pointer" />
          </Dropdown>
        </div>
      </Header>
      <Content className="p-6 bg-gray-50">
        <Outlet />
      </Content>
    </AntLayout>
  )
}

export default Layout
