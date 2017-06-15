// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sesiface provides an interface for the Amazon Simple Email Service.
package sesiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SESAPI is the interface type for ses.SES.
type SESAPI interface {
	DeleteIdentityRequest(*ses.DeleteIdentityInput) (*aws.Request, *ses.DeleteIdentityOutput)

	DeleteIdentity(*ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error)

	DeleteIdentityPolicyRequest(*ses.DeleteIdentityPolicyInput) (*aws.Request, *ses.DeleteIdentityPolicyOutput)

	DeleteIdentityPolicy(*ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error)

	DeleteVerifiedEmailAddressRequest(*ses.DeleteVerifiedEmailAddressInput) (*aws.Request, *ses.DeleteVerifiedEmailAddressOutput)

	DeleteVerifiedEmailAddress(*ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error)

	GetIdentityDKIMAttributesRequest(*ses.GetIdentityDKIMAttributesInput) (*aws.Request, *ses.GetIdentityDKIMAttributesOutput)

	GetIdentityDKIMAttributes(*ses.GetIdentityDKIMAttributesInput) (*ses.GetIdentityDKIMAttributesOutput, error)

	GetIdentityNotificationAttributesRequest(*ses.GetIdentityNotificationAttributesInput) (*aws.Request, *ses.GetIdentityNotificationAttributesOutput)

	GetIdentityNotificationAttributes(*ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error)

	GetIdentityPoliciesRequest(*ses.GetIdentityPoliciesInput) (*aws.Request, *ses.GetIdentityPoliciesOutput)

	GetIdentityPolicies(*ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error)

	GetIdentityVerificationAttributesRequest(*ses.GetIdentityVerificationAttributesInput) (*aws.Request, *ses.GetIdentityVerificationAttributesOutput)

	GetIdentityVerificationAttributes(*ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error)

	GetSendQuotaRequest(*ses.GetSendQuotaInput) (*aws.Request, *ses.GetSendQuotaOutput)

	GetSendQuota(*ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error)

	GetSendStatisticsRequest(*ses.GetSendStatisticsInput) (*aws.Request, *ses.GetSendStatisticsOutput)

	GetSendStatistics(*ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error)

	ListIdentitiesRequest(*ses.ListIdentitiesInput) (*aws.Request, *ses.ListIdentitiesOutput)

	ListIdentities(*ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error)

	ListIdentitiesPages(*ses.ListIdentitiesInput, func(*ses.ListIdentitiesOutput, bool) bool) error

	ListIdentityPoliciesRequest(*ses.ListIdentityPoliciesInput) (*aws.Request, *ses.ListIdentityPoliciesOutput)

	ListIdentityPolicies(*ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error)

	ListVerifiedEmailAddressesRequest(*ses.ListVerifiedEmailAddressesInput) (*aws.Request, *ses.ListVerifiedEmailAddressesOutput)

	ListVerifiedEmailAddresses(*ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error)

	PutIdentityPolicyRequest(*ses.PutIdentityPolicyInput) (*aws.Request, *ses.PutIdentityPolicyOutput)

	PutIdentityPolicy(*ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error)

	SendEmailRequest(*ses.SendEmailInput) (*aws.Request, *ses.SendEmailOutput)

	SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error)

	SendRawEmailRequest(*ses.SendRawEmailInput) (*aws.Request, *ses.SendRawEmailOutput)

	SendRawEmail(*ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)

	SetIdentityDKIMEnabledRequest(*ses.SetIdentityDKIMEnabledInput) (*aws.Request, *ses.SetIdentityDKIMEnabledOutput)

	SetIdentityDKIMEnabled(*ses.SetIdentityDKIMEnabledInput) (*ses.SetIdentityDKIMEnabledOutput, error)

	SetIdentityFeedbackForwardingEnabledRequest(*ses.SetIdentityFeedbackForwardingEnabledInput) (*aws.Request, *ses.SetIdentityFeedbackForwardingEnabledOutput)

	SetIdentityFeedbackForwardingEnabled(*ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)

	SetIdentityNotificationTopicRequest(*ses.SetIdentityNotificationTopicInput) (*aws.Request, *ses.SetIdentityNotificationTopicOutput)

	SetIdentityNotificationTopic(*ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error)

	VerifyDomainDKIMRequest(*ses.VerifyDomainDKIMInput) (*aws.Request, *ses.VerifyDomainDKIMOutput)

	VerifyDomainDKIM(*ses.VerifyDomainDKIMInput) (*ses.VerifyDomainDKIMOutput, error)

	VerifyDomainIdentityRequest(*ses.VerifyDomainIdentityInput) (*aws.Request, *ses.VerifyDomainIdentityOutput)

	VerifyDomainIdentity(*ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error)

	VerifyEmailAddressRequest(*ses.VerifyEmailAddressInput) (*aws.Request, *ses.VerifyEmailAddressOutput)

	VerifyEmailAddress(*ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error)

	VerifyEmailIdentityRequest(*ses.VerifyEmailIdentityInput) (*aws.Request, *ses.VerifyEmailIdentityOutput)

	VerifyEmailIdentity(*ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error)
}