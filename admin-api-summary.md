# OpenIM Chat Admin API Summary

This document summarizes all admin API endpoints available in the OpenIM Go server, including their URLs and request parameters.

## 1. Admin Account Management

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Login | `/account/login` | Admin login | `account`, `password` |
| Update Admin Info | `/account/update` | Modify admin information | `nickname`, `faceURL` |
| Get Admin Info | `/account/info` | Get admin information | - |
| Change Admin Password | `/account/change_password` | Change admin password | `currentPassword`, `newPassword` |
| Add Admin Account | `/account/add_admin` | Add new admin account | `account`, `password`, `faceURL`, `nickname` |
| Add User Account | `/account/add_user` | Add new user account | `user` object containing user details |
| Delete Admin Account | `/account/del_admin` | Delete admin account | `userID` |
| Search Admin Accounts | `/account/search` | Get admin list | `keyword`, `pagination` parameters |

## 2. User Import

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Import Users by JSON | `/user/import/json` | Import users via JSON | `users` array of user objects |
| Import Users by XLSX | `/user/import/xlsx` | Import users via Excel file | Excel file upload |
| Get Batch Import Template | `/user/import/xlsx` (GET) | Download import template | - |

## 3. Registration Control

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Get Allow Register | `/user/allow_register/get` | Check if registration is allowed | - |
| Set Allow Register | `/user/allow_register/set` | Enable/disable registration | `allow` boolean |

## 4. Default User Settings

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Add Default Friend | `/default/user/add` | Add default friend at registration | `userID` |
| Delete Default Friend | `/default/user/del` | Delete default friend | `userID` |
| Find Default Friends | `/default/user/find` | Get default friend list | - |
| Search Default Friends | `/default/user/search` | Search default friends | `keyword` |

## 5. Default Group Settings

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Add Default Group | `/default/group/add` | Add default group at registration | `groupID` |
| Delete Default Group | `/default/group/del` | Delete default group | `groupID` |
| Find Default Groups | `/default/group/find` | Get default group list | - |
| Search Default Groups | `/default/group/search` | Search default groups | `keyword` |

## 6. Invitation Code Management

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Add Invitation Code | `/invitation_code/add` | Add invitation code | `code`, `userID` (optional) |
| Generate Invitation Codes | `/invitation_code/gen` | Generate multiple codes | `count` |
| Delete Invitation Code | `/invitation_code/del` | Delete invitation code | `codes` array |
| Search Invitation Codes | `/invitation_code/search` | Search invitation codes | `keyword`, `status`, `pagination` parameters |

## 7. IP Restrictions

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Add IP Forbidden | `/forbidden/ip/add` | Block IP for registration/login | `ip`, `limitType` |
| Delete IP Forbidden | `/forbidden/ip/del` | Unblock IP | `ip` |
| Search IP Forbidden | `/forbidden/ip/search` | List blocked IPs | `keyword`, `pagination` parameters |

## 8. User IP Restrictions

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Add User IP Limit | `/forbidden/user/add` | Limit user login to specific IPs | `userID`, `ips` array |
| Delete User IP Limit | `/forbidden/user/del` | Remove user IP restriction | `userID`, `ip` |
| Search User IP Limits | `/forbidden/user/search` | List user IP restrictions | `userID`, `pagination` parameters |

## 9. Applet Management

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Add Applet | `/applet/add` | Add new applet | `name`, `appID`, `icon`, `url`, `status`, etc. |
| Delete Applet | `/applet/del` | Delete applet | `appID` |
| Update Applet | `/applet/update` | Modify applet | `appID` and updated fields |
| Search Applets | `/applet/search` | Search applets | `keyword`, `pagination` parameters |

## 10. User Blocking

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Block User | `/block/add` | Block user | `userID`, `reason` |
| Unblock User | `/block/del` | Unblock user | `userID` |
| Search Blocked Users | `/block/search` | List blocked users | `keyword`, `pagination` parameters |

## 11. User Management

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Reset User Password | `/user/password/reset` | Reset user password | `userID`, `newPassword` |

## 12. Client Configuration

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Get Client Config | `/client_config/get` | Get client initialization config | `keys` array (optional) |
| Set Client Config | `/client_config/set` | Set client initialization config | `configs` key-value pairs |
| Delete Client Config | `/client_config/del` | Delete client config items | `keys` array |

## 13. Statistics

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| New User Count | `/statistic/new_user_count` | Get new user statistics | `start`, `end` date range |
| Login User Count | `/statistic/login_user_count` | Get login statistics | `start`, `end` date range |

## 14. Application Version Management

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Add Application Version | `/application/add_version` | Add new app version | `version`, `platform`, `forceUpdate`, etc. |
| Update Application Version | `/application/update_version` | Update app version | `versionID` and updated fields |
| Delete Application Version | `/application/delete_version` | Delete app version | `versionID` |
| Latest Application Version | `/application/latest_version` | Get latest version | `platform` |
| Page Application Versions | `/application/page_versions` | List versions with pagination | `platform`, `pagination` parameters |

## 15. System Configuration Management

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Get Config List | `/config/get_config_list` | Get available configs | - |
| Get Config | `/config/get_config` | Get specific config | `name` |
| Set Config | `/config/set_config` | Update config | `name`, `content` |
| Set Multiple Configs | `/config/set_configs` | Update multiple configs | Array of config objects |
| Reset Config | `/config/reset_config` | Reset config to default | `name` |
| Get Enable Config Manager | `/config/get_enable_config_manager` | Check if config manager is enabled | - |
| Set Enable Config Manager | `/config/set_enable_config_manager` | Enable/disable config manager | `enable` boolean |

## 16. System Control

| Endpoint | URL | Description | Parameters |
|----------|-----|-------------|------------|
| Restart | `/restart` | Restart services | `services` array |