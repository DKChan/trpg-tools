# TRPG个人人物卡管理工具需求文档

## 1. 项目概述

### 1.1 项目背景
TRPG（桌面角色扮演游戏）玩家在进行游戏时，需要管理多个人物卡信息，包括属性、技能、装备等。传统的纸质或电子表格管理方式存在以下问题：
- 信息更新繁琐，容易出错
- 多个人物卡难以分类管理
- 不同战役的人物卡混在一起
- 无法快速切换不同规则系统的人物卡

### 1.2 项目目标
构建一个单机TRPG人物卡管理工具，为个人玩家提供：
- 便捷的人物卡创建和编辑功能
- 房间分类管理（用于区分不同战役）
- 支持多种TRPG规则系统
- 本地数据存储，无需网络连接

### 1.3 目标用户
- **TRPG玩家**：管理个人的人物卡库，按战役分类组织

### 1.4 项目定位
**个人工具**: 本工具专为个人使用设计，所有数据存储在本地，无需用户认证，不支持多人协作或实时同步。

## 2. 核心功能需求

### 2.1 房间管理（分类功能）

房间功能主要用于分类管理不同战役的人物卡，不是多用户协作空间。

#### 2.1.1 创建房间
- 设置房间名称（用于标识战役，如"冰风谷战役"）
- 选择TRPG规则系统（目前支持D&D 5e）
- 设置房间描述（可选，记录战役信息）
- 选择TRPG规则系统（初始仅支持DND5e）
- 设置房间密码（可选）
- 生成房间邀请码

#### 2.2.2 加入房间
- 玩家可以通过邀请码加入房间
- 输入房间密码（如需要）
- 验证房间是否已满
- 限制玩家数量（建议上限10人）

#### 2.2.3 房间列表
- 查看公开房间列表
- 显示房间基本信息（名称、DM、人数、规则系统）
- 搜索房间功能

#### 2.2.4 房间管理
- DM可以踢出玩家
- DM可以转让DM权限
- DM可以解散房间
- 玩家可以主动退出房间

### 2.3 人物卡管理

#### 2.3.1 创建人物卡（DND5e）
- 基本信息：
  - 角色名称
  - 种族
  - 职业
  - 等级
  - 背景故事
  - 阵营
- 属性值：
  - 力量（Strength）
  - 敏捷（Dexterity）
  - 体质（Constitution）
  - 智力（Intelligence）
  - 感知（Wisdom）
  - 魅力（Charisma）
- 战斗属性：
  - 护甲等级（AC）
  - 生命值（HP）
  - 速度
  - 先攻加值
  - 熟练加值
- 技能熟练度：
  - 运动、杂技、隐匿、巧手、奥秘、历史、调查、自然、宗教、察觉、求生、说服、欺瞒、威吓、表演
- 豁免熟练度：
  - 力量豁免、敏捷豁免、体质豁免、智力豁免、感知豁免、魅力豁免
- 装备：
  - 武器列表
  - 护甲列表
  - 物品列表
- 法术（如适用）：
  - 已知法术
  - 法术位
  - 每日法术使用情况

#### 2.3.2 编辑人物卡
- 修改人物卡所有字段
- 实时保存
- 历史版本记录（可选）

#### 2.3.3 查看人物卡
- 玩家只能查看自己的人物卡
- DM可以查看房间内所有玩家的人物卡
- 支持详细视图和简化视图

#### 2.3.4 删除人物卡
- 玩家可以删除自己的人物卡
- DM不能删除玩家的人物卡
- 删除前二次确认

### 2.4 实时同步

#### 2.4.1 WebSocket连接
- 建立房间内的WebSocket连接
- 自动重连机制
- 心跳检测

#### 2.4.2 实时更新
- 人物卡属性变更实时同步
- 房间成员变动实时通知
- DM操作实时通知

#### 2.4.3 消息通知
- 玩家加入/退出房间通知
- DM操作通知
- 系统消息通知

### 2.5 权限管理

#### 2.5.1 角色权限
- **DM权限**：
  - 创建/解散房间
  - 踢出玩家
  - 转让DM权限
  - 查看所有玩家人物卡
  - 修改房间设置
  - 发送系统通知
- **玩家权限**：
  - 创建/编辑/删除自己的人物卡
  - 查看自己的人物卡
  - 主动退出房间
  - 查看房间信息

#### 2.5.2 权限验证
- 所有API接口进行权限验证
- 前端路由权限控制
- 敏感操作二次确认

