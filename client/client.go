// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddTagsBatchRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbfsList    *string `json:"DbfsList,omitempty" xml:"DbfsList,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddTagsBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsBatchRequest) GoString() string {
	return s.String()
}

func (s *AddTagsBatchRequest) SetClientToken(v string) *AddTagsBatchRequest {
	s.ClientToken = &v
	return s
}

func (s *AddTagsBatchRequest) SetDbfsList(v string) *AddTagsBatchRequest {
	s.DbfsList = &v
	return s
}

func (s *AddTagsBatchRequest) SetRegionId(v string) *AddTagsBatchRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsBatchRequest) SetTags(v string) *AddTagsBatchRequest {
	s.Tags = &v
	return s
}

type AddTagsBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagsBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsBatchResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsBatchResponseBody) SetRequestId(v string) *AddTagsBatchResponseBody {
	s.RequestId = &v
	return s
}

type AddTagsBatchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddTagsBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTagsBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsBatchResponse) GoString() string {
	return s.String()
}

func (s *AddTagsBatchResponse) SetHeaders(v map[string]*string) *AddTagsBatchResponse {
	s.Headers = v
	return s
}

func (s *AddTagsBatchResponse) SetStatusCode(v int32) *AddTagsBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagsBatchResponse) SetBody(v *AddTagsBatchResponseBody) *AddTagsBatchResponse {
	s.Body = v
	return s
}

type ApplyAutoSnapshotPolicyRequest struct {
	DbfsIds  []*string `json:"DbfsIds,omitempty" xml:"DbfsIds,omitempty" type:"Repeated"`
	PolicyId *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyRequest) SetDbfsIds(v []*string) *ApplyAutoSnapshotPolicyRequest {
	s.DbfsIds = v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetPolicyId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetRegionId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

type ApplyAutoSnapshotPolicyShrinkRequest struct {
	DbfsIdsShrink *string `json:"DbfsIds,omitempty" xml:"DbfsIds,omitempty"`
	PolicyId      *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyShrinkRequest) SetDbfsIdsShrink(v string) *ApplyAutoSnapshotPolicyShrinkRequest {
	s.DbfsIdsShrink = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyShrinkRequest) SetPolicyId(v string) *ApplyAutoSnapshotPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyShrinkRequest) SetRegionId(v string) *ApplyAutoSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

type ApplyAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponseBody) SetRequestId(v string) *ApplyAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ApplyAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *ApplyAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetStatusCode(v int32) *ApplyAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyResponse) SetBody(v *ApplyAutoSnapshotPolicyResponseBody) *ApplyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type AttachDbfsRequest struct {
	AttachMode    *string `json:"AttachMode,omitempty" xml:"AttachMode,omitempty"`
	AttachPoint   *string `json:"AttachPoint,omitempty" xml:"AttachPoint,omitempty"`
	ECSInstanceId *string `json:"ECSInstanceId,omitempty" xml:"ECSInstanceId,omitempty"`
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerUrl     *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
}

func (s AttachDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachDbfsRequest) GoString() string {
	return s.String()
}

func (s *AttachDbfsRequest) SetAttachMode(v string) *AttachDbfsRequest {
	s.AttachMode = &v
	return s
}

func (s *AttachDbfsRequest) SetAttachPoint(v string) *AttachDbfsRequest {
	s.AttachPoint = &v
	return s
}

func (s *AttachDbfsRequest) SetECSInstanceId(v string) *AttachDbfsRequest {
	s.ECSInstanceId = &v
	return s
}

func (s *AttachDbfsRequest) SetFsId(v string) *AttachDbfsRequest {
	s.FsId = &v
	return s
}

func (s *AttachDbfsRequest) SetRegionId(v string) *AttachDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDbfsRequest) SetServerUrl(v string) *AttachDbfsRequest {
	s.ServerUrl = &v
	return s
}

type AttachDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDbfsResponseBody) SetRequestId(v string) *AttachDbfsResponseBody {
	s.RequestId = &v
	return s
}

type AttachDbfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachDbfsResponse) GoString() string {
	return s.String()
}

func (s *AttachDbfsResponse) SetHeaders(v map[string]*string) *AttachDbfsResponse {
	s.Headers = v
	return s
}

func (s *AttachDbfsResponse) SetStatusCode(v int32) *AttachDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDbfsResponse) SetBody(v *AttachDbfsResponseBody) *AttachDbfsResponse {
	s.Body = v
	return s
}

type CancelAutoSnapshotPolicyRequest struct {
	DbfsIds  []*string `json:"DbfsIds,omitempty" xml:"DbfsIds,omitempty" type:"Repeated"`
	PolicyId *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyRequest) SetDbfsIds(v []*string) *CancelAutoSnapshotPolicyRequest {
	s.DbfsIds = v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetPolicyId(v string) *CancelAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetRegionId(v string) *CancelAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

type CancelAutoSnapshotPolicyShrinkRequest struct {
	DbfsIdsShrink *string `json:"DbfsIds,omitempty" xml:"DbfsIds,omitempty"`
	PolicyId      *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelAutoSnapshotPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyShrinkRequest) SetDbfsIdsShrink(v string) *CancelAutoSnapshotPolicyShrinkRequest {
	s.DbfsIdsShrink = &v
	return s
}

func (s *CancelAutoSnapshotPolicyShrinkRequest) SetPolicyId(v string) *CancelAutoSnapshotPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyShrinkRequest) SetRegionId(v string) *CancelAutoSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

type CancelAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CancelAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CancelAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CancelAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CancelAutoSnapshotPolicyResponse) SetStatusCode(v int32) *CancelAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAutoSnapshotPolicyResponse) SetBody(v *CancelAutoSnapshotPolicyResponseBody) *CancelAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CreateAutoSnapshotPolicyRequest struct {
	PolicyName     *string   `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdays []*string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty" type:"Repeated"`
	RetentionDays  *int32    `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	TimePoints     []*string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) SetPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRegionId(v string) *CreateAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRepeatWeekdays(v []*string) *CreateAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTimePoints(v []*string) *CreateAutoSnapshotPolicyRequest {
	s.TimePoints = v
	return s
}

type CreateAutoSnapshotPolicyShrinkRequest struct {
	PolicyName           *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdaysShrink *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	RetentionDays        *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	TimePointsShrink     *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s CreateAutoSnapshotPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyShrinkRequest) SetPolicyName(v string) *CreateAutoSnapshotPolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyShrinkRequest) SetRegionId(v string) *CreateAutoSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyShrinkRequest) SetRepeatWeekdaysShrink(v string) *CreateAutoSnapshotPolicyShrinkRequest {
	s.RepeatWeekdaysShrink = &v
	return s
}

func (s *CreateAutoSnapshotPolicyShrinkRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyShrinkRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyShrinkRequest) SetTimePointsShrink(v string) *CreateAutoSnapshotPolicyShrinkRequest {
	s.TimePointsShrink = &v
	return s
}

type CreateAutoSnapshotPolicyResponseBody struct {
	PolicyId  *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetPolicyId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponseBody) SetRequestId(v string) *CreateAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *CreateAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetStatusCode(v int32) *CreateAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoSnapshotPolicyResponse) SetBody(v *CreateAutoSnapshotPolicyResponseBody) *CreateAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type CreateDbfsRequest struct {
	AdvancedFeatures     *string `json:"AdvancedFeatures,omitempty" xml:"AdvancedFeatures,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeleteSnapshot       *bool   `json:"DeleteSnapshot,omitempty" xml:"DeleteSnapshot,omitempty"`
	EnableRaid           *bool   `json:"EnableRaid,omitempty" xml:"EnableRaid,omitempty"`
	Encryption           *bool   `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	FsName               *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RaidStripeUnitNumber *int32  `json:"RaidStripeUnitNumber,omitempty" xml:"RaidStripeUnitNumber,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG                *int32  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	SnapshotId           *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	UsedScene            *string `json:"UsedScene,omitempty" xml:"UsedScene,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDbfsRequest) GoString() string {
	return s.String()
}

func (s *CreateDbfsRequest) SetAdvancedFeatures(v string) *CreateDbfsRequest {
	s.AdvancedFeatures = &v
	return s
}

func (s *CreateDbfsRequest) SetCategory(v string) *CreateDbfsRequest {
	s.Category = &v
	return s
}

func (s *CreateDbfsRequest) SetClientToken(v string) *CreateDbfsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDbfsRequest) SetDeleteSnapshot(v bool) *CreateDbfsRequest {
	s.DeleteSnapshot = &v
	return s
}

func (s *CreateDbfsRequest) SetEnableRaid(v bool) *CreateDbfsRequest {
	s.EnableRaid = &v
	return s
}

func (s *CreateDbfsRequest) SetEncryption(v bool) *CreateDbfsRequest {
	s.Encryption = &v
	return s
}

func (s *CreateDbfsRequest) SetFsName(v string) *CreateDbfsRequest {
	s.FsName = &v
	return s
}

func (s *CreateDbfsRequest) SetInstanceType(v string) *CreateDbfsRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateDbfsRequest) SetKMSKeyId(v string) *CreateDbfsRequest {
	s.KMSKeyId = &v
	return s
}

func (s *CreateDbfsRequest) SetPerformanceLevel(v string) *CreateDbfsRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateDbfsRequest) SetRaidStripeUnitNumber(v int32) *CreateDbfsRequest {
	s.RaidStripeUnitNumber = &v
	return s
}

func (s *CreateDbfsRequest) SetRegionId(v string) *CreateDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDbfsRequest) SetSizeG(v int32) *CreateDbfsRequest {
	s.SizeG = &v
	return s
}

func (s *CreateDbfsRequest) SetSnapshotId(v string) *CreateDbfsRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateDbfsRequest) SetUsedScene(v string) *CreateDbfsRequest {
	s.UsedScene = &v
	return s
}

func (s *CreateDbfsRequest) SetZoneId(v string) *CreateDbfsRequest {
	s.ZoneId = &v
	return s
}

type CreateDbfsResponseBody struct {
	FsId      *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDbfsResponseBody) SetFsId(v string) *CreateDbfsResponseBody {
	s.FsId = &v
	return s
}

func (s *CreateDbfsResponseBody) SetRequestId(v string) *CreateDbfsResponseBody {
	s.RequestId = &v
	return s
}

type CreateDbfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDbfsResponse) GoString() string {
	return s.String()
}

func (s *CreateDbfsResponse) SetHeaders(v map[string]*string) *CreateDbfsResponse {
	s.Headers = v
	return s
}

