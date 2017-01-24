package api

import "github.com/humpback/humpback-center/api/request"
import "github.com/humpback/humpback-center/api/response"
import "github.com/humpback/gounits/logger"

import (
	"net/http"
)

func getClusterGroups(c *Context) error {

	logger.INFO("[#api#] %s resolve getclustergroups request successed.", c.ID)
	groups := c.Cluster.GetGroups()
	logger.INFO("[#api#] %s get cluster groups %d.", c.ID, len(groups))
	resp := response.NewClusterGroupsResponse(groups)
	result := &response.ResponseResult{ResponseID: c.ID}
	result.SetError(request.RequestSuccessed, request.ErrRequestSuccessed, "cluster groups")
	result.SetResponse(resp)
	return c.JSON(http.StatusOK, result)
}

func getClusterGroup(c *Context) error {

	result := &response.ResponseResult{ResponseID: c.ID}
	req, err := request.ResolveClusterGroupRequest(c.Request())
	if err != nil {
		logger.ERROR("[#api#] %s resolve getclustergroup request faild, %s", c.ID, err.Error())
		result.SetError(request.RequestInvalid, request.ErrRequestInvalid, err.Error())
		return c.JSON(http.StatusBadRequest, result)
	}

	logger.INFO("[#api#] %s resolve getclustergroup request successed.", c.ID)
	group := c.Cluster.GetGroup(req.GroupID)
	if group == nil {
		logger.ERROR("[#api#] %s get cluster group %s not found.", c.ID, req.GroupID)
		result.SetError(request.RequestNotFound, request.ErrRequestNotFound, req.GroupID+" not found")
		return c.JSON(http.StatusNotFound, result)
	}
	logger.INFO("[#api#] %s get cluster group %p.", c.ID, group)
	resp := response.NewClusterGroupResponse(group)
	result.SetError(request.RequestSuccessed, request.ErrRequestSuccessed, "cluster group")
	result.SetResponse(resp)
	return c.JSON(http.StatusOK, result)
}
