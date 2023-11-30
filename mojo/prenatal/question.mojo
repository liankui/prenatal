
type Question {
    id: String @1

    category: String @2
    question: String @3

    options: String @10
    correct_answer: String @11
    explanation: String @12

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
