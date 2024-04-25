/// quiz服务
///
/// quiz服务相关的接口
interface Quiz {
    /// 新增question信息
    @entity("Question")
    @http.post("/v1/questions")
    create_question(question: Question @1) -> Question

    /// 获取question信息
    @entity("Question")
    @http.get("/v1/questions/{id}")
    get_question(id: String @1) -> Question

    /// 更新question信息
    @entity("Question")
    @http.put("/v1/questions/{id}")
    update_question(id: String @1
                   question: Question @2)
                    -> Question

    /// 删除question
    @entity("Question")
    @http.delete("/v1/questions/{id}")
    delete_question(id: String @1)

    /// 新增answer
    @entity("Answer")
    @http.post("/v1/Answers")
    create_answer(answer: Answer @1) -> Answer
}