func (s *CreateDbfsResponse) SetStatusCode(v int32) *CreateDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDbfsResponse) SetBody(v *CreateDbfsResponseBody) *CreateDbfsResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetClientToken(v string) *CreateServiceLinkedRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetentionDays *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SnapshotName  *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetFsId(v string) *CreateSnapshotRequest {
	s.FsId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

type CreateSnapshotResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetStatusCode(v int32) *CreateSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type DeleteAutoSnapshotPolicyRequest struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyRequest) SetPolicyId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) SetRegionId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

type DeleteAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyResponseBody) SetRequestId(v string) *DeleteAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *DeleteAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoSnapshotPolicyResponse) SetStatusCode(v int32) *DeleteAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyResponse) SetBody(v *DeleteAutoSnapshotPolicyResponseBody) *DeleteAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type DeleteDbfsRequest struct {
	Force    *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDbfsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDbfsRequest) SetForce(v bool) *DeleteDbfsRequest {
	s.Force = &v
	return s
}

func (s *DeleteDbfsRequest) SetFsId(v string) *DeleteDbfsRequest {
	s.FsId = &v
	return s
}

func (s *DeleteDbfsRequest) SetRegionId(v string) *DeleteDbfsRequest {
	s.RegionId = &v
	return s
}

type DeleteDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDbfsResponseBody) SetRequestId(v string) *DeleteDbfsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDbfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDbfsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDbfsResponse) SetHeaders(v map[string]*string) *DeleteDbfsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDbfsResponse) SetStatusCode(v int32) *DeleteDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDbfsResponse) SetBody(v *DeleteDbfsResponseBody) *DeleteDbfsResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	Force      *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetForce(v bool) *DeleteSnapshotRequest {
	s.Force = &v
	return s
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type DeleteSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetStatusCode(v int32) *DeleteSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DeleteTagsBatchRequest struct {
	DbfsList *string `json:"DbfsList,omitempty" xml:"DbfsList,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DeleteTagsBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsBatchRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagsBatchRequest) SetDbfsList(v string) *DeleteTagsBatchRequest {
	s.DbfsList = &v
	return s
}

func (s *DeleteTagsBatchRequest) SetRegionId(v string) *DeleteTagsBatchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTagsBatchRequest) SetTags(v string) *DeleteTagsBatchRequest {
	s.Tags = &v
	return s
}

type DeleteTagsBatchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagsBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagsBatchResponseBody) SetRequestId(v string) *DeleteTagsBatchResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagsBatchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTagsBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTagsBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagsBatchResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagsBatchResponse) SetHeaders(v map[string]*string) *DeleteTagsBatchResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagsBatchResponse) SetStatusCode(v int32) *DeleteTagsBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagsBatchResponse) SetBody(v *DeleteTagsBatchResponseBody) *DeleteTagsBatchResponse {
	s.Body = v
	return s
}

type DescribeDbfsSpecificationsRequest struct {
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDbfsSpecificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbfsSpecificationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbfsSpecificationsRequest) SetCategory(v string) *DescribeDbfsSpecificationsRequest {
	s.Category = &v
	return s
}

func (s *DescribeDbfsSpecificationsRequest) SetEcsInstanceType(v string) *DescribeDbfsSpecificationsRequest {
	s.EcsInstanceType = &v
	return s
}

func (s *DescribeDbfsSpecificationsRequest) SetRegionId(v string) *DescribeDbfsSpecificationsRequest {
	s.RegionId = &v
	return s
}

type DescribeDbfsSpecificationsResponseBody struct {
	MaxDbfsNumberPerEcs            map[string]interface{} `json:"MaxDbfsNumberPerEcs,omitempty" xml:"MaxDbfsNumberPerEcs,omitempty"`
	RequestId                      *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportedEcsInstanceTypeFamily []*string              `json:"SupportedEcsInstanceTypeFamily,omitempty" xml:"SupportedEcsInstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeDbfsSpecificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbfsSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbfsSpecificationsResponseBody) SetMaxDbfsNumberPerEcs(v map[string]interface{}) *DescribeDbfsSpecificationsResponseBody {
	s.MaxDbfsNumberPerEcs = v
	return s
}

func (s *DescribeDbfsSpecificationsResponseBody) SetRequestId(v string) *DescribeDbfsSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbfsSpecificationsResponseBody) SetSupportedEcsInstanceTypeFamily(v []*string) *DescribeDbfsSpecificationsResponseBody {
	s.SupportedEcsInstanceTypeFamily = v
	return s
}

type DescribeDbfsSpecificationsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDbfsSpecificationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDbfsSpecificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDbfsSpecificationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbfsSpecificationsResponse) SetHeaders(v map[string]*string) *DescribeDbfsSpecificationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbfsSpecificationsResponse) SetStatusCode(v int32) *DescribeDbfsSpecificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbfsSpecificationsResponse) SetBody(v *DescribeDbfsSpecificationsResponseBody) *DescribeDbfsSpecificationsResponse {
	s.Body = v
	return s
}

type DescribeInstanceTypesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) SetRegionId(v string) *DescribeInstanceTypesRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceTypesResponseBody struct {
	InstanceTypes []*DescribeInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBody) SetInstanceTypes(v []*DescribeInstanceTypesResponseBodyInstanceTypes) *DescribeInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetRequestId(v string) *DescribeInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceTypesResponseBodyInstanceTypes struct {
	CpuCoreCount            *float32 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	InstanceTypeDescription *string  `json:"InstanceTypeDescription,omitempty" xml:"InstanceTypeDescription,omitempty"`
	InstanceTypeId          *string  `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	MemorySize              *float32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetCpuCoreCount(v float32) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetInstanceTypeDescription(v string) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.InstanceTypeDescription = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetInstanceTypeId(v string) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetMemorySize(v float32) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.MemorySize = &v
	return s
}

type DescribeInstanceTypesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypesResponse) SetStatusCode(v int32) *DescribeInstanceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTypesResponse) SetBody(v *DescribeInstanceTypesResponseBody) *DescribeInstanceTypesResponse {
	s.Body = v
	return s
}

type DetachDbfsRequest struct {
	ECSInstanceId *string `json:"ECSInstanceId,omitempty" xml:"ECSInstanceId,omitempty"`
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachDbfsRequest) GoString() string {
	return s.String()
}

func (s *DetachDbfsRequest) SetECSInstanceId(v string) *DetachDbfsRequest {
	s.ECSInstanceId = &v
	return s
}

func (s *DetachDbfsRequest) SetFsId(v string) *DetachDbfsRequest {
	s.FsId = &v
	return s
}

func (s *DetachDbfsRequest) SetRegionId(v string) *DetachDbfsRequest {
	s.RegionId = &v
	return s
}

type DetachDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDbfsResponseBody) SetRequestId(v string) *DetachDbfsResponseBody {
	s.RequestId = &v
	return s
}

type DetachDbfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachDbfsResponse) GoString() string {
	return s.String()
}

func (s *DetachDbfsResponse) SetHeaders(v map[string]*string) *DetachDbfsResponse {
	s.Headers = v
	return s
}

func (s *DetachDbfsResponse) SetStatusCode(v int32) *DetachDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDbfsResponse) SetBody(v *DetachDbfsResponseBody) *DetachDbfsResponse {
	s.Body = v
	return s
}

type GetAutoSnapshotPolicyRequest struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetAutoSnapshotPolicyRequest) SetPolicyId(v string) *GetAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetAutoSnapshotPolicyRequest) SetRegionId(v string) *GetAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

type GetAutoSnapshotPolicyResponseBody struct {
	Data      *GetAutoSnapshotPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoSnapshotPolicyResponseBody) SetData(v *GetAutoSnapshotPolicyResponseBodyData) *GetAutoSnapshotPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBody) SetRequestId(v string) *GetAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetAutoSnapshotPolicyResponseBodyData struct {
	AccountId         *string   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AppliedDbfsNumber *int32    `json:"AppliedDbfsNumber,omitempty" xml:"AppliedDbfsNumber,omitempty"`
	CreatedTime       *string   `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	LastModified      *string   `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	PolicyId          *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName        *string   `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdays    []*string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty" type:"Repeated"`
	RetentionDays     *int32    `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	Status            *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDetail      *string   `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	TimePoints        []*string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s GetAutoSnapshotPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAutoSnapshotPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetAccountId(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetAppliedDbfsNumber(v int32) *GetAutoSnapshotPolicyResponseBodyData {
	s.AppliedDbfsNumber = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetCreatedTime(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetLastModified(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.LastModified = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetPolicyId(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetPolicyName(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.PolicyName = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetRegionId(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetRepeatWeekdays(v []*string) *GetAutoSnapshotPolicyResponseBodyData {
	s.RepeatWeekdays = v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetRetentionDays(v int32) *GetAutoSnapshotPolicyResponseBodyData {
	s.RetentionDays = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetStatus(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetStatusDetail(v string) *GetAutoSnapshotPolicyResponseBodyData {
	s.StatusDetail = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponseBodyData) SetTimePoints(v []*string) *GetAutoSnapshotPolicyResponseBodyData {
	s.TimePoints = v
	return s
}

type GetAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *GetAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetAutoSnapshotPolicyResponse) SetStatusCode(v int32) *GetAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoSnapshotPolicyResponse) SetBody(v *GetAutoSnapshotPolicyResponseBody) *GetAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type GetDbfsRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsRequest) GoString() string {
	return s.String()
}

func (s *GetDbfsRequest) SetFsId(v string) *GetDbfsRequest {
	s.FsId = &v
	return s
}

func (s *GetDbfsRequest) SetRegionId(v string) *GetDbfsRequest {
	s.RegionId = &v
	return s
}

type GetDbfsResponseBody struct {
	DBFSInfo  *GetDbfsResponseBodyDBFSInfo `json:"DBFSInfo,omitempty" xml:"DBFSInfo,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBody) SetDBFSInfo(v *GetDbfsResponseBodyDBFSInfo) *GetDbfsResponseBody {
	s.DBFSInfo = v
	return s
}

func (s *GetDbfsResponseBody) SetRequestId(v string) *GetDbfsResponseBody {
	s.RequestId = &v
	return s
}

type GetDbfsResponseBodyDBFSInfo struct {
	AttachNodeNumber *int32                                   `json:"AttachNodeNumber,omitempty" xml:"AttachNodeNumber,omitempty"`
	Category         *string                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	CreatedTime      *string                                  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DBFSClusterId    *string                                  `json:"DBFSClusterId,omitempty" xml:"DBFSClusterId,omitempty"`
	Description      *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EbsList          []*GetDbfsResponseBodyDBFSInfoEbsList    `json:"EbsList,omitempty" xml:"EbsList,omitempty" type:"Repeated"`
	EcsList          []*GetDbfsResponseBodyDBFSInfoEcsList    `json:"EcsList,omitempty" xml:"EcsList,omitempty" type:"Repeated"`
	EnableRaid       *bool                                    `json:"EnableRaid,omitempty" xml:"EnableRaid,omitempty"`
	Encryption       *bool                                    `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	FsId             *string                                  `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName           *string                                  `json:"FsName,omitempty" xml:"FsName,omitempty"`
	InstanceType     *string                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KMSKeyId         *string                                  `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	LastFailed       *string                                  `json:"LastFailed,omitempty" xml:"LastFailed,omitempty"`
	LastMountTime    *string                                  `json:"LastMountTime,omitempty" xml:"LastMountTime,omitempty"`
	LastUmountTime   *string                                  `json:"LastUmountTime,omitempty" xml:"LastUmountTime,omitempty"`
	PayType          *string                                  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PerformanceLevel *string                                  `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RaidStrip        *int32                                   `json:"RaidStrip,omitempty" xml:"RaidStrip,omitempty"`
	RegionId         *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG            *int32                                   `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	SnapshotInfo     *GetDbfsResponseBodyDBFSInfoSnapshotInfo `json:"SnapshotInfo,omitempty" xml:"SnapshotInfo,omitempty" type:"Struct"`
	Status           *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags             []*GetDbfsResponseBodyDBFSInfoTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UsedScene        *string                                  `json:"UsedScene,omitempty" xml:"UsedScene,omitempty"`
	ZoneId           *string                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfo) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfo) SetAttachNodeNumber(v int32) *GetDbfsResponseBodyDBFSInfo {
	s.AttachNodeNumber = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetCategory(v string) *GetDbfsResponseBodyDBFSInfo {
	s.Category = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetCreatedTime(v string) *GetDbfsResponseBodyDBFSInfo {
	s.CreatedTime = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetDBFSClusterId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.DBFSClusterId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetDescription(v string) *GetDbfsResponseBodyDBFSInfo {
	s.Description = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEbsList(v []*GetDbfsResponseBodyDBFSInfoEbsList) *GetDbfsResponseBodyDBFSInfo {
	s.EbsList = v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEcsList(v []*GetDbfsResponseBodyDBFSInfoEcsList) *GetDbfsResponseBodyDBFSInfo {
	s.EcsList = v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEnableRaid(v bool) *GetDbfsResponseBodyDBFSInfo {
	s.EnableRaid = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetEncryption(v bool) *GetDbfsResponseBodyDBFSInfo {
	s.Encryption = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetFsId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.FsId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetFsName(v string) *GetDbfsResponseBodyDBFSInfo {
	s.FsName = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetInstanceType(v string) *GetDbfsResponseBodyDBFSInfo {
	s.InstanceType = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetKMSKeyId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.KMSKeyId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetLastFailed(v string) *GetDbfsResponseBodyDBFSInfo {
	s.LastFailed = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetLastMountTime(v string) *GetDbfsResponseBodyDBFSInfo {
	s.LastMountTime = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetLastUmountTime(v string) *GetDbfsResponseBodyDBFSInfo {
	s.LastUmountTime = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetPayType(v string) *GetDbfsResponseBodyDBFSInfo {
	s.PayType = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetPerformanceLevel(v string) *GetDbfsResponseBodyDBFSInfo {
	s.PerformanceLevel = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetRaidStrip(v int32) *GetDbfsResponseBodyDBFSInfo {
	s.RaidStrip = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetRegionId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.RegionId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetSizeG(v int32) *GetDbfsResponseBodyDBFSInfo {
	s.SizeG = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetSnapshotInfo(v *GetDbfsResponseBodyDBFSInfoSnapshotInfo) *GetDbfsResponseBodyDBFSInfo {
	s.SnapshotInfo = v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetStatus(v string) *GetDbfsResponseBodyDBFSInfo {
	s.Status = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetTags(v []*GetDbfsResponseBodyDBFSInfoTags) *GetDbfsResponseBodyDBFSInfo {
	s.Tags = v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetUsedScene(v string) *GetDbfsResponseBodyDBFSInfo {
	s.UsedScene = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfo) SetZoneId(v string) *GetDbfsResponseBodyDBFSInfo {
	s.ZoneId = &v
	return s
}

type GetDbfsResponseBodyDBFSInfoEbsList struct {
	EbsId *string `json:"EbsId,omitempty" xml:"EbsId,omitempty"`
	SizeG *int32  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfoEbsList) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfoEbsList) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfoEbsList) SetEbsId(v string) *GetDbfsResponseBodyDBFSInfoEbsList {
	s.EbsId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoEbsList) SetSizeG(v int32) *GetDbfsResponseBodyDBFSInfoEbsList {
	s.SizeG = &v
	return s
}

type GetDbfsResponseBodyDBFSInfoEcsList struct {
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfoEcsList) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfoEcsList) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfoEcsList) SetEcsId(v string) *GetDbfsResponseBodyDBFSInfoEcsList {
	s.EcsId = &v
	return s
}

type GetDbfsResponseBodyDBFSInfoSnapshotInfo struct {
	LinkId        *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	PolicyId      *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SnapshotCount *int32  `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	TotalSize     *int64  `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfoSnapshotInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfoSnapshotInfo) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfoSnapshotInfo) SetLinkId(v string) *GetDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.LinkId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoSnapshotInfo) SetPolicyId(v string) *GetDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.PolicyId = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoSnapshotInfo) SetSnapshotCount(v int32) *GetDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.SnapshotCount = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoSnapshotInfo) SetTotalSize(v int64) *GetDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.TotalSize = &v
	return s
}

type GetDbfsResponseBodyDBFSInfoTags struct {
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetDbfsResponseBodyDBFSInfoTags) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponseBodyDBFSInfoTags) GoString() string {
	return s.String()
}

func (s *GetDbfsResponseBodyDBFSInfoTags) SetId(v int32) *GetDbfsResponseBodyDBFSInfoTags {
	s.Id = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoTags) SetTagKey(v string) *GetDbfsResponseBodyDBFSInfoTags {
	s.TagKey = &v
	return s
}

func (s *GetDbfsResponseBodyDBFSInfoTags) SetTagValue(v string) *GetDbfsResponseBodyDBFSInfoTags {
	s.TagValue = &v
	return s
}

type GetDbfsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDbfsResponse) GoString() string {
	return s.String()
}

func (s *GetDbfsResponse) SetHeaders(v map[string]*string) *GetDbfsResponse {
	s.Headers = v
	return s
}

func (s *GetDbfsResponse) SetStatusCode(v int32) *GetDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDbfsResponse) SetBody(v *GetDbfsResponseBody) *GetDbfsResponse {
	s.Body = v
	return s
}

type GetServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleRequest) SetRegionId(v string) *GetServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type GetServiceLinkedRoleResponseBody struct {
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DbfsLinkedRole *bool   `json:"DbfsLinkedRole,omitempty" xml:"DbfsLinkedRole,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleResponseBody) SetAccountId(v string) *GetServiceLinkedRoleResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetServiceLinkedRoleResponseBody) SetDbfsLinkedRole(v bool) *GetServiceLinkedRoleResponseBody {
	s.DbfsLinkedRole = &v
	return s
}

func (s *GetServiceLinkedRoleResponseBody) SetRegionId(v string) *GetServiceLinkedRoleResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetServiceLinkedRoleResponseBody) SetRequestId(v string) *GetServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceLinkedRoleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *GetServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLinkedRoleResponse) SetStatusCode(v int32) *GetServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLinkedRoleResponse) SetBody(v *GetServiceLinkedRoleResponseBody) *GetServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type GetSnapshotLinkRequest struct {
	LinkId   *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSnapshotLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotLinkRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotLinkRequest) SetLinkId(v string) *GetSnapshotLinkRequest {
	s.LinkId = &v
	return s
}

func (s *GetSnapshotLinkRequest) SetRegionId(v string) *GetSnapshotLinkRequest {
	s.RegionId = &v
	return s
}

type GetSnapshotLinkResponseBody struct {
	Data      *GetSnapshotLinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSnapshotLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotLinkResponseBody) SetData(v *GetSnapshotLinkResponseBodyData) *GetSnapshotLinkResponseBody {
	s.Data = v
	return s
}

func (s *GetSnapshotLinkResponseBody) SetRequestId(v string) *GetSnapshotLinkResponseBody {
	s.RequestId = &v
	return s
}

type GetSnapshotLinkResponseBodyData struct {
	Category      *string                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	EcsList       []*GetSnapshotLinkResponseBodyDataEcsList `json:"EcsList,omitempty" xml:"EcsList,omitempty" type:"Repeated"`
	FsId          *string                                   `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName        *string                                   `json:"FsName,omitempty" xml:"FsName,omitempty"`
	LinkId        *string                                   `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	SnapshotCount *int32                                    `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	SourceSize    *int32                                    `json:"SourceSize,omitempty" xml:"SourceSize,omitempty"`
	Status        *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalSize     *int64                                    `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetSnapshotLinkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSnapshotLinkResponseBodyData) SetCategory(v string) *GetSnapshotLinkResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetEcsList(v []*GetSnapshotLinkResponseBodyDataEcsList) *GetSnapshotLinkResponseBodyData {
	s.EcsList = v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetFsId(v string) *GetSnapshotLinkResponseBodyData {
	s.FsId = &v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetFsName(v string) *GetSnapshotLinkResponseBodyData {
	s.FsName = &v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetLinkId(v string) *GetSnapshotLinkResponseBodyData {
	s.LinkId = &v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetSnapshotCount(v int32) *GetSnapshotLinkResponseBodyData {
	s.SnapshotCount = &v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetSourceSize(v int32) *GetSnapshotLinkResponseBodyData {
	s.SourceSize = &v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetStatus(v string) *GetSnapshotLinkResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSnapshotLinkResponseBodyData) SetTotalSize(v int64) *GetSnapshotLinkResponseBodyData {
	s.TotalSize = &v
	return s
}

type GetSnapshotLinkResponseBodyDataEcsList struct {
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
}

func (s GetSnapshotLinkResponseBodyDataEcsList) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotLinkResponseBodyDataEcsList) GoString() string {
	return s.String()
}

func (s *GetSnapshotLinkResponseBodyDataEcsList) SetEcsId(v string) *GetSnapshotLinkResponseBodyDataEcsList {
	s.EcsId = &v
	return s
}

type GetSnapshotLinkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSnapshotLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSnapshotLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotLinkResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotLinkResponse) SetHeaders(v map[string]*string) *GetSnapshotLinkResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotLinkResponse) SetStatusCode(v int32) *GetSnapshotLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSnapshotLinkResponse) SetBody(v *GetSnapshotLinkResponseBody) *GetSnapshotLinkResponse {
	s.Body = v
	return s
}

type ListAutoSnapshotPoliciesRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAutoSnapshotPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPoliciesRequest) SetFilterKey(v string) *ListAutoSnapshotPoliciesRequest {
	s.FilterKey = &v
	return s
}

func (s *ListAutoSnapshotPoliciesRequest) SetFilterValue(v string) *ListAutoSnapshotPoliciesRequest {
	s.FilterValue = &v
	return s
}

func (s *ListAutoSnapshotPoliciesRequest) SetPageNumber(v int32) *ListAutoSnapshotPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAutoSnapshotPoliciesRequest) SetPageSize(v int32) *ListAutoSnapshotPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoSnapshotPoliciesRequest) SetRegionId(v string) *ListAutoSnapshotPoliciesRequest {
	s.RegionId = &v
	return s
}

type ListAutoSnapshotPoliciesResponseBody struct {
	PageNumber       *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotPolicies []*ListAutoSnapshotPoliciesResponseBodySnapshotPolicies `json:"SnapshotPolicies,omitempty" xml:"SnapshotPolicies,omitempty" type:"Repeated"`
	TotalCount       *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutoSnapshotPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPoliciesResponseBody) SetPageNumber(v int32) *ListAutoSnapshotPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBody) SetPageSize(v int32) *ListAutoSnapshotPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBody) SetRequestId(v string) *ListAutoSnapshotPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBody) SetSnapshotPolicies(v []*ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) *ListAutoSnapshotPoliciesResponseBody {
	s.SnapshotPolicies = v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBody) SetTotalCount(v int32) *ListAutoSnapshotPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAutoSnapshotPoliciesResponseBodySnapshotPolicies struct {
	AccountId         *string   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AppliedDbfsNumber *int32    `json:"AppliedDbfsNumber,omitempty" xml:"AppliedDbfsNumber,omitempty"`
	CreatedTime       *string   `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	LastModified      *string   `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	PolicyId          *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName        *string   `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdays    []*string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty" type:"Repeated"`
	RetentionDays     *int32    `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	Status            *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDetail      *string   `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	TimePoints        []*string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetAccountId(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.AccountId = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetAppliedDbfsNumber(v int32) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.AppliedDbfsNumber = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetCreatedTime(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.CreatedTime = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetLastModified(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.LastModified = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetPolicyId(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetPolicyName(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetRegionId(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.RegionId = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetRepeatWeekdays(v []*string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.RepeatWeekdays = v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetRetentionDays(v int32) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.RetentionDays = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetStatus(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.Status = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetStatusDetail(v string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.StatusDetail = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies) SetTimePoints(v []*string) *ListAutoSnapshotPoliciesResponseBodySnapshotPolicies {
	s.TimePoints = v
	return s
}

type ListAutoSnapshotPoliciesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAutoSnapshotPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAutoSnapshotPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPoliciesResponse) SetHeaders(v map[string]*string) *ListAutoSnapshotPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoSnapshotPoliciesResponse) SetStatusCode(v int32) *ListAutoSnapshotPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoSnapshotPoliciesResponse) SetBody(v *ListAutoSnapshotPoliciesResponseBody) *ListAutoSnapshotPoliciesResponse {
	s.Body = v
	return s
}

type ListAutoSnapshotPolicyAppliedDbfsRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyId    *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAutoSnapshotPolicyAppliedDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyAppliedDbfsRequest) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyAppliedDbfsRequest) SetFilterKey(v string) *ListAutoSnapshotPolicyAppliedDbfsRequest {
	s.FilterKey = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsRequest) SetFilterValue(v string) *ListAutoSnapshotPolicyAppliedDbfsRequest {
	s.FilterValue = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsRequest) SetPageNumber(v int32) *ListAutoSnapshotPolicyAppliedDbfsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsRequest) SetPageSize(v int32) *ListAutoSnapshotPolicyAppliedDbfsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsRequest) SetPolicyId(v string) *ListAutoSnapshotPolicyAppliedDbfsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsRequest) SetRegionId(v string) *ListAutoSnapshotPolicyAppliedDbfsRequest {
	s.RegionId = &v
	return s
}

type ListAutoSnapshotPolicyAppliedDbfsResponseBody struct {
	DbfsList   []*ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList `json:"DbfsList,omitempty" xml:"DbfsList,omitempty" type:"Repeated"`
	PageNumber *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutoSnapshotPolicyAppliedDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyAppliedDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBody) SetDbfsList(v []*ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) *ListAutoSnapshotPolicyAppliedDbfsResponseBody {
	s.DbfsList = v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBody) SetPageNumber(v int32) *ListAutoSnapshotPolicyAppliedDbfsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBody) SetPageSize(v int32) *ListAutoSnapshotPolicyAppliedDbfsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBody) SetRequestId(v string) *ListAutoSnapshotPolicyAppliedDbfsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBody) SetTotalCount(v int32) *ListAutoSnapshotPolicyAppliedDbfsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList struct {
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName        *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG         *int64  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	SnapshotCount *int32  `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalSize     *int64  `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) SetFsId(v string) *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList {
	s.FsId = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) SetFsName(v string) *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList {
	s.FsName = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) SetRegionId(v string) *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList {
	s.RegionId = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) SetSizeG(v int64) *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList {
	s.SizeG = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) SetSnapshotCount(v int32) *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList {
	s.SnapshotCount = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) SetStatus(v string) *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList {
	s.Status = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList) SetTotalSize(v int64) *ListAutoSnapshotPolicyAppliedDbfsResponseBodyDbfsList {
	s.TotalSize = &v
	return s
}

type ListAutoSnapshotPolicyAppliedDbfsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAutoSnapshotPolicyAppliedDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAutoSnapshotPolicyAppliedDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyAppliedDbfsResponse) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponse) SetHeaders(v map[string]*string) *ListAutoSnapshotPolicyAppliedDbfsResponse {
	s.Headers = v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponse) SetStatusCode(v int32) *ListAutoSnapshotPolicyAppliedDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoSnapshotPolicyAppliedDbfsResponse) SetBody(v *ListAutoSnapshotPolicyAppliedDbfsResponseBody) *ListAutoSnapshotPolicyAppliedDbfsResponse {
	s.Body = v
	return s
}

type ListAutoSnapshotPolicyUnappliedDbfsRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAutoSnapshotPolicyUnappliedDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyUnappliedDbfsRequest) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsRequest) SetFilterKey(v string) *ListAutoSnapshotPolicyUnappliedDbfsRequest {
	s.FilterKey = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsRequest) SetFilterValue(v string) *ListAutoSnapshotPolicyUnappliedDbfsRequest {
	s.FilterValue = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsRequest) SetPageNumber(v int32) *ListAutoSnapshotPolicyUnappliedDbfsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsRequest) SetPageSize(v int32) *ListAutoSnapshotPolicyUnappliedDbfsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsRequest) SetRegionId(v string) *ListAutoSnapshotPolicyUnappliedDbfsRequest {
	s.RegionId = &v
	return s
}

type ListAutoSnapshotPolicyUnappliedDbfsResponseBody struct {
	DbfsList   []*ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList `json:"DbfsList,omitempty" xml:"DbfsList,omitempty" type:"Repeated"`
	PageNumber *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutoSnapshotPolicyUnappliedDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyUnappliedDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBody) SetDbfsList(v []*ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) *ListAutoSnapshotPolicyUnappliedDbfsResponseBody {
	s.DbfsList = v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBody) SetPageNumber(v int32) *ListAutoSnapshotPolicyUnappliedDbfsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBody) SetPageSize(v int32) *ListAutoSnapshotPolicyUnappliedDbfsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBody) SetRequestId(v string) *ListAutoSnapshotPolicyUnappliedDbfsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBody) SetTotalCount(v int32) *ListAutoSnapshotPolicyUnappliedDbfsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList struct {
	FsId          *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName        *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG         *int64  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	SnapshotCount *int32  `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalSize     *int64  `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) SetFsId(v string) *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList {
	s.FsId = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) SetFsName(v string) *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList {
	s.FsName = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) SetRegionId(v string) *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList {
	s.RegionId = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) SetSizeG(v int64) *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList {
	s.SizeG = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) SetSnapshotCount(v int32) *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList {
	s.SnapshotCount = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) SetStatus(v string) *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList {
	s.Status = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList) SetTotalSize(v int64) *ListAutoSnapshotPolicyUnappliedDbfsResponseBodyDbfsList {
	s.TotalSize = &v
	return s
}

type ListAutoSnapshotPolicyUnappliedDbfsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAutoSnapshotPolicyUnappliedDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAutoSnapshotPolicyUnappliedDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutoSnapshotPolicyUnappliedDbfsResponse) GoString() string {
	return s.String()
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponse) SetHeaders(v map[string]*string) *ListAutoSnapshotPolicyUnappliedDbfsResponse {
	s.Headers = v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponse) SetStatusCode(v int32) *ListAutoSnapshotPolicyUnappliedDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoSnapshotPolicyUnappliedDbfsResponse) SetBody(v *ListAutoSnapshotPolicyUnappliedDbfsResponseBody) *ListAutoSnapshotPolicyUnappliedDbfsResponse {
	s.Body = v
	return s
}

type ListDbfsRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortKey     *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortType    *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsRequest) GoString() string {
	return s.String()
}

func (s *ListDbfsRequest) SetFilterKey(v string) *ListDbfsRequest {
	s.FilterKey = &v
	return s
}

func (s *ListDbfsRequest) SetFilterValue(v string) *ListDbfsRequest {
	s.FilterValue = &v
	return s
}

func (s *ListDbfsRequest) SetPageNumber(v int32) *ListDbfsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDbfsRequest) SetPageSize(v int32) *ListDbfsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDbfsRequest) SetRegionId(v string) *ListDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDbfsRequest) SetSortKey(v string) *ListDbfsRequest {
	s.SortKey = &v
	return s
}

func (s *ListDbfsRequest) SetSortType(v string) *ListDbfsRequest {
	s.SortType = &v
	return s
}

func (s *ListDbfsRequest) SetTags(v string) *ListDbfsRequest {
	s.Tags = &v
	return s
}

type ListDbfsResponseBody struct {
	DBFSInfo   []*ListDbfsResponseBodyDBFSInfo `json:"DBFSInfo,omitempty" xml:"DBFSInfo,omitempty" type:"Repeated"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBody) SetDBFSInfo(v []*ListDbfsResponseBodyDBFSInfo) *ListDbfsResponseBody {
	s.DBFSInfo = v
	return s
}

func (s *ListDbfsResponseBody) SetPageNumber(v int32) *ListDbfsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDbfsResponseBody) SetPageSize(v int32) *ListDbfsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDbfsResponseBody) SetRequestId(v string) *ListDbfsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDbfsResponseBody) SetTotalCount(v int32) *ListDbfsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDbfsResponseBodyDBFSInfo struct {
	AttachNodeNumber *int32                                    `json:"AttachNodeNumber,omitempty" xml:"AttachNodeNumber,omitempty"`
	Category         *string                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	CreatedTime      *string                                   `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DBFSClusterId    *string                                   `json:"DBFSClusterId,omitempty" xml:"DBFSClusterId,omitempty"`
	EbsList          []*ListDbfsResponseBodyDBFSInfoEbsList    `json:"EbsList,omitempty" xml:"EbsList,omitempty" type:"Repeated"`
	EcsList          []*ListDbfsResponseBodyDBFSInfoEcsList    `json:"EcsList,omitempty" xml:"EcsList,omitempty" type:"Repeated"`
	EnableRaid       *bool                                     `json:"EnableRaid,omitempty" xml:"EnableRaid,omitempty"`
	Encryption       *bool                                     `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	FsId             *string                                   `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName           *string                                   `json:"FsName,omitempty" xml:"FsName,omitempty"`
	InstanceType     *string                                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KMSKeyId         *string                                   `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	LastFailed       *string                                   `json:"LastFailed,omitempty" xml:"LastFailed,omitempty"`
	LastMountTime    *string                                   `json:"LastMountTime,omitempty" xml:"LastMountTime,omitempty"`
	LastUmountTime   *string                                   `json:"LastUmountTime,omitempty" xml:"LastUmountTime,omitempty"`
	PayType          *string                                   `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PerformanceLevel *string                                   `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RaidStrip        *int32                                    `json:"RaidStrip,omitempty" xml:"RaidStrip,omitempty"`
	RegionId         *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SizeG            *int32                                    `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
	SnapshotInfo     *ListDbfsResponseBodyDBFSInfoSnapshotInfo `json:"SnapshotInfo,omitempty" xml:"SnapshotInfo,omitempty" type:"Struct"`
	Status           *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags             []*ListDbfsResponseBodyDBFSInfoTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UsedScene        *string                                   `json:"UsedScene,omitempty" xml:"UsedScene,omitempty"`
	ZoneId           *string                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfo) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfo) SetAttachNodeNumber(v int32) *ListDbfsResponseBodyDBFSInfo {
	s.AttachNodeNumber = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetCategory(v string) *ListDbfsResponseBodyDBFSInfo {
	s.Category = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetCreatedTime(v string) *ListDbfsResponseBodyDBFSInfo {
	s.CreatedTime = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetDBFSClusterId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.DBFSClusterId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEbsList(v []*ListDbfsResponseBodyDBFSInfoEbsList) *ListDbfsResponseBodyDBFSInfo {
	s.EbsList = v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEcsList(v []*ListDbfsResponseBodyDBFSInfoEcsList) *ListDbfsResponseBodyDBFSInfo {
	s.EcsList = v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEnableRaid(v bool) *ListDbfsResponseBodyDBFSInfo {
	s.EnableRaid = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetEncryption(v bool) *ListDbfsResponseBodyDBFSInfo {
	s.Encryption = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetFsId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.FsId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetFsName(v string) *ListDbfsResponseBodyDBFSInfo {
	s.FsName = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetInstanceType(v string) *ListDbfsResponseBodyDBFSInfo {
	s.InstanceType = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetKMSKeyId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.KMSKeyId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetLastFailed(v string) *ListDbfsResponseBodyDBFSInfo {
	s.LastFailed = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetLastMountTime(v string) *ListDbfsResponseBodyDBFSInfo {
	s.LastMountTime = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetLastUmountTime(v string) *ListDbfsResponseBodyDBFSInfo {
	s.LastUmountTime = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetPayType(v string) *ListDbfsResponseBodyDBFSInfo {
	s.PayType = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetPerformanceLevel(v string) *ListDbfsResponseBodyDBFSInfo {
	s.PerformanceLevel = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetRaidStrip(v int32) *ListDbfsResponseBodyDBFSInfo {
	s.RaidStrip = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetRegionId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.RegionId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetSizeG(v int32) *ListDbfsResponseBodyDBFSInfo {
	s.SizeG = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetSnapshotInfo(v *ListDbfsResponseBodyDBFSInfoSnapshotInfo) *ListDbfsResponseBodyDBFSInfo {
	s.SnapshotInfo = v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetStatus(v string) *ListDbfsResponseBodyDBFSInfo {
	s.Status = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetTags(v []*ListDbfsResponseBodyDBFSInfoTags) *ListDbfsResponseBodyDBFSInfo {
	s.Tags = v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetUsedScene(v string) *ListDbfsResponseBodyDBFSInfo {
	s.UsedScene = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfo) SetZoneId(v string) *ListDbfsResponseBodyDBFSInfo {
	s.ZoneId = &v
	return s
}

type ListDbfsResponseBodyDBFSInfoEbsList struct {
	EbsId *string `json:"EbsId,omitempty" xml:"EbsId,omitempty"`
	SizeG *int32  `json:"SizeG,omitempty" xml:"SizeG,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfoEbsList) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfoEbsList) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfoEbsList) SetEbsId(v string) *ListDbfsResponseBodyDBFSInfoEbsList {
	s.EbsId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoEbsList) SetSizeG(v int32) *ListDbfsResponseBodyDBFSInfoEbsList {
	s.SizeG = &v
	return s
}

type ListDbfsResponseBodyDBFSInfoEcsList struct {
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfoEcsList) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfoEcsList) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfoEcsList) SetEcsId(v string) *ListDbfsResponseBodyDBFSInfoEcsList {
	s.EcsId = &v
	return s
}

type ListDbfsResponseBodyDBFSInfoSnapshotInfo struct {
	LinkId        *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	PolicyId      *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SnapshotCount *int32  `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	TotalSize     *int64  `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfoSnapshotInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfoSnapshotInfo) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfoSnapshotInfo) SetLinkId(v string) *ListDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.LinkId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoSnapshotInfo) SetPolicyId(v string) *ListDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.PolicyId = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoSnapshotInfo) SetSnapshotCount(v int32) *ListDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.SnapshotCount = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoSnapshotInfo) SetTotalSize(v int64) *ListDbfsResponseBodyDBFSInfoSnapshotInfo {
	s.TotalSize = &v
	return s
}

type ListDbfsResponseBodyDBFSInfoTags struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListDbfsResponseBodyDBFSInfoTags) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponseBodyDBFSInfoTags) GoString() string {
	return s.String()
}

func (s *ListDbfsResponseBodyDBFSInfoTags) SetId(v int64) *ListDbfsResponseBodyDBFSInfoTags {
	s.Id = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoTags) SetTagKey(v string) *ListDbfsResponseBodyDBFSInfoTags {
	s.TagKey = &v
	return s
}

func (s *ListDbfsResponseBodyDBFSInfoTags) SetTagValue(v string) *ListDbfsResponseBodyDBFSInfoTags {
	s.TagValue = &v
	return s
}

type ListDbfsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsResponse) GoString() string {
	return s.String()
}

func (s *ListDbfsResponse) SetHeaders(v map[string]*string) *ListDbfsResponse {
	s.Headers = v
	return s
}

func (s *ListDbfsResponse) SetStatusCode(v int32) *ListDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDbfsResponse) SetBody(v *ListDbfsResponseBody) *ListDbfsResponse {
	s.Body = v
	return s
}

type ListDbfsAttachableEcsInstancesRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDbfsAttachableEcsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachableEcsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachableEcsInstancesRequest) SetFilterKey(v string) *ListDbfsAttachableEcsInstancesRequest {
	s.FilterKey = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesRequest) SetFilterValue(v string) *ListDbfsAttachableEcsInstancesRequest {
	s.FilterValue = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesRequest) SetPageNumber(v int32) *ListDbfsAttachableEcsInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesRequest) SetPageSize(v int32) *ListDbfsAttachableEcsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesRequest) SetRegionId(v string) *ListDbfsAttachableEcsInstancesRequest {
	s.RegionId = &v
	return s
}

type ListDbfsAttachableEcsInstancesResponseBody struct {
	EcsLabelInfo []*ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo `json:"EcsLabelInfo,omitempty" xml:"EcsLabelInfo,omitempty" type:"Repeated"`
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDbfsAttachableEcsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachableEcsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachableEcsInstancesResponseBody) SetEcsLabelInfo(v []*ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) *ListDbfsAttachableEcsInstancesResponseBody {
	s.EcsLabelInfo = v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBody) SetRequestId(v string) *ListDbfsAttachableEcsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBody) SetTotalCount(v int32) *ListDbfsAttachableEcsInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo struct {
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	OSName             *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Label              *string `json:"label,omitempty" xml:"label,omitempty"`
	Value              *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) SetImageId(v string) *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo {
	s.ImageId = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) SetInstanceTypeFamily(v string) *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo {
	s.InstanceTypeFamily = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) SetOSName(v string) *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo {
	s.OSName = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) SetStatus(v string) *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo {
	s.Status = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) SetZoneId(v string) *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo {
	s.ZoneId = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) SetLabel(v string) *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo {
	s.Label = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo) SetValue(v string) *ListDbfsAttachableEcsInstancesResponseBodyEcsLabelInfo {
	s.Value = &v
	return s
}

type ListDbfsAttachableEcsInstancesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDbfsAttachableEcsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDbfsAttachableEcsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachableEcsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachableEcsInstancesResponse) SetHeaders(v map[string]*string) *ListDbfsAttachableEcsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponse) SetStatusCode(v int32) *ListDbfsAttachableEcsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDbfsAttachableEcsInstancesResponse) SetBody(v *ListDbfsAttachableEcsInstancesResponseBody) *ListDbfsAttachableEcsInstancesResponse {
	s.Body = v
	return s
}

type ListDbfsAttachedEcsInstancesRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDbfsAttachedEcsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachedEcsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachedEcsInstancesRequest) SetFsId(v string) *ListDbfsAttachedEcsInstancesRequest {
	s.FsId = &v
	return s
}

func (s *ListDbfsAttachedEcsInstancesRequest) SetRegionId(v string) *ListDbfsAttachedEcsInstancesRequest {
	s.RegionId = &v
	return s
}

type ListDbfsAttachedEcsInstancesResponseBody struct {
	EcsLabelInfo []*ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo `json:"EcsLabelInfo,omitempty" xml:"EcsLabelInfo,omitempty" type:"Repeated"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDbfsAttachedEcsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachedEcsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachedEcsInstancesResponseBody) SetEcsLabelInfo(v []*ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) *ListDbfsAttachedEcsInstancesResponseBody {
	s.EcsLabelInfo = v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponseBody) SetRequestId(v string) *ListDbfsAttachedEcsInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo struct {
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	MountPoint         *string `json:"MountPoint,omitempty" xml:"MountPoint,omitempty"`
	MountedTime        *string `json:"MountedTime,omitempty" xml:"MountedTime,omitempty"`
	OSName             *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	Label              *string `json:"label,omitempty" xml:"label,omitempty"`
	Value              *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) SetInstanceTypeFamily(v string) *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo {
	s.InstanceTypeFamily = &v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) SetMountPoint(v string) *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo {
	s.MountPoint = &v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) SetMountedTime(v string) *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo {
	s.MountedTime = &v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) SetOSName(v string) *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo {
	s.OSName = &v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) SetLabel(v string) *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo {
	s.Label = &v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo) SetValue(v string) *ListDbfsAttachedEcsInstancesResponseBodyEcsLabelInfo {
	s.Value = &v
	return s
}

type ListDbfsAttachedEcsInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDbfsAttachedEcsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDbfsAttachedEcsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDbfsAttachedEcsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDbfsAttachedEcsInstancesResponse) SetHeaders(v map[string]*string) *ListDbfsAttachedEcsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponse) SetStatusCode(v int32) *ListDbfsAttachedEcsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDbfsAttachedEcsInstancesResponse) SetBody(v *ListDbfsAttachedEcsInstancesResponseBody) *ListDbfsAttachedEcsInstancesResponse {
	s.Body = v
	return s
}

type ListSnapshotRequest struct {
	FilterKey    *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue  *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	FsId         *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotIds  *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SortKey      *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	SortType     *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotRequest) SetFilterKey(v string) *ListSnapshotRequest {
	s.FilterKey = &v
	return s
}

func (s *ListSnapshotRequest) SetFilterValue(v string) *ListSnapshotRequest {
	s.FilterValue = &v
	return s
}

func (s *ListSnapshotRequest) SetFsId(v string) *ListSnapshotRequest {
	s.FsId = &v
	return s
}

func (s *ListSnapshotRequest) SetPageNumber(v int32) *ListSnapshotRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotRequest) SetPageSize(v int32) *ListSnapshotRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotRequest) SetRegionId(v string) *ListSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotRequest) SetSnapshotIds(v string) *ListSnapshotRequest {
	s.SnapshotIds = &v
	return s
}

func (s *ListSnapshotRequest) SetSnapshotName(v string) *ListSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotRequest) SetSnapshotType(v string) *ListSnapshotRequest {
	s.SnapshotType = &v
	return s
}

func (s *ListSnapshotRequest) SetSortKey(v string) *ListSnapshotRequest {
	s.SortKey = &v
	return s
}

func (s *ListSnapshotRequest) SetSortType(v string) *ListSnapshotRequest {
	s.SortType = &v
	return s
}

func (s *ListSnapshotRequest) SetStatus(v string) *ListSnapshotRequest {
	s.Status = &v
	return s
}

type ListSnapshotResponseBody struct {
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots  []*ListSnapshotResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotResponseBody) SetPageNumber(v int32) *ListSnapshotResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotResponseBody) SetPageSize(v int32) *ListSnapshotResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotResponseBody) SetRequestId(v string) *ListSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotResponseBody) SetSnapshots(v []*ListSnapshotResponseBodySnapshots) *ListSnapshotResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotResponseBody) SetTotalCount(v int32) *ListSnapshotResponseBody {
	s.TotalCount = &v
	return s
}

type ListSnapshotResponseBodySnapshots struct {
	Category            *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CreationTime        *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LastModifiedTime    *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Progress            *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RemainTime          *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	RetentionDays       *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SnapshotId          *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName        *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType        *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SourceFsId          *string `json:"SourceFsId,omitempty" xml:"SourceFsId,omitempty"`
	SourceFsSize        *int32  `json:"SourceFsSize,omitempty" xml:"SourceFsSize,omitempty"`
	SourceFsStripeWidth *int32  `json:"SourceFsStripeWidth,omitempty" xml:"SourceFsStripeWidth,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSnapshotResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotResponseBodySnapshots) SetCategory(v string) *ListSnapshotResponseBodySnapshots {
	s.Category = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetCreationTime(v string) *ListSnapshotResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetDescription(v string) *ListSnapshotResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetLastModifiedTime(v string) *ListSnapshotResponseBodySnapshots {
	s.LastModifiedTime = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetProgress(v string) *ListSnapshotResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetRemainTime(v int32) *ListSnapshotResponseBodySnapshots {
	s.RemainTime = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetRetentionDays(v int32) *ListSnapshotResponseBodySnapshots {
	s.RetentionDays = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSnapshotId(v string) *ListSnapshotResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSnapshotName(v string) *ListSnapshotResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSnapshotType(v string) *ListSnapshotResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSourceFsId(v string) *ListSnapshotResponseBodySnapshots {
	s.SourceFsId = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSourceFsSize(v int32) *ListSnapshotResponseBodySnapshots {
	s.SourceFsSize = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetSourceFsStripeWidth(v int32) *ListSnapshotResponseBodySnapshots {
	s.SourceFsStripeWidth = &v
	return s
}

func (s *ListSnapshotResponseBodySnapshots) SetStatus(v string) *ListSnapshotResponseBodySnapshots {
	s.Status = &v
	return s
}

type ListSnapshotResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotResponse) SetHeaders(v map[string]*string) *ListSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotResponse) SetStatusCode(v int32) *ListSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotResponse) SetBody(v *ListSnapshotResponseBody) *ListSnapshotResponse {
	s.Body = v
	return s
}

type ListSnapshotLinksRequest struct {
	FilterKey   *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	FsIds       *string `json:"FsIds,omitempty" xml:"FsIds,omitempty"`
	LinkIds     *string `json:"LinkIds,omitempty" xml:"LinkIds,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSnapshotLinksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotLinksRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotLinksRequest) SetFilterKey(v string) *ListSnapshotLinksRequest {
	s.FilterKey = &v
	return s
}

func (s *ListSnapshotLinksRequest) SetFilterValue(v string) *ListSnapshotLinksRequest {
	s.FilterValue = &v
	return s
}

func (s *ListSnapshotLinksRequest) SetFsIds(v string) *ListSnapshotLinksRequest {
	s.FsIds = &v
	return s
}

func (s *ListSnapshotLinksRequest) SetLinkIds(v string) *ListSnapshotLinksRequest {
	s.LinkIds = &v
	return s
}

func (s *ListSnapshotLinksRequest) SetPageNumber(v int32) *ListSnapshotLinksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotLinksRequest) SetPageSize(v int32) *ListSnapshotLinksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotLinksRequest) SetRegionId(v string) *ListSnapshotLinksRequest {
	s.RegionId = &v
	return s
}

type ListSnapshotLinksResponseBody struct {
	PageNumber    *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotLinks []*ListSnapshotLinksResponseBodySnapshotLinks `json:"SnapshotLinks,omitempty" xml:"SnapshotLinks,omitempty" type:"Repeated"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotLinksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotLinksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotLinksResponseBody) SetPageNumber(v int32) *ListSnapshotLinksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotLinksResponseBody) SetPageSize(v int32) *ListSnapshotLinksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotLinksResponseBody) SetRequestId(v string) *ListSnapshotLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotLinksResponseBody) SetSnapshotLinks(v []*ListSnapshotLinksResponseBodySnapshotLinks) *ListSnapshotLinksResponseBody {
	s.SnapshotLinks = v
	return s
}

func (s *ListSnapshotLinksResponseBody) SetTotalCount(v int32) *ListSnapshotLinksResponseBody {
	s.TotalCount = &v
	return s
}

type ListSnapshotLinksResponseBodySnapshotLinks struct {
	EcsList       []*ListSnapshotLinksResponseBodySnapshotLinksEcsList `json:"EcsList,omitempty" xml:"EcsList,omitempty" type:"Repeated"`
	FsId          *string                                              `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName        *string                                              `json:"FsName,omitempty" xml:"FsName,omitempty"`
	LinkId        *string                                              `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	SnapshotCount *int32                                               `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	SourceSize    *int32                                               `json:"SourceSize,omitempty" xml:"SourceSize,omitempty"`
	Status        *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalSize     *int64                                               `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListSnapshotLinksResponseBodySnapshotLinks) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotLinksResponseBodySnapshotLinks) GoString() string {
	return s.String()
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetEcsList(v []*ListSnapshotLinksResponseBodySnapshotLinksEcsList) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.EcsList = v
	return s
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetFsId(v string) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.FsId = &v
	return s
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetFsName(v string) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.FsName = &v
	return s
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetLinkId(v string) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.LinkId = &v
	return s
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetSnapshotCount(v int32) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.SnapshotCount = &v
	return s
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetSourceSize(v int32) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.SourceSize = &v
	return s
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetStatus(v string) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.Status = &v
	return s
}

func (s *ListSnapshotLinksResponseBodySnapshotLinks) SetTotalSize(v int64) *ListSnapshotLinksResponseBodySnapshotLinks {
	s.TotalSize = &v
	return s
}

type ListSnapshotLinksResponseBodySnapshotLinksEcsList struct {
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
}

func (s ListSnapshotLinksResponseBodySnapshotLinksEcsList) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotLinksResponseBodySnapshotLinksEcsList) GoString() string {
	return s.String()
}

func (s *ListSnapshotLinksResponseBodySnapshotLinksEcsList) SetEcsId(v string) *ListSnapshotLinksResponseBodySnapshotLinksEcsList {
	s.EcsId = &v
	return s
}

type ListSnapshotLinksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSnapshotLinksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSnapshotLinksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotLinksResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotLinksResponse) SetHeaders(v map[string]*string) *ListSnapshotLinksResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotLinksResponse) SetStatusCode(v int32) *ListSnapshotLinksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotLinksResponse) SetBody(v *ListSnapshotLinksResponseBody) *ListSnapshotLinksResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

type ListTagKeysResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*string) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

type ListTagValuesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ModifyAutoSnapshotPolicyRequest struct {
	PolicyId       *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName     *string   `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdays []*string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty" type:"Repeated"`
	RetentionDays  *int32    `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	TimePoints     []*string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s ModifyAutoSnapshotPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyRequest) SetPolicyId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetPolicyName(v string) *ModifyAutoSnapshotPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRegionId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRepeatWeekdays(v []*string) *ModifyAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetTimePoints(v []*string) *ModifyAutoSnapshotPolicyRequest {
	s.TimePoints = v
	return s
}

type ModifyAutoSnapshotPolicyShrinkRequest struct {
	PolicyId             *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName           *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdaysShrink *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	RetentionDays        *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	TimePointsShrink     *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s ModifyAutoSnapshotPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyShrinkRequest) SetPolicyId(v string) *ModifyAutoSnapshotPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyShrinkRequest) SetPolicyName(v string) *ModifyAutoSnapshotPolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyShrinkRequest) SetRegionId(v string) *ModifyAutoSnapshotPolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyShrinkRequest) SetRepeatWeekdaysShrink(v string) *ModifyAutoSnapshotPolicyShrinkRequest {
	s.RepeatWeekdaysShrink = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyShrinkRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyShrinkRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyShrinkRequest) SetTimePointsShrink(v string) *ModifyAutoSnapshotPolicyShrinkRequest {
	s.TimePointsShrink = &v
	return s
}

type ModifyAutoSnapshotPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoSnapshotPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyResponseBody) SetRequestId(v string) *ModifyAutoSnapshotPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAutoSnapshotPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAutoSnapshotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAutoSnapshotPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyResponse) SetHeaders(v map[string]*string) *ModifyAutoSnapshotPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoSnapshotPolicyResponse) SetStatusCode(v int32) *ModifyAutoSnapshotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyResponse) SetBody(v *ModifyAutoSnapshotPolicyResponseBody) *ModifyAutoSnapshotPolicyResponse {
	s.Body = v
	return s
}

type ModifySnapshotAttributeRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId   *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s ModifySnapshotAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySnapshotAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeRequest) SetDescription(v string) *ModifySnapshotAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetRegionId(v string) *ModifySnapshotAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotId(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotName(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotName = &v
	return s
}

type ModifySnapshotAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySnapshotAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySnapshotAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeResponseBody) SetRequestId(v string) *ModifySnapshotAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifySnapshotAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySnapshotAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySnapshotAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySnapshotAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeResponse) SetHeaders(v map[string]*string) *ModifySnapshotAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifySnapshotAttributeResponse) SetStatusCode(v int32) *ModifySnapshotAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySnapshotAttributeResponse) SetBody(v *ModifySnapshotAttributeResponseBody) *ModifySnapshotAttributeResponse {
	s.Body = v
	return s
}

type RenameDbfsRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	FsName   *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenameDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDbfsRequest) GoString() string {
	return s.String()
}

func (s *RenameDbfsRequest) SetFsId(v string) *RenameDbfsRequest {
	s.FsId = &v
	return s
}

func (s *RenameDbfsRequest) SetFsName(v string) *RenameDbfsRequest {
	s.FsName = &v
	return s
}

func (s *RenameDbfsRequest) SetRegionId(v string) *RenameDbfsRequest {
	s.RegionId = &v
	return s
}

type RenameDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *RenameDbfsResponseBody) SetRequestId(v string) *RenameDbfsResponseBody {
	s.RequestId = &v
	return s
}

type RenameDbfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenameDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDbfsResponse) GoString() string {
	return s.String()
}

func (s *RenameDbfsResponse) SetHeaders(v map[string]*string) *RenameDbfsResponse {
	s.Headers = v
	return s
}

func (s *RenameDbfsResponse) SetStatusCode(v int32) *RenameDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameDbfsResponse) SetBody(v *RenameDbfsResponseBody) *RenameDbfsResponse {
	s.Body = v
	return s
}

type ResizeDbfsRequest struct {
	FsId     *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	NewSizeG *int32  `json:"NewSizeG,omitempty" xml:"NewSizeG,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResizeDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeDbfsRequest) GoString() string {
	return s.String()
}

func (s *ResizeDbfsRequest) SetFsId(v string) *ResizeDbfsRequest {
	s.FsId = &v
	return s
}

func (s *ResizeDbfsRequest) SetNewSizeG(v int32) *ResizeDbfsRequest {
	s.NewSizeG = &v
	return s
}

func (s *ResizeDbfsRequest) SetRegionId(v string) *ResizeDbfsRequest {
	s.RegionId = &v
	return s
}

type ResizeDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeDbfsResponseBody) SetRequestId(v string) *ResizeDbfsResponseBody {
	s.RequestId = &v
	return s
}

type ResizeDbfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResizeDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeDbfsResponse) GoString() string {
	return s.String()
}

func (s *ResizeDbfsResponse) SetHeaders(v map[string]*string) *ResizeDbfsResponse {
	s.Headers = v
	return s
}

func (s *ResizeDbfsResponse) SetStatusCode(v int32) *ResizeDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeDbfsResponse) SetBody(v *ResizeDbfsResponseBody) *ResizeDbfsResponse {
	s.Body = v
	return s
}

type TagDbfsRequest struct {
	DbfsId   *string `json:"DbfsId,omitempty" xml:"DbfsId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s TagDbfsRequest) GoString() string {
	return s.String()
}

func (s *TagDbfsRequest) SetDbfsId(v string) *TagDbfsRequest {
	s.DbfsId = &v
	return s
}

func (s *TagDbfsRequest) SetRegionId(v string) *TagDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *TagDbfsRequest) SetTags(v string) *TagDbfsRequest {
	s.Tags = &v
	return s
}

type TagDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *TagDbfsResponseBody) SetRequestId(v string) *TagDbfsResponseBody {
	s.RequestId = &v
	return s
}

type TagDbfsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s TagDbfsResponse) GoString() string {
	return s.String()
}

func (s *TagDbfsResponse) SetHeaders(v map[string]*string) *TagDbfsResponse {
	s.Headers = v
	return s
}

func (s *TagDbfsResponse) SetStatusCode(v int32) *TagDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *TagDbfsResponse) SetBody(v *TagDbfsResponseBody) *TagDbfsResponse {
	s.Body = v
	return s
}

type UpdateDbfsRequest struct {
	AdvancedFeatures *string `json:"AdvancedFeatures,omitempty" xml:"AdvancedFeatures,omitempty"`
	FsId             *string `json:"FsId,omitempty" xml:"FsId,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UsedScene        *string `json:"UsedScene,omitempty" xml:"UsedScene,omitempty"`
}

func (s UpdateDbfsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDbfsRequest) GoString() string {
	return s.String()
}

func (s *UpdateDbfsRequest) SetAdvancedFeatures(v string) *UpdateDbfsRequest {
	s.AdvancedFeatures = &v
	return s
}

func (s *UpdateDbfsRequest) SetFsId(v string) *UpdateDbfsRequest {
	s.FsId = &v
	return s
}

func (s *UpdateDbfsRequest) SetInstanceType(v string) *UpdateDbfsRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateDbfsRequest) SetRegionId(v string) *UpdateDbfsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDbfsRequest) SetUsedScene(v string) *UpdateDbfsRequest {
	s.UsedScene = &v
	return s
}

type UpdateDbfsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDbfsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDbfsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDbfsResponseBody) SetRequestId(v string) *UpdateDbfsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDbfsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDbfsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDbfsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDbfsResponse) GoString() string {
	return s.String()
}

