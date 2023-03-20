/// user service
///
/// get userinfo
interface User {
    /// get userinfo
    @http.get("/user/v1/users/{id}")
    get_userinfo(name: String @1) //< get userinfo by id
           -> Userinfo
}
