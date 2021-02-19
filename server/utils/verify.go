package utils

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"UUID": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	AppVerify              = Rules{"Name": {NotEmpty()}, "KafkaInputTopic": {NotEmpty()}, "KafkaOutputTopic": {NotEmpty()}}
	TemplateAlarmVerify    = Rules{"Interval": {NotEmpty()}, "Name": {NotEmpty()}, "TemplateId": {NotEmpty(), Ge("0")}, "App": {NotEmpty()}, "Email": {EmailLegal()}}
	LevelAlarmVerify    = Rules{"Interval": {NotEmpty()}, "Name": {NotEmpty()}, "App": {NotEmpty()}, "Level": {NotEmpty()}, "Email": {EmailLegal()}}
	RegexAlarmVerify = Rules{"Interval": {NotEmpty()}, "Name": {NotEmpty()}, "App": {NotEmpty()}, "Regex": {NotEmpty()}, "Email": {EmailLegal()}}
)
