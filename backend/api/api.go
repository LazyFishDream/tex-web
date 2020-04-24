package api

import (
    "github.com/labstack/echo"
    "github.com/yellia1989/tex-web/backend/api/gm"
)

func RegisterHandler(group *echo.Group) {
    group.POST("/login", UserLogin)                     // 登陆

    group.GET("/menu/list", MenuList)                   // 菜单列表
    group.POST("/menu/update", MenuUpdate)              // 更新菜单
    group.GET("/user/list", UserList)                   // 用户列表
    group.POST("/user/add", UserAdd)                    // 用户增加
    group.POST("/user/del", UserDel)                    // 用户删除
    group.POST("/user/update/role", UserUpdateRole)     // 用户角色变更
    group.POST("/user/update", UserUpdate)              // 用户更新
    group.POST("/user/pwd", UserPwd)                    // 用户改密码
    group.GET("/perm/list", PermList)                   // 权限列表
    group.POST("/perm/add", PermAdd)                    // 权限增加
    group.POST("/perm/del", PermDel)                    // 权限删除
    group.POST("/perm/update", PermUpdate)              // 权限编辑
    group.GET("/role/list", RoleList)                   // 角色列表
    group.POST("/role/add", RoleAdd)                    // 角色增加
    group.POST("/role/del", RoleDel)                    // 角色删除
    group.POST("/role/update", RoleUpdate)              // 角色更新

    group.POST("/gm/game/cmd", gm.GameCmd)              // 执行gm命令
    group.GET("/gm/zone/simplelist", gm.ZoneSimpleList) // 获取分区列表
    group.GET("/gm/zone/list", gm.ZoneList)             // 获取分区列表
    group.POST("/gm/zone/add", gm.ZoneAdd)              // 增加新分区
    group.POST("/gm/zone/del", gm.ZoneDel)              // 删除分区
    group.POST("/gm/zone/update", gm.ZoneUpdate)        // 更新分区
    group.POST("/gm/zone/version", gm.ZoneUpdateVersion)        // 批量更新分区版本号

    group.GET("/gm/channel/list", gm.ChannelList)             // 获取渠道列表
    group.POST("/gm/channel/add", gm.ChannelAdd)              // 增加新渠道
    group.POST("/gm/channel/del", gm.ChannelDel)              // 删除渠道
    group.POST("/gm/channel/update", gm.ChannelUpdate)        // 更新渠道

    group.GET("/gm/registry/list", gm.RegistryList)             // 获取registry列表
    group.POST("/gm/registry/add", gm.RegistryAdd)              // 增加registry
    group.POST("/gm/registry/del", gm.RegistryDel)              // 删除registry
}
