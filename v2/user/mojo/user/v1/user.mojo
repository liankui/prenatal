/// user service
///
/// get userinfo
interface User {
    /// get user info
    @http.get("/user/v1/users/{id}")
    get_userinfo(name: String @1) //< get user info by id
           -> Userinfo
}
