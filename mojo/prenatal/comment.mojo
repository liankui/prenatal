
type Comment {
    id: String @1

    question_id: String @5
    question_category: String @6

    user_id: String @10
    parent_comment_id: String @11
    comment: String @12

    likes: Int @20

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
