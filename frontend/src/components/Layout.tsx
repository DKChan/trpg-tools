import { Outlet } from 'react-router-dom'
import { Layout as AntLayout, Menu } from 'antd'
import { HomeOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'

const { Header, Content } = AntLayout

function Layout() {
  const navigate = useNavigate()

  const menuItems = [
    {
      key: 'home',
      icon: <HomeOutlined />,
      label: '首页',
      onClick: () => navigate('/'),
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
      </Header>
      <Content className="p-6 bg-gray-50">
        <Outlet />
      </Content>
    </AntLayout>
  )
}

export default Layout
