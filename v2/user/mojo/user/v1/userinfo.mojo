
/// userinfo object
type Userinfo {
    id: String @1 //< the user id
    name: String @2 //< the user name

    password: String @10
    email: String @11
    phoneNumber: String @12 //< 电话号码
    industry: String @13 //< 所在行业
    company: String @21 //< 公司名称
    job: String @22 //< 职业
    role: String @23 //< 职业角色
    birthday: String @24 //< 生日
    education: String @25 //< 学历 1-大专及以下 2-本科 3-硕士 4-博士及以上
    major: String @26 //< 专业

    openid: String @51 //< 用户在当前小程序的唯一标识
    unionid: String @52 //< 微信开放平台帐号下的唯一标识
    nickname: String @53 //< 用户昵称
    avatarUrl: String @54 //< 用户头像图片的URL
    gender: Int64 @55 //< 用户性别 0-未知 1-男性 2-女性
    country: String @56 //< 用户所在国家
    province: String @57 //< 用户所在省份
    city: String @58 //< 用户所在城市
    language: String @59 @default(zh_CN)//< 显示用户信息的语言(default: zh_CN)
    desc: String @60 //< 声明获取用户个人信息后的用途

    source: String @90//< 用户来源
    eventId: Int64 @91//< 事件Id
    inviterId: Int64 @92//< 邀请者Id

    create_time: Timestamp @100
    update_time: Timestamp @101
    delete_time: Timestamp @102
}
