# OpenIM Chat Admin API 功能列表

## 一、管理员账户管理 (Admin Account Management)

*   `Login` (管理员登录): 管理员使用账户和密码进行登录。
*   `ChangePassword` (修改密码): 修改当前登录管理员的密码。
*   `AdminUpdateInfo` (更新管理员信息): 更新管理员的账户信息，如昵称、头像、等级等。
*   `GetAdminInfo` (获取管理员信息): 获取当前登录管理员的详细信息。
*   `AddAdminAccount` (添加管理员账户): 创建新的管理员账户。
*   `ChangeAdminPassword` (修改指定管理员密码): 修改指定管理员账户的密码。
*   `DelAdminAccount` (删除管理员账户): 删除指定的管理员账户。
*   `SearchAdminAccount` (搜索管理员账户): 根据条件搜索管理员账户列表。

## 二、用户与群组默认设置 (User and Group Default Settings)

*   `AddDefaultFriend` (添加默认好友): 设置用户注册后默认添加的好友。
*   `DelDefaultFriend` (删除默认好友): 移除用户注册后默认添加的好友。
*   `FindDefaultFriend` (查询默认好友列表): 获取当前的默认好友列表。
*   `SearchDefaultFriend` (搜索默认好友): 根据关键词搜索默认好友。
*   `AddDefaultGroup` (添加默认群组): 设置用户注册后默认加入的群组。
*   `DelDefaultGroup` (删除默认群组): 移除用户注册后默认加入的群组。
*   `FindDefaultGroup` (查询默认群组列表): 获取当前的默认群组列表。
*   `SearchDefaultGroup` (搜索默认群组): 根据关键词搜索默认群组。

## 三、邀请码管理 (Invitation Code Management)

*   `AddInvitationCode` (添加邀请码): 手动添加邀请码。
*   `GenInvitationCode` (生成邀请码): 批量生成邀请码。
*   `FindInvitationCode` (查询邀请码): 根据邀请码查询其状态和使用情况。
*   `UseInvitationCode` (使用邀请码): 标记邀请码已被特定用户使用（通常在注册时由系统内部调用）。
*   `DelInvitationCode` (删除邀请码): 删除指定的邀请码。
*   `SearchInvitationCode` (搜索邀请码): 根据条件（状态、用户ID、邀请码本身、关键词）搜索邀请码。

## 四、IP与用户登录限制 (IP and User Login Restrictions)

*   `SearchUserIPLimitLogin` (搜索用户IP登录限制): 查询特定用户登录IP的限制列表。
*   `AddUserIPLimitLogin` (添加用户IP登录限制): 限制特定用户只能通过指定的IP地址登录。
*   `DelUserIPLimitLogin` (删除用户IP登录限制): 移除对特定用户登录IP的限制。
*   `SearchIPForbidden` (搜索IP禁用列表): 查询已被禁用的IP地址列表（禁止注册/登录）。
*   `AddIPForbidden` (添加IP禁用): 禁止指定的IP地址进行注册或登录。
*   `DelIPForbidden` (删除IP禁用): 解除对指定IP地址的注册或登录禁用。
*   `CheckRegisterForbidden` (检查IP是否禁止注册): 检查某个IP地址是否被禁止注册。
*   `CheckLoginForbidden` (检查IP或用户是否禁止登录): 检查某个IP地址或用户是否被禁止登录。

## 五、用户账户操作 (User Account Operations)

*   `CancellationUser` (注销用户): 注销指定用户的账户。
*   `BlockUser` (封禁用户): 封禁指定用户的账户。
*   `UnblockUser` (解封用户): 解封指定用户的账户。
*   `SearchBlockUser` (搜索被封禁用户): 查询被封禁的用户列表。
*   `FindUserBlockInfo` (查询用户封禁信息): 获取特定用户的封禁状态和信息。

## 六、令牌管理 (Token Management)

*   `CreateToken` (创建令牌): 为指定用户（管理员或普通用户）创建认证令牌。
*   `ParseToken` (解析令牌): 解析令牌，获取令牌中包含的用户ID和用户类型。
*   `InvalidateToken` (使令牌无效): 使指定用户的所有令牌失效（强制下线）。
*   `GetUserToken` (获取用户令牌): 获取指定用户当前的有效令牌信息。

## 七、应用与配置管理 (Application and Configuration Management)

*   `AddApplet` (添加小程序): 添加新的小程序/应用信息。
*   `DelApplet` (删除小程序): 删除指定的小程序/应用信息。
*   `UpdateApplet` (更新小程序): 更新指定的小程序/应用信息。
*   `FindApplet` (查询小程序列表): 获取所有小程序/应用的信息列表。
*   `SearchApplet` (搜索小程序): 根据条件搜索小程序/应用。
*   `GetClientConfig` (获取客户端配置): 获取客户端初始化所需的配置信息。
*   `SetClientConfig` (设置客户端配置): 设置客户端初始化所需的配置信息。
*   `DelClientConfig` (删除客户端配置): 删除某些客户端初始化配置项。
*   `LatestApplicationVersion` (获取最新应用版本): 获取指定平台的最新应用版本信息。
*   `AddApplicationVersion` (添加应用版本): 添加新的应用版本信息。
*   `UpdateApplicationVersion` (更新应用版本): 更新指定的应用版本信息。
*   `DeleteApplicationVersion` (删除应用版本): 删除指定的应用版本信息。
*   `PageApplicationVersion` (分页查询应用版本): 分页获取应用版本信息。