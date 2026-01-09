import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Layout from './components/Layout'
import Login from './pages/Login'
import Register from './pages/Register'
import Home from './pages/Home'
import RoomDetail from './pages/RoomDetail'
import CharacterCard from './pages/CharacterCard'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="rooms/:id" element={<RoomDetail />} />
          <Route path="rooms/:roomId/characters/:id" element={<CharacterCard />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
