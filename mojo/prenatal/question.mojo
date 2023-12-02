
type Question {
    id: String @1
    category: String @2
    question: String @3

    options: [QuestionOption] @10 @db.json
    explanation: String @15

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