## 3. 非功能需求

### 3.1 性能要求
- 页面加载时间 < 2秒
- API响应时间 < 500ms
- WebSocket消息延迟 < 100ms
- 支持至少100个并发房间
- 每个房间支持最多10名玩家

### 3.2 安全要求
- 所有API使用HTTPS
- 密码使用bcrypt加密（cost >= 10）
- JWT Token过期时间设置为24小时
- 防止SQL注入
- 防止XSS攻击
- 防止CSRF攻击
- 敏感信息不记录日志

### 3.3 可用性要求
- 系统可用性 >= 99.5%
- 数据备份频率：每日一次
- 支持数据恢复

### 3.4 可扩展性要求
- 支持水平扩展
- 支持多种TRPG规则系统（DND5e、COC、PF等）
- 支持插件化扩展

### 3.5 兼容性要求
- 支持主流浏览器（Chrome、Firefox、Safari、Edge最新版本）
- 支持移动端浏览器
- 响应式设计，适配不同屏幕尺寸

## 4. 数据模型设计

### 4.1 用户模型（users）
```go
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Email     string    `gorm:"uniqueIndex;not null"`
    Password  string    `gorm:"not null"`
    Nickname  string    `gorm:"not null"`
    Avatar    string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 4.2 房间模型（rooms）
```go
type Room struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"not null"`
    Description string
    RuleSystem  string    `gorm:"not null;default:'DND5e'"`
    Password    string
    InviteCode  string    `gorm:"uniqueIndex;not null"`
    DMID        uint      `gorm:"not null"`
    MaxPlayers  int       `gorm:"not null;default:10"`
    IsPublic    bool      `gorm:"not null;default:true"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### 4.3 房间成员模型（room_members）
```go
type RoomMember struct {
    ID        uint      `gorm:"primaryKey"`
    RoomID    uint      `gorm:"not null"`
    UserID    uint      `gorm:"not null"`
    Role      string    `gorm:"not null;default:'player'"`
    JoinedAt  time.Time
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 4.4 人物卡模型（character_cards）
```go
type CharacterCard struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"not null"`
    RoomID      uint      `gorm:"not null"`
    Name        string    `gorm:"not null"`
    Race        string
    Class       string
    Level       int       `gorm:"default:1"`
    Background  string
    Alignment   string
    Strength    int       `gorm:"default:10"`
    Dexterity   int       `gorm:"default:10"`
    Constitution int      `gorm:"default:10"`
    Intelligence int       `gorm:"default:10"`
    Wisdom      int       `gorm:"default:10"`
    Charisma    int       `gorm:"default:10"`
    AC          int       `gorm:"default:10"`
    HP          int       `gorm:"default:10"`
    MaxHP       int       `gorm:"default:10"`
    Speed       int       `gorm:"default:30"`
    Proficiency int       `gorm:"default:2"`
    Skills      string    `gorm:"type:jsonb"`
    Saves       string    `gorm:"type:jsonb"`
    Equipment   string    `gorm:"type:jsonb"`
    Spells      string    `gorm:"type:jsonb"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### 4.5 游戏会话模型（game_sessions）
```go
type GameSession struct {
    ID        uint      `gorm:"primaryKey"`
    RoomID    uint      `gorm:"not null"`
    StartTime time.Time
    EndTime   *time.Time
    Notes     string    `gorm:"type:text"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

## 5. API接口设计

### 5.1 用户相关接口
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `GET /api/v1/user/profile` - 获取用户信息
- `PUT /api/v1/user/profile` - 更新用户信息
- `PUT /api/v1/user/password` - 修改密码

### 5.2 房间相关接口
- `POST /api/v1/rooms` - 创建房间
- `GET /api/v1/rooms` - 获取房间列表
- `GET /api/v1/rooms/:id` - 获取房间详情
- `POST /api/v1/rooms/:id/join` - 加入房间
- `POST /api/v1/rooms/:id/leave` - 退出房间
- `DELETE /api/v1/rooms/:id` - 解散房间
- `PUT /api/v1/rooms/:id/members/:userId/kick` - 踢出玩家
- `PUT /api/v1/rooms/:id/transfer-dm` - 转让DM权限