func (s *UpdateDbfsResponse) SetHeaders(v map[string]*string) *UpdateDbfsResponse {
	s.Headers = v
	return s
}

func (s *UpdateDbfsResponse) SetStatusCode(v int32) *UpdateDbfsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDbfsResponse) SetBody(v *UpdateDbfsResponseBody) *UpdateDbfsResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("dbfs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("dbfs.aliyuncs.com"),
		"cn-edge-1":                   tea.String("dbfs.aliyuncs.com"),
		"cn-fujian":                   tea.String("dbfs.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("dbfs.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("dbfs.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("dbfs.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("dbfs.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("dbfs.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("dbfs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("dbfs.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("dbfs.aliyuncs.com"),
		"cn-wuhan":                    tea.String("dbfs.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("dbfs.aliyuncs.com"),
		"cn-yushanfang":               tea.String("dbfs.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("dbfs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("dbfs.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("dbfs.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("dbfs.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("dbfs.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("dbfs.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dbfs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTagsBatchWithOptions(request *AddTagsBatchRequest, runtime *util.RuntimeOptions) (_result *AddTagsBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbfsList)) {
		query["DbfsList"] = request.DbfsList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTagsBatch"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTagsBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTagsBatch(request *AddTagsBatchRequest) (_result *AddTagsBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsBatchResponse{}
	_body, _err := client.AddTagsBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyAutoSnapshotPolicyWithOptions(tmpReq *ApplyAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyAutoSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DbfsIds)) {
		request.DbfsIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DbfsIds, tea.String("DbfsIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbfsIdsShrink)) {
		query["DbfsIds"] = request.DbfsIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyAutoSnapshotPolicy"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyAutoSnapshotPolicy(request *ApplyAutoSnapshotPolicyRequest) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.ApplyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachDbfsWithOptions(request *AttachDbfsRequest, runtime *util.RuntimeOptions) (_result *AttachDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachMode)) {
		query["AttachMode"] = request.AttachMode
	}

	if !tea.BoolValue(util.IsUnset(request.AttachPoint)) {
		query["AttachPoint"] = request.AttachPoint
	}

	if !tea.BoolValue(util.IsUnset(request.ECSInstanceId)) {
		query["ECSInstanceId"] = request.ECSInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerUrl)) {
		query["ServerUrl"] = request.ServerUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachDbfs(request *AttachDbfsRequest) (_result *AttachDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachDbfsResponse{}
	_body, _err := client.AttachDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelAutoSnapshotPolicyWithOptions(tmpReq *CancelAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CancelAutoSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DbfsIds)) {
		request.DbfsIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DbfsIds, tea.String("DbfsIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbfsIdsShrink)) {
		query["DbfsIds"] = request.DbfsIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelAutoSnapshotPolicy"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelAutoSnapshotPolicy(request *CancelAutoSnapshotPolicyRequest) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CancelAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAutoSnapshotPolicyWithOptions(tmpReq *CreateAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAutoSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RepeatWeekdays)) {
		request.RepeatWeekdaysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatWeekdays, tea.String("RepeatWeekdays"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TimePoints)) {
		request.TimePointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TimePoints, tea.String("TimePoints"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatWeekdaysShrink)) {
		query["RepeatWeekdays"] = request.RepeatWeekdaysShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.TimePointsShrink)) {
		query["TimePoints"] = request.TimePointsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAutoSnapshotPolicy"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CreateAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDbfsWithOptions(request *CreateDbfsRequest, runtime *util.RuntimeOptions) (_result *CreateDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdvancedFeatures)) {
		query["AdvancedFeatures"] = request.AdvancedFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteSnapshot)) {
		query["DeleteSnapshot"] = request.DeleteSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.EnableRaid)) {
		query["EnableRaid"] = request.EnableRaid
	}

	if !tea.BoolValue(util.IsUnset(request.Encryption)) {
		query["Encryption"] = request.Encryption
	}

	if !tea.BoolValue(util.IsUnset(request.FsName)) {
		query["FsName"] = request.FsName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.KMSKeyId)) {
		query["KMSKeyId"] = request.KMSKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.PerformanceLevel)) {
		query["PerformanceLevel"] = request.PerformanceLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RaidStripeUnitNumber)) {
		query["RaidStripeUnitNumber"] = request.RaidStripeUnitNumber
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SizeG)) {
		query["SizeG"] = request.SizeG
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.UsedScene)) {
		query["UsedScene"] = request.UsedScene
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDbfs(request *CreateDbfsRequest) (_result *CreateDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDbfsResponse{}
	_body, _err := client.CreateDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceLinkedRole"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAutoSnapshotPolicyWithOptions(request *DeleteAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutoSnapshotPolicy"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAutoSnapshotPolicy(request *DeleteAutoSnapshotPolicyRequest) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.DeleteAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDbfsWithOptions(request *DeleteDbfsRequest, runtime *util.RuntimeOptions) (_result *DeleteDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDbfs(request *DeleteDbfsRequest) (_result *DeleteDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDbfsResponse{}
	_body, _err := client.DeleteDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagsBatchWithOptions(request *DeleteTagsBatchRequest, runtime *util.RuntimeOptions) (_result *DeleteTagsBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbfsList)) {
		query["DbfsList"] = request.DbfsList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTagsBatch"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTagsBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTagsBatch(request *DeleteTagsBatchRequest) (_result *DeleteTagsBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagsBatchResponse{}
	_body, _err := client.DeleteTagsBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDbfsSpecificationsWithOptions(request *DescribeDbfsSpecificationsRequest, runtime *util.RuntimeOptions) (_result *DescribeDbfsSpecificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceType)) {
		query["EcsInstanceType"] = request.EcsInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDbfsSpecifications"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDbfsSpecificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDbfsSpecifications(request *DescribeDbfsSpecificationsRequest) (_result *DescribeDbfsSpecificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDbfsSpecificationsResponse{}
	_body, _err := client.DescribeDbfsSpecificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTypesWithOptions(request *DescribeInstanceTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceTypes"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (_result *DescribeInstanceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.DescribeInstanceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachDbfsWithOptions(request *DetachDbfsRequest, runtime *util.RuntimeOptions) (_result *DetachDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ECSInstanceId)) {
		query["ECSInstanceId"] = request.ECSInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachDbfs(request *DetachDbfsRequest) (_result *DetachDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachDbfsResponse{}
	_body, _err := client.DetachDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutoSnapshotPolicyWithOptions(request *GetAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *GetAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutoSnapshotPolicy"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutoSnapshotPolicy(request *GetAutoSnapshotPolicyRequest) (_result *GetAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoSnapshotPolicyResponse{}
	_body, _err := client.GetAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDbfsWithOptions(request *GetDbfsRequest, runtime *util.RuntimeOptions) (_result *GetDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDbfs(request *GetDbfsRequest) (_result *GetDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDbfsResponse{}
	_body, _err := client.GetDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceLinkedRoleWithOptions(request *GetServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *GetServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceLinkedRole"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceLinkedRole(request *GetServiceLinkedRoleRequest) (_result *GetServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceLinkedRoleResponse{}
	_body, _err := client.GetServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSnapshotLinkWithOptions(request *GetSnapshotLinkRequest, runtime *util.RuntimeOptions) (_result *GetSnapshotLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSnapshotLink"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSnapshotLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSnapshotLink(request *GetSnapshotLinkRequest) (_result *GetSnapshotLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSnapshotLinkResponse{}
	_body, _err := client.GetSnapshotLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAutoSnapshotPoliciesWithOptions(request *ListAutoSnapshotPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListAutoSnapshotPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutoSnapshotPolicies"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAutoSnapshotPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAutoSnapshotPolicies(request *ListAutoSnapshotPoliciesRequest) (_result *ListAutoSnapshotPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutoSnapshotPoliciesResponse{}
	_body, _err := client.ListAutoSnapshotPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAutoSnapshotPolicyAppliedDbfsWithOptions(request *ListAutoSnapshotPolicyAppliedDbfsRequest, runtime *util.RuntimeOptions) (_result *ListAutoSnapshotPolicyAppliedDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutoSnapshotPolicyAppliedDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAutoSnapshotPolicyAppliedDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAutoSnapshotPolicyAppliedDbfs(request *ListAutoSnapshotPolicyAppliedDbfsRequest) (_result *ListAutoSnapshotPolicyAppliedDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutoSnapshotPolicyAppliedDbfsResponse{}
	_body, _err := client.ListAutoSnapshotPolicyAppliedDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAutoSnapshotPolicyUnappliedDbfsWithOptions(request *ListAutoSnapshotPolicyUnappliedDbfsRequest, runtime *util.RuntimeOptions) (_result *ListAutoSnapshotPolicyUnappliedDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutoSnapshotPolicyUnappliedDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAutoSnapshotPolicyUnappliedDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAutoSnapshotPolicyUnappliedDbfs(request *ListAutoSnapshotPolicyUnappliedDbfsRequest) (_result *ListAutoSnapshotPolicyUnappliedDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutoSnapshotPolicyUnappliedDbfsResponse{}
	_body, _err := client.ListAutoSnapshotPolicyUnappliedDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDbfsWithOptions(request *ListDbfsRequest, runtime *util.RuntimeOptions) (_result *ListDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortKey)) {
		query["SortKey"] = request.SortKey
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDbfs(request *ListDbfsRequest) (_result *ListDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDbfsResponse{}
	_body, _err := client.ListDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDbfsAttachableEcsInstancesWithOptions(request *ListDbfsAttachableEcsInstancesRequest, runtime *util.RuntimeOptions) (_result *ListDbfsAttachableEcsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDbfsAttachableEcsInstances"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDbfsAttachableEcsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDbfsAttachableEcsInstances(request *ListDbfsAttachableEcsInstancesRequest) (_result *ListDbfsAttachableEcsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDbfsAttachableEcsInstancesResponse{}
	_body, _err := client.ListDbfsAttachableEcsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDbfsAttachedEcsInstancesWithOptions(request *ListDbfsAttachedEcsInstancesRequest, runtime *util.RuntimeOptions) (_result *ListDbfsAttachedEcsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDbfsAttachedEcsInstances"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDbfsAttachedEcsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDbfsAttachedEcsInstances(request *ListDbfsAttachedEcsInstancesRequest) (_result *ListDbfsAttachedEcsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDbfsAttachedEcsInstancesResponse{}
	_body, _err := client.ListDbfsAttachedEcsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSnapshotWithOptions(request *ListSnapshotRequest, runtime *util.RuntimeOptions) (_result *ListSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotType)) {
		query["SnapshotType"] = request.SnapshotType
	}

	if !tea.BoolValue(util.IsUnset(request.SortKey)) {
		query["SortKey"] = request.SortKey
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSnapshot"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSnapshot(request *ListSnapshotRequest) (_result *ListSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSnapshotResponse{}
	_body, _err := client.ListSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSnapshotLinksWithOptions(request *ListSnapshotLinksRequest, runtime *util.RuntimeOptions) (_result *ListSnapshotLinksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.FsIds)) {
		query["FsIds"] = request.FsIds
	}

	if !tea.BoolValue(util.IsUnset(request.LinkIds)) {
		query["LinkIds"] = request.LinkIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSnapshotLinks"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSnapshotLinksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSnapshotLinks(request *ListSnapshotLinksRequest) (_result *ListSnapshotLinksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSnapshotLinksResponse{}
	_body, _err := client.ListSnapshotLinksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAutoSnapshotPolicyWithOptions(tmpReq *ModifyAutoSnapshotPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAutoSnapshotPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RepeatWeekdays)) {
		request.RepeatWeekdaysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatWeekdays, tea.String("RepeatWeekdays"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TimePoints)) {
		request.TimePointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TimePoints, tea.String("TimePoints"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatWeekdaysShrink)) {
		query["RepeatWeekdays"] = request.RepeatWeekdaysShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.TimePointsShrink)) {
		query["TimePoints"] = request.TimePointsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAutoSnapshotPolicy"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAutoSnapshotPolicy(request *ModifyAutoSnapshotPolicyRequest) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.ModifyAutoSnapshotPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySnapshotAttributeWithOptions(request *ModifySnapshotAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifySnapshotAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySnapshotAttribute"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySnapshotAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (_result *ModifySnapshotAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySnapshotAttributeResponse{}
	_body, _err := client.ModifySnapshotAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameDbfsWithOptions(request *RenameDbfsRequest, runtime *util.RuntimeOptions) (_result *RenameDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.FsName)) {
		query["FsName"] = request.FsName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameDbfs(request *RenameDbfsRequest) (_result *RenameDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameDbfsResponse{}
	_body, _err := client.RenameDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeDbfsWithOptions(request *ResizeDbfsRequest, runtime *util.RuntimeOptions) (_result *ResizeDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.NewSizeG)) {
		query["NewSizeG"] = request.NewSizeG
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeDbfs(request *ResizeDbfsRequest) (_result *ResizeDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeDbfsResponse{}
	_body, _err := client.ResizeDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagDbfsWithOptions(request *TagDbfsRequest, runtime *util.RuntimeOptions) (_result *TagDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbfsId)) {
		query["DbfsId"] = request.DbfsId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagDbfs(request *TagDbfsRequest) (_result *TagDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagDbfsResponse{}
	_body, _err := client.TagDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDbfsWithOptions(request *UpdateDbfsRequest, runtime *util.RuntimeOptions) (_result *UpdateDbfsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdvancedFeatures)) {
		query["AdvancedFeatures"] = request.AdvancedFeatures
	}

	if !tea.BoolValue(util.IsUnset(request.FsId)) {
		query["FsId"] = request.FsId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UsedScene)) {
		query["UsedScene"] = request.UsedScene
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDbfs"),
		Version:     tea.String("2020-04-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDbfsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDbfs(request *UpdateDbfsRequest) (_result *UpdateDbfsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDbfsResponse{}
	_body, _err := client.UpdateDbfsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
