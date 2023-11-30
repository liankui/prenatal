
type Comment {
    id: String @1

    question_id: String @5
    question_category: String @6

    user_id: String @10
    user_name: String @11
    parent_comment_id: String @15
    comment: String @16

    likes: Int @20

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
