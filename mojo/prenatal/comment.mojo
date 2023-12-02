
type Comment {
    id: String @1

    question_id: String @5 @db.index
    question_category: String @6

    parent_comment_id: String @30
    comment: String @31
    likes: Int @35

    user_id: String @50 @db.index
    user_name: String @51

    create_time: Timestamp @100
    update_time: Timestamp @101
///    delete_time: db.DeleteTime @102
}
