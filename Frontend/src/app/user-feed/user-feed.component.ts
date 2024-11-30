import { Component } from '@angular/core';
import { PostModel, UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';

@Component({
  selector: 'app-user-feed',
  templateUrl: './user-feed.component.html',
  styleUrls: ['./user-feed.component.css'],
})
export class UserFeedComponent {
  public selectedUser = {
    id: 0,
    firstName: 'Please choose a name...',
    lastName: '',
    username: '',
  } as UserModel;

  public availableUsers: UserModel[] = [];

  // this should be turned into an endpoint, and posts would be retrieved from a database
  public allPosts: PostModel[] = [
    // {
    //   content: "Can't believe it took me this long to get on FakeBook (TM)!",
    //   author: "brownm26csp"
    // },
    // {
    //   content: "I would give anything for a good night's sleep",
    //   author: "brownm26csp"
    // },
    // {
    //   content: "Just won a million dollars. See you never!!!",
    //   author: "averytribbett"
    // },
    // {
    //   content: "My neighbor gave me a cybertruck, but it has a really dumb bumper sticker. Best way to remove without damaging my new car?",
    //   author: "averytribbett"
    // },
    // {
    //   content: "Send cat memes",
    //   author: "cadegithub"
    // },
    // {
    //   content: "Where should I go on my next vacation?",
    //   author: "cadegithub"
    // },
    // {
    //   content: "Just finished running a marathon!",
    //   author: "youssefgithub"
    // },
    // {
    //   content: "Got a huge promotion so now I'm going to buy FakeBook (TM)",
    //   author: "youssefgithub"
    // },
  ];

  public relevantPosts: PostModel[] = [];

  constructor(private userService: UserServiceService) {}

  ngOnInit(): void {
  }

  public changeSelectedUser(user: any): void {
    // const allPosts = this.allPosts.filter(x => x.author === user.target.value) as PostModel[];
    // console.log(allPosts);
    // this.relevantPosts = allPosts;
  }
}
