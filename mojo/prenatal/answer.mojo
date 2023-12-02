
type Answer {
    id: String @1

    question_id: String @5
    question_category: String @6

    answer: String @30
    correct_answer: String @31

    user_id: String @50
    user_name: String @51

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