### 5.3 人物卡相关接口
- `POST /api/v1/rooms/:roomId/characters` - 创建人物卡
- `GET /api/v1/rooms/:roomId/characters` - 获取房间内所有人物卡
- `GET /api/v1/rooms/:roomId/characters/:id` - 获取人物卡详情
- `PUT /api/v1/rooms/:roomId/characters/:id` - 更新人物卡
- `DELETE /api/v1/rooms/:roomId/characters/:id` - 删除人物卡

### 5.4 WebSocket接口
- `WS /api/v1/ws/rooms/:roomId` - 房间WebSocket连接

## 6. 技术架构

### 6.1 后端技术栈
- **Web框架**：Gin
- **ORM**：GORM
- **数据库**：PostgreSQL
- **认证**：JWT
- **实时通信**：WebSocket（gorilla/websocket）
- **日志**：zap
- **配置管理**：viper

### 6.2 前端技术栈
- **框架**：React 18
- **构建工具**：Vite
- **状态管理**：Zustand
- **路由**：React Router v6
- **UI库**：Ant Design
- **API请求**：Axios
- **样式**：TailwindCSS
- **类型检查**：TypeScript 5
- **WebSocket客户端**：原生WebSocket API

## 7. 界面设计要求

### 7.1 登录/注册页面
- 简洁的登录表单
- 邮箱和密码输入
- 记住我选项
- 注册链接

### 7.2 主页面
- 顶部导航栏（用户信息、退出登录）
- 房间列表（卡片式布局）
- 创建房间按钮
- 搜索框

### 7.3 房间详情页面
- 房间信息展示
- 成员列表
- 玩家人物卡列表
- DM操作按钮（仅DM可见）

### 7.4 人物卡编辑页面
- 分组展示人物卡信息
- 实时保存提示
- 预览功能

## 8. 后续扩展功能

### 8.1 多规则系统支持
- COC（克苏鲁的呼唤）
- PF（开拓者）
- 自定义规则系统

### 8.2 人物卡模板系统
- 预设人物卡模板
- 自定义模板
- 模板分享

### 8.3 战斗管理
- 战斗轮次管理
- 伤害计算
- 状态效果追踪
- 战斗日志

### 8.4 骰子系统
- 在线掷骰子
- 支持多种骰子类型
- 骰子结果记录
- 骰子宏

### 8.5 地图系统
- 在线地图共享
- 标记和注释
- 迷雾系统

### 8.6 语音/视频通话
- 集成语音通话功能
- 集成视频通话功能
- 屏幕共享

## 9. 验收标准

### 9.1 功能验收
- 所有核心功能正常运行
- 权限控制正确
- 实时同步稳定
- 数据一致性保证

### 9.2 性能验收
- 满足性能要求指标
- 压力测试通过
- 负载测试通过

### 9.3 安全验收
- 安全漏洞扫描通过
- 渗透测试通过
- 代码审计通过

### 9.4 兼容性验收
- 主流浏览器测试通过
- 移动端测试通过
- 不同屏幕尺寸适配测试通过

## 10. 项目里程碑

### 10.1 第一阶段（MVP）
- 用户注册登录
- 房间创建和加入
- 人物卡创建和编辑（DND5e）
- 基础实时同步

### 10.2 第二阶段
- 完善人物卡功能
- 权限管理优化
- UI/UX优化

### 10.3 第三阶段
- 多规则系统支持
- 战斗管理功能
- 骰子系统

### 10.4 第四阶段
- 地图系统
- 语音/视频通话
- 性能优化

## 11. 风险评估

### 11.1 技术风险
- WebSocket连接稳定性
- 实时同步数据一致性
- 高并发性能瓶颈

### 11.2 业务风险
- 用户需求变更
- 竞品功能迭代
- 规则系统复杂度

### 11.3 应对措施
- 技术预研和POC验证
- 敏捷开发，快速迭代
- 定期用户反馈收集
- 模块化设计，便于扩展

## 12. 附录

### 12.1 参考资料
- DND5e规则书
- TRPG社区最佳实践
- WebSocket最佳实践

### 12.2 术语表
- **DM**：Dungeon Master，地下城主，游戏的主持人
- **TRPG**：Tabletop Role-Playing Game，桌面角色扮演游戏
- **DND5e**：Dungeons & Dragons 5th Edition，龙与地下城第五版
- **COC**：Call of Cthulhu，克苏鲁的呼唤
- **PF**：Pathfinder，开拓者

---

**文档版本**：v1.0.0  
**创建日期**：2026-01-05  
**最后更新**：2026-01-05
