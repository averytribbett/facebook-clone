export interface UserModel {
    id?: number,
    firstName: string,
    lastName: string,
    bio: string,
    username: string
}

export interface PostModel {
    content: string,
    author: string
}