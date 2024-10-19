export interface UserModel {
    id?: number,
    firstName: string,
    lastName: string,
    bio: string,
    username: string
}

export interface DisplayNameUserModel {
    firstName: string,
    lastName: string
}

export interface PostModel {
    likes: number,
    comments: number,
    postText: string,
    userAvatar: string,
    userFirstName: string,
    userLastName: string
}