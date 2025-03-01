package kubeController

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dto/kubernetes"
	"github.com/noovertime7/kubemanage/middleware"
	"github.com/noovertime7/kubemanage/pkg/core/kubemanage/v1/kube"
	"github.com/wonderivan/logger"
)

var PersistentVolumeClaim persistentVolumeClaim

type persistentVolumeClaim struct{}

func PersistentVolumeClaimRegister(router *gin.RouterGroup) {
	router.DELETE("/del", PersistentVolumeClaim.DeletePersistentVolumeClaim)
	router.PUT("/update", PersistentVolumeClaim.UpdatePersistentVolumeClaim)
	router.GET("/list", PersistentVolumeClaim.GetPersistentVolumeClaimList)
	router.GET("/detail", PersistentVolumeClaim.GetPersistentVolumeClaimDetail)
}

// DeletePersistentVolumeClaim 删除PersistentVolumeClaim
// ListPage godoc
// @Summary      删除PersistentVolumeClaim
// @Description  删除PersistentVolumeClaim
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/del
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "PersistentVolumeClaim名称"
// @Param        namespace    query  string  true  "命名空间"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "删除成功}"
// @Router       /api/k8s/persistentvolumeclaim/del [delete]
func (s *persistentVolumeClaim) DeletePersistentVolumeClaim(ctx *gin.Context) {
	params := &kubernetes.PersistentVolumeClaimNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := kube.PersistentVolumeClaim.DeletePersistentVolumeClaim(params.Name, params.NameSpace); err != nil {
		logger.Error("删除PersistentVolumeClaim失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}

// UpdatePersistentVolumeClaim 更新PersistentVolumeClaim
// ListPage godoc
// @Summary      更新PersistentVolumeClaim
// @Description  更新PersistentVolumeClaim
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/update
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "无状态控制器名称"
// @Param        namespace  query  string  true  "命名空间"
// @Param        content    query  string  true  "更新内容"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": "更新成功}"
// @Router       /api/k8s/persistentvolumeclaim/update [put]
func (s *persistentVolumeClaim) UpdatePersistentVolumeClaim(ctx *gin.Context) {
	params := &kubernetes.PersistentVolumeClaimUpdateInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	if err := kube.PersistentVolumeClaim.UpdatePersistentVolumeClaim(params.Content, params.NameSpace); err != nil {
		logger.Error("更新PersistentVolumeClaim失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "更新成功")
}

// GetPersistentVolumeClaimList 查看PersistentVolumeClaim列表
// ListPage godoc
// @Summary      查看PersistentVolumeClaim列表
// @Description  查看PersistentVolumeClaim列表
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/list
// @Accept       json
// @Produce      json
// @Param        filter_name  query  string  false  "过滤"
// @Param        namespace  query  string  false  "命名空间"
// @Param        page         query  int     false  "页码"
// @Param        limit        query  int     false  "分页限制"
// @Success       200  {object}  middleware.Response"{"code": 200, msg="","data": }"
// @Router       /api/k8s/persistentvolumeclaim/list [get]
func (s *persistentVolumeClaim) GetPersistentVolumeClaimList(ctx *gin.Context) {
	params := &kubernetes.PersistentVolumeClaimListInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.PersistentVolumeClaim.GetPersistentVolumeClaims(params.FilterName, params.NameSpace, params.Limit, params.Page)
	if err != nil {
		logger.Error("获取PersistentVolumeClaim列表失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

// GetPersistentVolumeClaimDetail 获取PersistentVolumeClaim详情
// ListPage godoc
// @Summary      获取PersistentVolumeClaim详情
// @Description  获取PersistentVolumeClaim详情
// @Tags         PersistentVolumeClaim
// @ID           /api/k8s/persistentvolumeclaim/detail
// @Accept       json
// @Produce      json
// @Param        name       query  string  true  "PersistentVolumeClaim名称"
// @Param        namespace  query  string  true  "命名空间"
// @Success      200        {object}  middleware.Response"{"code": 200, msg="","data":v1.Deployment }"
// @Router       /api/k8s/persistentvolumeclaim/detail [get]
func (s *persistentVolumeClaim) GetPersistentVolumeClaimDetail(ctx *gin.Context) {
	params := &kubernetes.PersistentVolumeClaimNameNS{}
	if err := params.BindingValidParams(ctx); err != nil {
		logger.Error("绑定参数失败:", err)
		middleware.ResponseError(ctx, 20001, err)
		return
	}
	data, err := kube.PersistentVolumeClaim.GetPersistentVolumeClaimDetail(params.Name, params.NameSpace)
	if err != nil {
		logger.Error("获取PersistentVolumeClaim详情失败", err)
		middleware.ResponseError(ctx, 20002, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}
