# TRPG-Sync 前端测试指南

## 测试框架

- **Vitest**: 现代化的测试框架
- **React Testing Library**: React 组件测试
- **@testing-library/user-event**: 用户交互模拟
- **jsdom**: 浏览器环境模拟

## 测试结构

```
frontend/
├── src/
│   ├── test/
│   │   ├── setup.ts                    # 测试环境配置
│   │   ├── renderWithRouter.tsx        # Router 渲染辅助
│   │   ├── mockService.ts              # API mock
│   │   └── testData.ts                # 测试数据
│   ├── store/
│   │   └── authStore.test.ts
│   ├── services/
│   │   ├── api.test.ts
│   │   └── index.test.ts
│   └── pages/
│       ├── Login.test.tsx
│       ├── Register.test.tsx
│       └── Home.test.tsx
├── vitest.config.ts                    # Vitest 配置
└── package.json                         # 测试脚本
```

## 运行测试

### 运行所有测试
```bash
cd frontend
npm test
```

### 运行测试并监听文件变化
```bash
npm test:ui
```

### 运行测试一次
```bash
npm test:run
```

### 生成测试覆盖率报告
```bash
npm test:coverage
```

## 测试覆盖范围

### 已实现的测试

#### 高优先级模块
- ✅ `store/authStore.ts` - 状态管理
- ✅ `services/api.ts` - API 拦截器
- ✅ `services/index.ts` - 所有 API 服务方法
- ✅ `pages/Login.tsx` - 登录页面
- ✅ `pages/Register.tsx` - 注册页面
- ✅ `pages/Home.tsx` - 房间列表页面

### 测试覆盖率目标

- 核心业务逻辑: ≥ 90%
- 组件渲染: ≥ 70%
- 状态管理: ≥ 80%
- 整体覆盖率: ≥ 70%

## 测试辅助工具

### renderWithRouter
```typescript
import { renderWithRouter } from '../test/renderWithRouter'

renderWithRouter(<MyComponent />)
```

### 模拟数据
```typescript
import { mockUser, mockRoom, mockCharacter } from '../test/testData'

const user = mockUser
const room = mockRoom
```

### Mock API
```typescript
import { mockApi, mockSuccessResponse } from '../test/mockService'

mockApi.get.mockResolvedValue(mockSuccessResponse(data))
```

## 测试最佳实践

1. **遵循 "The Testing Library" 哲学**: 测试用户行为而非实现细节
2. **使用 `waitFor` 处理异步操作**: 确保断言在异步操作完成后执行
3. **Mock 外部依赖**: 使用 vi.mock 隔离 API 调用和路由
4. **测试所有用户交互**: 输入、点击、表单提交等
5. **测试边界条件和错误场景**: 网络错误、验证失败等

## 添加新测试

1. 在对应目录创建 `*.test.tsx` 或 `*.test.ts` 文件
2. 导入必要的测试工具和组件
3. 使用 `describe` 组织测试用例
4. 使用 `it` 编写单个测试
5. Mock 外部依赖（API、路由等）
6. 编写断言验证预期行为
7. 运行 `npm test` 验证测试通过

## 示例测试

### 组件测试
```typescript
import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import MyComponent from './MyComponent'

describe('MyComponent', () => {
  it('renders correctly', () => {
    render(<MyComponent />)
    expect(screen.getByText('Hello')).toBeInTheDocument()
  })
})
```

### 服务测试
```typescript
import { describe, it, expect, vi } from 'vitest'
import { myService } from './services'

describe('MyService', () => {
  it('calls correct API', () => {
    myService.getData()
    expect(api.get).toHaveBeenCalledWith('/data')
  })
})
```

## 已知问题

1. **npm 命令不可用**: 需要手动安装依赖
2. **部分组件测试未完成**: Layout, RoomDetail, CharacterCard 等
3. **集成测试缺失**: 端到端用户流程测试

## 下一步

1. 完成剩余组件的测试
2. 添加端到端测试（Playwright 或 Cypress）
3. 配置 CI/CD 集成
4. 持续监控测试覆盖率
