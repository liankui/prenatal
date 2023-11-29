/// entity服务
///
/// entity服务相关的接口
interface Quiz {
    /// 新增question
    @entity("Question")
    @http.post("/v1/questions")
    create_question(question: Question @1) -> Question

    /// 获取question详情
    @entity("Question")
    @http.get("/v1/questions/{id}")
    get_question(id: String @1) -> Question

    /// 更新question信息
    @entity("Question")
    @http.put("/v1/questions/{id}")
    update_question(question: Question @1) -> Question

    /// 删除question
    @entity("Question")
    @http.delete("/v1/questions/{id}")
    delete_question(id: String @1)
}
