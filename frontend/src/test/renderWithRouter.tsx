import { render } from '@testing-library/react'
import { BrowserRouter } from 'react-router-dom'

/**
 * 渲染组件并包裹在 Router 中
 */
export function renderWithRouter(ui: React.ReactElement) {
  return render(<BrowserRouter>{ui}</BrowserRouter>)
}

/**
 * 渲染组件并包裹在 Router 中，指定路径
 */
export function renderWithRouterAndPath(ui: React.ReactElement, path: string) {
  window.history.pushState({}, 'Test page', path)
  return render(<BrowserRouter>{ui}</BrowserRouter>)
}
