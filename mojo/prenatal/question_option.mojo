
type QuestionOption {
    id: String @1

    question_id: String @5 @db.index
    question_category: String @6

    option: String @10 //< 选项
    content: String @11 //< 选项的描述
    is_correct: Bool @15
    explanation: String @20

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
