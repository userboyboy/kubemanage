package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/wonderivan/logger"
)

var Secret secret

type secret struct{}

func SecretRegister(router *gin.RouterGroup) {
	router.DELETE("/del", Secret.DeleteSecret)
	router.PUT("/update", Secret.UpdateSecret)
	router.GET("/list", Secret.GetSecretList)
	router.GET("/detail", Secret.GetSecretDetail)
}

// DeleteSecret 删除Secret
// ListPage godoc
// @Summary      删除Secret
// @Description  删除Secret
// @Tags         Secret
// @ID           /api/k8s/Secret/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Secret名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/Secret/del [delete]
func (s *secret) DeleteSecret(ctx *gin.Context) {
	params := &kubernetes.SecretNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := kube.Secret.DeleteSecrets(params.Name, params.NameSpace); err != nil {
		logger.Error("删除Secret失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdateSecret 更新Secret
// ListPage godoc
// @Summary      更新Secret
// @Description  更新Secret
// @Tags         Secret
// @ID           /api/k8s/secret/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/secret/update [put]
func (s *secret) UpdateSecret(ctx *gin.Context) {
	params := &kubernetes.SecretUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := kube.Secret.UpdateSecrets(params.Content, params.NameSpace); err != nil {
		logger.Error("更新Secret失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetSecretList 查看Secret列表
// ListPage godoc
// @Summary      查看Secret列表
// @Description  查看Secret列表
// @Tags         Secret
// @ID           /api/k8s/Secret/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/Secret/list [get]
func (s *secret) GetSecretList(ctx *gin.Context) {
	params := &kubernetes.SecretListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.Secret.GetSecrets(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取Secret列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetSecretDetail 获取Secret详情
// ListPage godoc
// @Summary      获取Secret详情
// @Description  获取Secret详情
// @Tags         Secret
// @ID           /api/k8s/Secret/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "Secret名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/Secret/detail [get]
func (s *secret) GetSecretDetail(ctx *gin.Context) {
	params := &kubernetes.SecretNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.Secret.GetSecretsDetail(params.Name, params.NameSpace)
	if err != nil {
		logger.Error("获取Secret详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
