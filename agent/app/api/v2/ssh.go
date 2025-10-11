package v2

import (
	"encoding/base64"

	"github.com/1Panel-dev/1Panel/agent/app/api/v2/helper"
	"github.com/1Panel-dev/1Panel/agent/app/dto"
	"github.com/gin-gonic/gin"
)

// @Tags SSH
// @Summary Load host SSH setting info
// @Success 200 {object} dto.SSHInfo
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/search [post]
func (b *BaseApi) GetSSHInfo(c *gin.Context) {
	info, err := sshService.GetSSHInfo()
	if err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.SuccessWithData(c, info)
}

// @Tags SSH
// @Summary Operate SSH
// @Accept json
// @Param request body dto.Operate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/operate [post]
// @x-panel-log {"bodyKeys":["operation"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"[operation] SSH ","formatEN":"[operation] SSH"}
func (b *BaseApi) OperateSSH(c *gin.Context) {
	var req dto.Operate
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := sshService.OperateSSH(req.Operation); err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.Success(c)
}

// @Tags SSH
// @Summary Update host SSH setting
// @Accept json
// @Param request body dto.SSHUpdate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/update [post]
// @x-panel-log {"bodyKeys":["key","value"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"修改 SSH 配置 [key] => [value]","formatEN":"update SSH setting [key] => [value]"}
func (b *BaseApi) UpdateSSH(c *gin.Context) {
	var req dto.SSHUpdate
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := sshService.Update(req); err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.Success(c)
}

// @Tags SSH
// @Summary Update host SSH setting by file
// @Accept json
// @Param request body dto.SSHConf true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/conffile/update [post]
// @x-panel-log {"bodyKeys":[],"paramKeys":[],"BeforeFunctions":[],"formatZH":"修改 SSH 配置文件","formatEN":"update SSH conf"}
func (b *BaseApi) UpdateSSHByfile(c *gin.Context) {
	var req dto.SSHConf
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := sshService.UpdateByFile(req.File); err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.Success(c)
}

// @Tags SSH
// @Summary Generate host SSH secret
// @Accept json
// @Param request body dto.CreateRootCert true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/cert [post]
// @x-panel-log {"bodyKeys":[],"paramKeys":[],"BeforeFunctions":[],"formatZH":"生成 SSH 密钥 ","formatEN":"generate SSH secret"}
func (b *BaseApi) CreateRootCert(c *gin.Context) {
	var req dto.CreateRootCert
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}
	if len(req.PassPhrase) != 0 {
		passPhrase, err := base64.StdEncoding.DecodeString(req.PassPhrase)
		if err != nil {
			helper.BadRequest(c, err)
			return
		}
		req.PassPhrase = string(passPhrase)
	}
	if len(req.PrivateKey) != 0 {
		privateKey, err := base64.StdEncoding.DecodeString(req.PrivateKey)
		if err != nil {
			helper.BadRequest(c, err)
			return
		}
		req.PrivateKey = string(privateKey)
	}
	if len(req.PublicKey) != 0 {
		publicKey, err := base64.StdEncoding.DecodeString(req.PublicKey)
		if err != nil {
			helper.BadRequest(c, err)
			return
		}
		req.PublicKey = string(publicKey)
	}

	if err := sshService.CreateRootCert(req); err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.Success(c)
}

// @Tags SSH
// @Summary Sycn host SSH secret
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/cert/sync [post]
// @x-panel-log {"bodyKeys":[],"paramKeys":[],"BeforeFunctions":[],"formatZH":"同步 SSH 密钥 ","formatEN":"sync SSH secret"}
func (b *BaseApi) SyncRootCert(c *gin.Context) {
	if err := sshService.SyncRootCert(); err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.Success(c)
}

// @Tags SSH
// @Summary Load host SSH secret
// @Accept json
// @Param request body dto.SearchWithPage true "request"
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/cert/search [post]
func (b *BaseApi) SearchRootCert(c *gin.Context) {
	var req dto.SearchWithPage
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	total, data, err := sshService.SearchRootCerts(req)
	if err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.SuccessWithData(c, dto.PageResult{
		Total: total,
		Items: data,
	})
}

// @Tags SSH
// @Summary Delete host SSH secret
// @Accept json
// @Param request body dto.ForceDelete true "request"
// @Success 200
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/cert/delete [post]
// @x-panel-log {"bodyKeys":[],"paramKeys":[],"BeforeFunctions":[],"formatZH":"删除 SSH 密钥 ","formatEN":"delete SSH secret"}
func (b *BaseApi) DeleteRootCert(c *gin.Context) {
	var req dto.ForceDelete
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	if err := sshService.DeleteRootCerts(req); err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.Success(c)
}

// @Tags SSH
// @Summary Load host SSH logs
// @Accept json
// @Param request body dto.SearchSSHLog true "request"
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/log [post]
func (b *BaseApi) LoadSSHLogs(c *gin.Context) {
	var req dto.SearchSSHLog
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	total, data, err := sshService.LoadLog(c, req)
	if err != nil {
		helper.InternalServer(c, err)
		return
	}

	helper.SuccessWithData(c, dto.PageResult{
		Total: total,
		Items: data,
	})
}

// @Tags SSH
// @Summary Export host SSH logs
// @Accept json
// @Param request body dto.SearchSSHLog true "request"
// @Success 200 {string} path
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/log/export [post]
func (b *BaseApi) ExportSSHLogs(c *gin.Context) {
	var req dto.SearchSSHLog
	if err := helper.CheckBindAndValidate(&req, c); err != nil {
		return
	}
	tmpFile, err := sshService.ExportLog(c, req)
	if err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.SuccessWithData(c, tmpFile)
}

// @Tags SSH
// @Summary Load host SSH conf
// @Success 200 {string} conf
// @Security ApiKeyAuth
// @Security Timestamp
// @Router /hosts/ssh/conf [get]
func (b *BaseApi) LoadSSHConf(c *gin.Context) {
	data, err := sshService.LoadSSHConf()
	if err != nil {
		helper.InternalServer(c, err)
		return
	}
	helper.SuccessWithData(c, data)
}
