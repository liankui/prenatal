
type Answer {
    id: String @1

    question_id: String @5
    question_category: String @6

    user_id: String @10
    user_name: String @12
    answer: String @15

    is_correct: Bool @20
    user_comment: String @21

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
