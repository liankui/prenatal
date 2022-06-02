package model

type User struct {
	BaseModel
	// 微信信息
	Openid    string `json:"openid" gorm:"column:openid;type:string;size:100;not null;default:'';comment:用户在当前小程序的唯一标识"`
	Unionid   string `json:"unionid" gorm:"column:unionid;type:string;size:100;not null;default:'';comment:微信开放平台帐号下的唯一标识"` // 若当前小程序已绑定到微信开放平台帐号
	Nickname  string `json:"nickname" gorm:"column:nickname;type:string;size:100;not null;default:'';comment:用户昵称"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url;type:string;size:1000;not null;default:'';comment:用户头像图片的URL"`
	Gender    int    `json:"gender" gorm:"column:gender;type:int;not null;default:0;comment:用户性别 0-未知 1-男性2-女性"`        // 不再返回，强制返回0
	Country   string `json:"country" gorm:"column:country;type:string;size:50;not null;default:'';comment:用户所在国家"`      // 不再返回
	Province  string `json:"province" gorm:"column:province;type:string;size:50;not null;default:'';comment:用户所在省份"`    // 不再返回
	City      string `json:"city" gorm:"column:city;type:string;size:50;not null;default:'';comment:用户所在城市"`            // 不再返回
	Language  string `json:"language" gorm:"column:language;type:string;size:50;not null;default:'';comment:显示用户信息的语言"` // 不再返回，强制返回"zh_CN"
	Desc      string `json:"desc" gorm:"column:desc;type:string;size:30;not null;default:'';comment:声明获取用户个人信息后的用途"`    // 不超过30个字符

	// 个人信息
	Password    string `json:"password" gorm:"column:password;type:string;size:100;not null;default:'';comment:密码"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:string;size:20;not null;default:'';comment:电话号码"`
	Email       string `json:"email" gorm:"column:email;type:string;size:30;not null;default:''"`
	Company     string `json:"company" gorm:"column:company;type:string;size:100;not null;default:'';comment:公司名称"`
	Job         string `json:"job" gorm:"column:job;type:string;size:30;not null;default:'';comment:职业"`
	Industry    string `json:"industry" gorm:"column:industry;type:string;size:30;not null;default:'';comment:所在行业"`
	Role        string `json:"role" gorm:"column:role;type:string;size:30;not null;default:'';comment:职业角色"`
	Birthday    string `json:"birthday" gorm:"column:birthday;type:string;size:30;not null;default:'';comment:生日"`
	Education   string `json:"education" gorm:"column:education;type:string;size:30;not null;default:'';comment:学历 1-大专及以下 2-本科 3-硕士 4-博士及以上"`
	Major       string `json:"major" gorm:"column:major;type:string;size:30;not null;default:'';comment:专业"`

	// 来源信息
	InviterId int    `json:"inviter_id" gorm:"column:inviter_id;type:int;not null;default:0;comment:邀请者Id"`
	EventId   int    `json:"event_id" gorm:"column:event_id;type:int;not null;default:0;comment:事件Id"`
	Source    string `json:"source" gorm:"column:source;type:string;size:30;not null;default:'';comment:用户来源"`
}

func (*User) TableName() string {
	return "user"
}
